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
	stringType    = funcs.Must(abi.NewType("string", "string", nil))
	bytesType     = funcs.Must(abi.NewType("bytes", "bytes", nil))
	bytes32Type   = funcs.Must(abi.NewType("bytes32", "bytes32", nil))
	uint8Type     = funcs.Must(abi.NewType("uint8", "uint8", nil))
	uint32Type    = funcs.Must(abi.NewType("uint32", "uint32", nil))
	uint256Type   = funcs.Must(abi.NewType("uint256", "uint256", nil))
	addressType   = funcs.Must(abi.NewType("address", "address", nil))
	addressesType = funcs.Must(abi.NewType("address[]", "address[]", nil))

	uint256ArrayType = funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))
	stringArrayType  = funcs.Must(abi.NewType("string[]", "string[]", nil))

	callContractWithTokenPayloadArguments = abi.Arguments{{Type: uint256Type}, {Type: stringArrayType}, {Type: uint256ArrayType}}
	callContractWithTokenArguments        = abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: bytesType}, {Type: stringType}, {Type: uint256Type}}
)

func CreateRedeemSessionKey(custodianGroupUid string) utils.Key {
	return redeemSessionPrefix.Append(utils.KeyFromStr(custodianGroupUid))
}
func CreateUTXOSnapshotKey(custodianGroupUid string) utils.Key {
	return utxoSnapshotPrefix.Append(utils.KeyFromStr(custodianGroupUid))
}
func (k Keeper) SetRedeemSession(ctx sdk.Context, redeemSession *cov.RedeemSession) {
	k.getStore(ctx).Set(CreateRedeemSessionKey(redeemSession.CustodianGroupUid), redeemSession)
}

func (k Keeper) GetRedeemSessionByCustodianGroupUid(ctx sdk.Context, custodianGroupUid string) (*cov.RedeemSession, bool) {
	var results cov.RedeemSession

	ok := k.getStore(ctx).Get(CreateRedeemSessionKey(custodianGroupUid), &results)
	if !ok {
		k.Logger(ctx).Error("redeem session not found", "custodianGroupUid", custodianGroupUid)
		return nil, false
	}

	return &results, true
}

func (k Keeper) SetUtxoSnapshot(ctx sdk.Context, utxoSnapshot *cov.UTXOSnapshot) {
	k.getStore(ctx).Set(CreateUTXOSnapshotKey(utxoSnapshot.CustodianGroupUid), utxoSnapshot)
}

func (k Keeper) GetUtxoSnapshotByCustodianGroupUid(ctx sdk.Context, custodianGroupUid string) (*cov.UTXOSnapshot, bool) {
	var results cov.UTXOSnapshot

	ok := k.getStore(ctx).Get(CreateUTXOSnapshotKey(custodianGroupUid), &results)
	if !ok {
		k.Logger(ctx).Error("utxo snapshot not found", "custodianGroupUid", custodianGroupUid)
		return nil, false
	}

	return &results, true
}

// findAvailableUtxos uses knapsack algorithm to find optimal UTXO combination
func (k Keeper) reserveUtxos(ctx sdk.Context, custodianGroupUid string, requestID string, amount uint64) ([]*cov.UTXO, error) {
	redeemSession, ok := k.GetRedeemSessionByCustodianGroupUid(ctx, custodianGroupUid)
	if !ok {
		return nil, fmt.Errorf("redeem session not found")
	}
	if redeemSession.CurrentPhase != cov.Preparing {
		return nil, fmt.Errorf("redeem session is not in preparing phase")
	}
	utxoSnapshot, ok := k.GetUtxoSnapshotByCustodianGroupUid(ctx, custodianGroupUid)
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
