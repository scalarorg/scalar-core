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

	RedeemTokenPayloadArguments = abi.Arguments{{Type: uint64Type}, {Type: bytesType}, {Type: stringArrayType}, {Type: uint32ArrayType}, {Type: uint64ArrayType}, {Type: bytes32Type}}
	RedeemTokenArguments        = abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: bytesType}, {Type: stringType}, {Type: uint256Type}, {Type: bytes32Type}, {Type: uint64Type}}
)
