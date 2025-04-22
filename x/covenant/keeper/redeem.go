package keeper

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	"github.com/scalarorg/scalar-core/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/key"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	covExported "github.com/scalarorg/scalar-core/x/covenant/exported"
	cov "github.com/scalarorg/scalar-core/x/covenant/types"
)

func CreateRedeemSessionKey(uid exported.Hash) utils.Key {
	return redeemSessionPrefix.Append(utils.KeyFromBz(uid[:]))
}

func CreateUTXOSnapshotKey(uid exported.Hash) utils.Key {
	return utxoSnapshotPrefix.Append(utils.KeyFromBz(uid[:]))
}

func (k Keeper) GetRedeemSession(ctx sdk.Context, CustodianGroupUID exported.Hash) (*cov.RedeemSession, bool) {
	var results cov.RedeemSession

	ok := k.getStore(ctx).Get(CreateRedeemSessionKey(CustodianGroupUID), &results)
	if !ok {
		k.Logger(ctx).Error("redeem session not found", "custodianGroupUid", CustodianGroupUID)
		return nil, false
	}

	return &results, true
}

func (k Keeper) SetSwitchingForRedeemSession(ctx sdk.Context, custodianGroupUID exported.Hash) error {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("redeem session not found")
	}
	if redeemSession.IsSwitching {
		return fmt.Errorf("redeem session is switching")
	}

	redeemSession.IsSwitching = true

	k.SetRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) UpdatePreparingToExecuting(ctx sdk.Context, custodianGroupUID exported.Hash, sequence uint64) error {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("redeem session not found")
	}

	if redeemSession.CurrentPhase != covExported.Preparing {
		return fmt.Errorf("redeem session is not in preparing phase")
	}

	if !redeemSession.IsSwitching {
		return fmt.Errorf("redeem session is not switching")
	}
	redeemSession.Sequence = sequence
	redeemSession.CurrentPhase = covExported.Executing
	redeemSession.IsSwitching = false
	redeemSession.PhaseExpiredAt = 0 // reset the phase expired at
	// Update the redeem session in storage
	k.SetRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) UpdateExecutingToPreparing(ctx sdk.Context, custodianGroupUID exported.Hash, sequence uint64) error {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("redeem session not found")
	}
	//Check if it the recovering request
	if sequence > 0 {
		//In normal mode
		if redeemSession.Sequence != sequence-1 && redeemSession.CurrentPhase != covExported.Executing {
			return fmt.Errorf("redeem session is not in executing phase")
		}
		if !redeemSession.IsSwitching {
			return fmt.Errorf("redeem session is not in switching state")
		}
	} else {
		log.Debug().Msg("[UpdateExecutingToPreparing] redeem session is in recovering mode")
		//In recovering mode we don check sequence and phase, this is the first request set the phase to preparing
	}

	redeemSession.Sequence = sequence
	redeemSession.CurrentPhase = covExported.Preparing
	redeemSession.IsSwitching = false
	redeemSession.PhaseExpiredAt = uint64(ctx.BlockHeight()) + k.GetParams(ctx).BlockLimitPerSession
	// Update the redeem session in storage
	k.SetRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) RenewRedeemSession(ctx sdk.Context, custodianGroupUID exported.Hash) error {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("redeem session not found")
	}
	if redeemSession.CurrentPhase != covExported.Preparing {
		return fmt.Errorf("redeem session is not in preparing phase")
	}

	if redeemSession.IsSwitching {
		return fmt.Errorf("redeem session is in switching state")
	}

	redeemSession.PhaseExpiredAt = uint64(ctx.BlockHeight()) + k.GetParams(ctx).BlockLimitPerSession
	// Update the redeem session in storage
	k.SetRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) SetRedeemSession(ctx sdk.Context, redeemSession *cov.RedeemSession) {
	k.getStore(ctx).Set(CreateRedeemSessionKey(redeemSession.CustodianGroupUID), redeemSession)
}

