package types

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/scalarorg/scalar-core/utils/funcs"
)

var (
	addressType      = funcs.Must(abi.NewType("address", "address", nil))
	stringType       = funcs.Must(abi.NewType("string", "string", nil))
	bytesType        = funcs.Must(abi.NewType("bytes", "bytes", nil))
	bytes32Type      = funcs.Must(abi.NewType("bytes32", "bytes32", nil))
	uint8Type        = funcs.Must(abi.NewType("uint8", "uint8", nil))
	uint256Type      = funcs.Must(abi.NewType("uint256", "uint256", nil))
	uint64Type       = funcs.Must(abi.NewType("uint64", "uint64", nil))
	bytesArrayType   = funcs.Must(abi.NewType("bytes[]", "bytes[]", nil))
	uint256ArrayType = funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))
	uint32ArrayType  = funcs.Must(abi.NewType("uint32[]", "uint32[]", nil))
	uint64ArrayType  = funcs.Must(abi.NewType("uint64[]", "uint64[]", nil))
	bytes32ArrayType = funcs.Must(abi.NewType("bytes32[]", "bytes32[]", nil))
	stringArrayType  = funcs.Must(abi.NewType("string[]", "string[]", nil))
	addressArrayType = funcs.Must(abi.NewType("address[]", "address[]", nil))

	RedeemCustodianOnlyPayloadAbi = abi.Arguments{
		{Type: uint64Type, Name: "amount"},
		{Type: bytesType, Name: "lockingScript"},
		{Type: stringArrayType, Name: "txIds"},
		{Type: uint32ArrayType, Name: "vouts"},
		{Type: uint64ArrayType, Name: "amounts"},
		{Type: bytes32Type, Name: "requestId"},
	}

	RedeemTokenArguments = abi.Arguments{
		{Type: stringType, Name: "destinationChain"},
		{Type: stringType, Name: "destinationAddress"},
		{Type: bytesType, Name: "payload"},
		{Type: stringType, Name: "symbol"},
		{Type: uint256Type, Name: "amount"},
		{Type: bytes32Type, Name: "custodianGroupUID"},
		{Type: uint64Type, Name: "sessionSequence"},
	}
)
