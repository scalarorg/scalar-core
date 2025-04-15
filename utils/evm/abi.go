package evm

import (
	"fmt"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/scalarorg/scalar-core/utils/funcs"
)

var (
	AddressType      = funcs.Must(ethabi.NewType("address", "address", nil))
	StringType       = funcs.Must(ethabi.NewType("string", "string", nil))
	BytesType        = funcs.Must(ethabi.NewType("bytes", "bytes", nil))
	Bytes32Type      = funcs.Must(ethabi.NewType("bytes32", "bytes32", nil))
	Uint8Type        = funcs.Must(ethabi.NewType("uint8", "uint8", nil))
	Uint256Type      = funcs.Must(ethabi.NewType("uint256", "uint256", nil))
	Uint64Type       = funcs.Must(ethabi.NewType("uint64", "uint64", nil))
	BytesArrayType   = funcs.Must(ethabi.NewType("bytes[]", "bytes[]", nil))
	Uint256ArrayType = funcs.Must(ethabi.NewType("uint256[]", "uint256[]", nil))
	Uint32ArrayType  = funcs.Must(ethabi.NewType("uint32[]", "uint32[]", nil))
	Uint64ArrayType  = funcs.Must(ethabi.NewType("uint64[]", "uint64[]", nil))
	Bytes32ArrayType = funcs.Must(ethabi.NewType("bytes32[]", "bytes32[]", nil))
	StringArrayType  = funcs.Must(ethabi.NewType("string[]", "string[]", nil))
	AddressArrayType = funcs.Must(ethabi.NewType("address[]", "address[]", nil))
)

func AbiUnpack(data []byte, types ...string) ([]interface{}, error) {
	var arguments ethabi.Arguments
	for _, t := range types {
		typ, err := ethabi.NewType(t, t, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create type: %w", err)
		}
		arguments = append(arguments, ethabi.Argument{Type: typ})
	}
	args, err := arguments.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to get arguments: %w", err)
	}
	return args, nil
}