func (k Keeper) SetUtxoSnapshot(ctx sdk.Context, utxoSnapshot *cov.UTXOSnapshot) error {
	key := CreateUTXOSnapshotKey(utxoSnapshot.CustodianGroupUID)

	var results cov.UTXOSnapshot
	ok := k.getStore(ctx).Get(key, &results)
	if !ok {
		k.Logger(ctx).Debug("init utxo snapshot not found", "custodianGroupUID", utxoSnapshot.CustodianGroupUID)
		k.setUtxoSnapshot(ctx, utxoSnapshot)
		return nil
	}

	// if len(results.Utxos) == 0 {
	// 	k.setUtxoSnapshot(ctx, utxoSnapshot)
	// 	return nil
	// }
	// Allway update the utxo snapshot with higher block height
	// Update the utxo snapshot
	if results.BlockHeight < utxoSnapshot.BlockHeight {
		k.setUtxoSnapshot(ctx, utxoSnapshot)
		return nil
	}

	return fmt.Errorf("utxo snapshot already exists")
}

func (k Keeper) setUtxoSnapshot(ctx sdk.Context, utxoSnapshot *cov.UTXOSnapshot) {
	k.getStore(ctx).Set(CreateUTXOSnapshotKey(utxoSnapshot.CustodianGroupUID), utxoSnapshot)
}

func (k Keeper) GetUtxoSnapshot(ctx sdk.Context, custodianGroupUID exported.Hash) (*cov.UTXOSnapshot, bool) {
	var results cov.UTXOSnapshot

	ok := k.getStore(ctx).Get(CreateUTXOSnapshotKey(custodianGroupUID), &results)
	if !ok {
		k.Logger(ctx).Error("utxo snapshot not found", "custodianGroupUID", custodianGroupUID)
		return nil, false
	}

	return &results, true
}

func (k Keeper) AppendUtxo(ctx sdk.Context, txID exported.Hash, vout uint32, scriptPubkey []byte, amountInSats uint64) error {
	custodianGroups, ok := k.GetAllCustodianGroups(ctx)
	if !ok {
		return fmt.Errorf("custodian groups not found")
	}

	for _, custodianGroup := range custodianGroups {
		if bytes.Equal(custodianGroup.BitcoinPubkey, scriptPubkey) {
			utxoSnapshot, ok := k.GetUtxoSnapshot(ctx, custodianGroup.UID)
			if !ok {
				return fmt.Errorf("utxo snapshot not found")
			}
			utxoSnapshot.Utxos = append(utxoSnapshot.Utxos, &cov.UTXO{
				TxID:         txID,
				Vout:         vout,
				ScriptPubkey: scriptPubkey,
				AmountInSats: amountInSats,
			})
			k.SetUtxoSnapshot(ctx, utxoSnapshot)
			break
		}
	}
	return nil
}

// findAvailableUtxos uses knapsack algorithm to find optimal UTXO combination
func (k Keeper) reserveUtxos(ctx sdk.Context, custodianGroupUID [32]byte, requestID string, amount uint64) ([]*cov.UTXO, error) {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return nil, fmt.Errorf("redeem session not found")
	}
	if redeemSession.IsSwitching {
		return nil, fmt.Errorf("redeem phase is switching")
	}
	if redeemSession.CurrentPhase != covExported.Preparing {
		return nil, fmt.Errorf("redeem session is not in preparing phase")
	}
	utxoSnapshot, ok := k.GetUtxoSnapshot(ctx, custodianGroupUID)
	if !ok {
		return nil, fmt.Errorf("utxo snapshot not found")
	}

	if len(utxoSnapshot.Utxos) == 0 {
		return nil, fmt.Errorf("no utxos found")
	}

	// Find optimal UTXO combination using knapsack algorithm
	reserveUtxos, err := utxoSnapshot.ReserveUtxos(requestID, amount)
	if err != nil {
		k.Logger(ctx).Error("failed to reserve utxos", "error", err)
		return nil, err
	}

	// Update the redeem session in storage
	k.setUtxoSnapshot(ctx, utxoSnapshot)

	return reserveUtxos, nil
}

