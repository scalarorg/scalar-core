package keeper

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/scalarorg/scalar-core/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/key"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	covExported "github.com/scalarorg/scalar-core/x/covenant/exported"
	cov "github.com/scalarorg/scalar-core/x/covenant/types"
)

var (
	stringType       = funcs.Must(abi.NewType("string", "string", nil))
	bytesType        = funcs.Must(abi.NewType("bytes", "bytes", nil))
	bytes32Type      = funcs.Must(abi.NewType("bytes32", "bytes32", nil))
	uint256Type      = funcs.Must(abi.NewType("uint256", "uint256", nil))
	uint64Type       = funcs.Must(abi.NewType("uint64", "uint64", nil))
	uint256ArrayType = funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))
	uint32ArrayType  = funcs.Must(abi.NewType("uint32[]", "uint32[]", nil))
	uint64ArrayType  = funcs.Must(abi.NewType("uint64[]", "uint64[]", nil))
	stringArrayType  = funcs.Must(abi.NewType("string[]", "string[]", nil))

	RedeemTokenPayloadArguments    = abi.Arguments{{Type: uint64Type}, {Type: bytesType}, {Type: stringArrayType}, {Type: uint32ArrayType}, {Type: uint64ArrayType}, {Type: bytes32Type}}
	callContractWithTokenArguments = abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: bytesType}, {Type: stringType}, {Type: uint256Type}}
)

func CreateRedeemSessionKey(uid []byte) utils.Key {
	return redeemSessionPrefix.Append(utils.KeyFromBz(uid))
}

func CreateUTXOSnapshotKey(uid []byte) utils.Key {
	return utxoSnapshotPrefix.Append(utils.KeyFromBz(uid))
}

func (k Keeper) GetRedeemSession(ctx sdk.Context, CustodianGroupUID []byte) (*cov.RedeemSession, bool) {
	var results cov.RedeemSession

	ok := k.getStore(ctx).Get(CreateRedeemSessionKey(CustodianGroupUID), &results)
	if !ok {
		k.Logger(ctx).Error("redeem session not found", "custodianGroupUid", CustodianGroupUID)
		return nil, false
	}

	return &results, true
}

func (k Keeper) SetSwitchingForRedeemSession(ctx sdk.Context, custodianGroupUID []byte) error {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("redeem session not found")
	}
	if redeemSession.IsSwitching {
		return fmt.Errorf("redeem session is switching")
	}

	redeemSession.IsSwitching = true

	k.setRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) UpdatePreparingToExecuting(ctx sdk.Context, custodianGroupUID []byte) error {
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

	redeemSession.CurrentPhase = covExported.Executing
	redeemSession.IsSwitching = false
	redeemSession.PhaseExpiredAt = 0 // reset the phase expired at
	// Update the redeem session in storage
	k.setRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) UpdateExecutingToPreparing(ctx sdk.Context, custodianGroupUID []byte, sequence uint64) error {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("redeem session not found")
	}

	if redeemSession.CurrentPhase != covExported.Executing {
		return fmt.Errorf("redeem session is not in executing phase")
	}

	if !redeemSession.IsSwitching {
		return fmt.Errorf("redeem session is not in switching state")
	}
	redeemSession.Sequence = sequence
	redeemSession.CurrentPhase = covExported.Preparing
	redeemSession.IsSwitching = false
	redeemSession.PhaseExpiredAt = uint64(ctx.BlockHeight()) + k.GetParams(ctx).BlockLimitPerSession
	// Update the redeem session in storage
	k.setRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) RenewRedeemSession(ctx sdk.Context, custodianGroupUID []byte) error {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("redeem session not found")
	}
	if redeemSession.CurrentPhase != covExported.Preparing {
		return fmt.Errorf("redeem session is not in executing phase")
	}

	if redeemSession.IsSwitching {
		return fmt.Errorf("redeem session is not in switching state")
	}

	redeemSession.PhaseExpiredAt = uint64(ctx.BlockHeight()) + k.GetParams(ctx).BlockLimitPerSession
	// Update the redeem session in storage
	k.setRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) setRedeemSession(ctx sdk.Context, redeemSession *cov.RedeemSession) {
	k.getStore(ctx).Set(CreateRedeemSessionKey(redeemSession.CustodianGroupUID[:]), redeemSession)
}

func (k Keeper) SetUtxoSnapshot(ctx sdk.Context, utxoSnapshot *cov.UTXOSnapshot) error {
	key := CreateUTXOSnapshotKey(utxoSnapshot.CustodianGroupUID.Bytes())

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
	k.getStore(ctx).Set(CreateUTXOSnapshotKey(utxoSnapshot.CustodianGroupUID.Bytes()), utxoSnapshot)
}

func (k Keeper) GetUtxoSnapshot(ctx sdk.Context, custodianGroupUID []byte) (*cov.UTXOSnapshot, bool) {
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
			utxoSnapshot, ok := k.GetUtxoSnapshot(ctx, custodianGroup.UID.Bytes())
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
func (k Keeper) reserveUtxos(ctx sdk.Context, custodianGroupUID []byte, requestID string, amount uint64) ([]*cov.UTXO, error) {
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
	k.SetUtxoSnapshot(ctx, utxoSnapshot)

	return reserveUtxos, nil
}

func (k Keeper) CreateRedeemParams(ctx sdk.Context, req *cov.ReserveRedeemUtxoRequest, custodianGrUID []byte) ([]byte, *cov.CommandID, error) {

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(ctx.BlockHeight()))

	// TODO: ensure the reqId is unique and prevent double spending
	// TODO: cache the reqId to check if it's already used

	amountz := make([]byte, 8)
	binary.BigEndian.PutUint64(amountz, req.Amount)

	reqId := crypto.Keccak256(bz, req.Sender.Bytes(), []byte(req.Address), []byte(req.SourceChain), []byte(req.DestChain), []byte(req.Symbol), amountz)
	reservedUtxos, err := k.reserveUtxos(ctx, custodianGrUID, hex.EncodeToString(reqId), req.Amount)
	if err != nil {
		return nil, nil, err
	}
	// TODO: optimize the datatype
	txIds := make([]string, len(reservedUtxos))
	vouts := make([]uint32, len(reservedUtxos))
	amounts := make([]uint64, len(reservedUtxos))
	for i, utxo := range reservedUtxos {
		txIds[i] = utxo.TxID.Hex()
		vouts[i] = utxo.Vout
		amounts[i] = utxo.AmountInSats
	}

	cmdId := cov.NewCommandID(reqId)

	payload, err := RedeemTokenPayloadArguments.Pack(
		req.Amount,
		req.LockingScript,
		txIds,
		vouts,
		amounts,
		reqId)
	if err != nil {
		return nil, nil, err
	}
	params, err := callContractWithTokenArguments.Pack(
		req.DestChain,
		req.Address,
		payload,
		req.Symbol,
		req.Amount,
	)

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
