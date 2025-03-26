package keeper

import (
	"fmt"

	"github.com/scalarorg/scalar-core/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/scalarorg/scalar-core/utils/funcs"
	cov "github.com/scalarorg/scalar-core/x/covenant/types"
)

var (
	stringType       = funcs.Must(abi.NewType("string", "string", nil))
	bytesType        = funcs.Must(abi.NewType("bytes", "bytes", nil))
	uint256Type      = funcs.Must(abi.NewType("uint256", "uint256", nil))
	uint256ArrayType = funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))
	stringArrayType  = funcs.Must(abi.NewType("string[]", "string[]", nil))

	callContractWithTokenPayloadArguments = abi.Arguments{{Type: uint256Type}, {Type: stringArrayType}, {Type: uint256ArrayType}}
	callContractWithTokenArguments        = abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: bytesType}, {Type: stringType}, {Type: uint256Type}}
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

	if redeemSession.CurrentPhase != cov.Preparing {
		return fmt.Errorf("redeem session is not in preparing phase")
	}

	if !redeemSession.IsSwitching {
		return fmt.Errorf("redeem session is not switching")
	}

	redeemSession.CurrentPhase = cov.Executing
	redeemSession.IsSwitching = false
	// Update the redeem session in storage
	k.setRedeemSession(ctx, redeemSession)
	return nil
}

func (k Keeper) UpdateExecutingToPreparing(ctx sdk.Context, custodianGroupUID []byte) error {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("redeem session not found")
	}

	if redeemSession.CurrentPhase != cov.Executing {
		return fmt.Errorf("redeem session is not in executing phase")
	}

	if redeemSession.IsSwitching {
		return fmt.Errorf("redeem session is switching")
	}

	redeemSession.CurrentPhase = cov.Preparing
	redeemSession.IsSwitching = false
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
		k.Logger(ctx).Error("init utxo snapshot not found", "custodianGroupUID", utxoSnapshot.CustodianGroupUID)
		k.setUtxoSnapshot(ctx, utxoSnapshot)
		return nil
	}

	if len(results.Utxos) == 0 {
		k.setUtxoSnapshot(ctx, utxoSnapshot)
		return nil
	}

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

// findAvailableUtxos uses knapsack algorithm to find optimal UTXO combination
func (k Keeper) reserveUtxos(ctx sdk.Context, custodianGroupUID []byte, requestID string, amount uint64) ([]*cov.UTXO, error) {
	redeemSession, ok := k.GetRedeemSession(ctx, custodianGroupUID)
	if !ok {
		return nil, fmt.Errorf("redeem session not found")
	}
	if redeemSession.CurrentPhase != cov.Preparing {
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

func (k Keeper) createRedeemPayload(ctx sdk.Context, destChain string, destAddress string, symbol string,
	amount uint64,
	reservedTx []*cov.UTXO) ([]byte, error) {
	txIds := make([]string, len(reservedTx))
	vouts := make([]uint32, len(reservedTx))
	for i, utxo := range reservedTx {
		txIds[i] = utxo.TxID.Hex()
		vouts[i] = utxo.Vout
	}
	payload, err := callContractWithTokenPayloadArguments.Pack(amount, txIds, vouts)
	if err != nil {
		return nil, err
	}
	return callContractWithTokenArguments.Pack(
		destChain,
		destAddress,
		payload,
		symbol,
		amount,
	)
}