func (k Keeper) CreateRedeemParams(ctx sdk.Context, req *cov.ReserveRedeemUtxoRequest, custodianGrUID exported.Hash, sequence uint64) ([]byte, *cov.CommandID, error) {

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(ctx.BlockHeight()))

	// TODO: ensure the reqId is unique and prevent double spending
	// TODO: cache the reqId to check if it's already used

	amountz := make([]byte, 8)
	binary.BigEndian.PutUint64(amountz, req.Amount)

	dataHash := crypto.Keccak256(bz, req.Sender.Bytes(), []byte(req.Address), []byte(req.SourceChain), []byte(req.DestChain), []byte(req.Symbol), amountz)
	reservedUtxos, err := k.reserveUtxos(ctx, custodianGrUID, hex.EncodeToString(dataHash), req.Amount)
	if err != nil {
		return nil, nil, err
	}
	cmdId := cov.NewCommandID(dataHash)
	payload := &cov.RedeemTokenPayloadWithType{
		RedeemTokenPayload: cov.RedeemTokenPayload{
			Amount:        req.Amount,
			LockingScript: req.LockingScript,
			Utxos:         reservedUtxos,
			RequestId:     cmdId.Bytes(),
		},
		PayloadType: encode.ContractCallWithTokenPayloadType_CustodianOnly,
	}
	redeemTokenParams := &cov.RedeemTokenParams{
		DestinationChain:   req.DestChain.String(),
		DestinationAddress: req.Address,
		Payload:            *payload,
		Symbol:             req.Symbol,
		Amount:             req.Amount,
		CustodianGroupUID:  custodianGrUID,
		SessionSequence:    sequence,
	}
	params, err := redeemTokenParams.AbiPack()
	if err != nil {
		return nil, nil, err
	}

	return params, &cmdId, err
}

func (k Keeper) GetReserveUTXOCommandByID(ctx sdk.Context, id []byte) cov.StandaloneCommand {
	md := k.getStandaloneCommandMetadata(ctx, id, reserveUtxoCommandPrefix)

	setter := func(m cov.StandaloneCommandMetadata) {
		k.SetStandaloneCommandMetadata(ctx, m, reserveUtxoCommandPrefix)
	}

	return cov.NewStandaloneCommand(md, setter)
}

func (k Keeper) getStandaloneCommandMetadata(ctx sdk.Context, id []byte, prefix key.Key) cov.StandaloneCommandMetadata {
	var md cov.StandaloneCommandMetadata
	k.getStore(ctx).GetNew(prefix.Append(key.FromBz(id)), &md)
	return md
}

func (k Keeper) SetStandaloneCommandMetadata(ctx sdk.Context, meta cov.StandaloneCommandMetadata, prefix key.Key) {
	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(prefix.Append(key.FromBz(meta.ID)), &meta))
}

func (k Keeper) SetUnsignedStandaloneCommandID(ctx sdk.Context, id []byte) {
	k.getStore(ctx).SetRawNew(unsignedStandaloneIDKey, id)
}

func (k Keeper) MarkReservedUtxo(ctx sdk.Context, uid exported.Hash, payload []byte) error {
	p := &cov.RedeemTokenPayloadWithType{}
	err := p.AbiUnpack(payload)
	if err != nil {
		return err
	}

	utxoSnapshot, ok := k.GetUtxoSnapshot(ctx, uid)
	if !ok {
		return fmt.Errorf("utxo snapshot not found")
	}

	reqId := hex.EncodeToString(p.RequestId[:])
	for _, reqUtxo := range p.Utxos {
		for _, utxo := range utxoSnapshot.Utxos {
			if utxo.TxID.Hex() == reqUtxo.TxID.Hex() && utxo.Vout == reqUtxo.Vout {
				if utxo.Reserved == nil {
					utxo.Reserved = make(map[string]uint64)
				}
				utxo.Reserved[reqId] = reqUtxo.AmountInSats
				break
			}
		}
	}

	k.setUtxoSnapshot(ctx, utxoSnapshot)
	return nil
}
