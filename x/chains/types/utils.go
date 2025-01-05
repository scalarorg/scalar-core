package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils/slices"
)

func packArguments(chainID sdk.Int, commandIDs []CommandID, commands []CommandType, commandParams [][]byte) ([]byte, error) {
	if len(commandIDs) != len(commands) || len(commandIDs) != len(commandParams) {
		return nil, fmt.Errorf("length mismatch for command arguments")
	}

	uint256Type, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return nil, err
	}

	bytes32ArrayType, err := abi.NewType("bytes32[]", "bytes32[]", nil)
	if err != nil {
		return nil, err
	}

	stringArrayType, err := abi.NewType("string[]", "string[]", nil)
	if err != nil {
		return nil, err
	}

	bytesArrayType, err := abi.NewType("bytes[]", "bytes[]", nil)
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{{Type: uint256Type}, {Type: bytes32ArrayType}, {Type: stringArrayType}, {Type: bytesArrayType}}
	result, err := arguments.Pack(
		chainID.BigInt(),
		commandIDs,
		slices.Map(commands, CommandType.String),
		commandParams,
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetSignHash(commandData []byte) common.Hash {
	hash := crypto.Keccak256(commandData)

	// TODO: is this the same across any EVM chain?
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(hash), hash)

	return crypto.Keccak256Hash([]byte(msg))
}

func validateBurnerCode(burnerCode []byte) error {
	burnerCodeHash := crypto.Keccak256Hash(burnerCode).Hex()
	switch burnerCodeHash {
	case BurnerCodeHashV1,
		BurnerCodeHashV2,
		BurnerCodeHashV3,
		BurnerCodeHashV4,
		BurnerCodeHashV5,
		BurnerCodeHashV6:
		break
	default:
		return fmt.Errorf("unsupported burner code with hash %s", burnerCodeHash)
	}

	return nil
}
