package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/scalarorg/scalar-core/utils/funcs"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

const (
	approveContractCallMaxGasCost = 100000
)

var (
	stringType       = funcs.Must(abi.NewType("string", "string", nil))
	addressType      = funcs.Must(abi.NewType("address", "address", nil))
	addressesType    = funcs.Must(abi.NewType("address[]", "address[]", nil))
	bytes32Type      = funcs.Must(abi.NewType("bytes32", "bytes32", nil))
	uint8Type        = funcs.Must(abi.NewType("uint8", "uint8", nil))
	uint256Type      = funcs.Must(abi.NewType("uint256", "uint256", nil))
	uint256ArrayType = funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))

	approveDestCallArguments = abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: addressType}, {Type: bytes32Type}, {Type: bytes32Type}, {Type: uint256Type}}
)

// TODO: design for generic bridge call

func NewApproveBridgeCallCommandGeneric(
	chainID sdk.Int,
	keyID multisig.KeyID,
	contractAddress common.Address,
	payloadHash common.Hash,
	sourceTxID common.Hash,
	sourceChain nexus.ChainName,
	sender string,
	sourceEventIndex uint64,
	ID string,
) Command {
	commandID := NewCommandID([]byte(ID), chainID)
	return Command{
		ID:         commandID,
		Type:       COMMAND_TYPE_APPROVE_BRIDGE_CALL,
		Params:     createApproveBridgeCallParamsGeneric(contractAddress, payloadHash, sourceTxID, string(sourceChain), sender, sourceEventIndex),
		KeyID:      keyID,
		MaxGasCost: approveContractCallMaxGasCost,
	}
}

func createApproveBridgeCallParamsGeneric(
	contractAddress common.Address,
	payloadHash common.Hash,
	txID common.Hash,
	sourceChain string,
	sender string,
	sourceEventIndex uint64) []byte {

	return funcs.Must(approveDestCallArguments.Pack(
		sourceChain,
		sender,
		contractAddress,
		payloadHash,
		txID,
		new(big.Int).SetUint64(sourceEventIndex),
	))
}
