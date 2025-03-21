package keeper

import (
	"fmt"

	"github.com/scalarorg/scalar-core/utils"
	types "github.com/scalarorg/scalar-core/x/covenant/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/scalarorg/scalar-core/utils/funcs"
	cov "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (k Keeper) AddRedeemSession(ctx sdk.Context, tokenSymbol string, session *types.RedeemSession) {
	k.getStore(ctx).Set(redeemSessionPrefix.Append(utils.KeyFromStr(tokenSymbol)), session)
}

func (k Keeper) GetRedeemSession(ctx sdk.Context, tokenSymbol string) (*types.RedeemSession, bool) {
	session := types.RedeemSession{}
	ok := k.getStore(ctx).Get(redeemSessionPrefix.Append(utils.KeyFromStr(tokenSymbol)), &session)
	return &session, ok
}

var (
	stringType  = funcs.Must(abi.NewType("string", "string", nil))
	bytesType   = funcs.Must(abi.NewType("bytes", "bytes", nil))
	uint256Type = funcs.Must(abi.NewType("uint256", "uint256", nil))

	addressType      = funcs.Must(abi.NewType("address", "address", nil))
	addressesType    = funcs.Must(abi.NewType("address[]", "address[]", nil))
	bytes32Type      = funcs.Must(abi.NewType("bytes32", "bytes32", nil))
	uint8Type        = funcs.Must(abi.NewType("uint8", "uint8", nil))
	uint256ArrayType = funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))
	//     function callContractWithToken(
	//     string calldata destinationChain,
	//     string calldata destinationContractAddress,
	//     bytes calldata payload,
	//     string calldata symbol,
	//     uint256 amount
	// )
	callContractWithTokenArguments = abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: bytesType}, {Type: stringType}, {Type: uint256Type}}
)

func CreateRedeemSessionKey(symbol string) utils.Key {
	return redeemSessionPrefix.Append(utils.KeyFromStr(symbol))
}

func (k Keeper) SetRedeemSession(ctx sdk.Context, redeemSession *cov.RedeemSession) {
	k.getStore(ctx).Set(CreateRedeemSessionKey(redeemSession.Symbol), redeemSession)
}

func (k Keeper) GetRedeemSessionBySymbol(ctx sdk.Context, symbol string) (*cov.RedeemSession, bool) {
	var results cov.RedeemSession

	ok := k.getStore(ctx).Get(CreateRedeemSessionKey(symbol), &results)
	if !ok {
		k.Logger(ctx).Error("redeem session not found", "symbol", symbol)
		return nil, false
	}

	return &results, true
}

// findAvailableUtxos uses knapsack algorithm to find optimal UTXO combination
func (k Keeper) reserveUtxos(ctx sdk.Context, symbol string, requestID string, amount uint64) ([]*cov.UTXO, error) {
	redeemSession, ok := k.GetRedeemSessionBySymbol(ctx, symbol)
	if !ok {
		return nil, fmt.Errorf("redeem session not found")
	}

	if redeemSession.IsEmpty() {
		return nil, fmt.Errorf("redeem session not found")
	}

	// if len(redeemSession.Utxos) == 0 {
	// 	return nil, fmt.Errorf("no utxos found")
	// }

	// Find optimal UTXO combination using knapsack algorithm
	// reserveUtxos, err := redeemSession.ReserveUtxos(requestID, amount)
	// if err != nil {
	// 	k.Logger(ctx).Error("failed to reserve utxos", "error", err)
	// 	return nil, err
	// }

	// Update the redeem session in storage
	// k.SetRedeemSession(ctx, redeemSession)

	// return reserveUtxos, nil
	return []*cov.UTXO{}, nil
}

func (k Keeper) createRedeemPayload(ctx sdk.Context, destChain string, destAddress string, symbol string,
	amount uint64,
	reservedTx []*cov.UTXO) ([]byte, error) {
	var payload []byte

	return callContractWithTokenArguments.Pack(
		destChain,
		destAddress,
		payload,
		symbol,
		amount,
	)
}
