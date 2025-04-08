package evm

import (
	"fmt"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
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
