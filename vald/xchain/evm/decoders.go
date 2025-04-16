package evm

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	covExported "github.com/scalarorg/scalar-core/x/covenant/exported"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var (
	stringType = funcs.Must(abi.NewType("string", "string", nil))
	bytesType  = funcs.Must(abi.NewType("bytes", "bytes", nil))
	uint8Type  = funcs.Must(abi.NewType("uint8", "uint8", nil))
)

// Smart contract event signatures
var (
	ERC20TransferSig                = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	ERC20TokenDeploymentSig         = crypto.Keccak256Hash([]byte("TokenDeployed(string,address)"))
	MultisigTransferOperatorshipSig = crypto.Keccak256Hash([]byte("OperatorshipTransferred(bytes)"))
	ContractCallSig                 = crypto.Keccak256Hash([]byte("ContractCall(address,string,string,bytes32,bytes)"))
	ContractCallWithTokenSig        = crypto.Keccak256Hash([]byte("ContractCallWithToken(address,string,string,bytes32,bytes,string,uint256)"))
	TokenSentSig                    = crypto.Keccak256Hash([]byte("TokenSent(address,string,string,string,uint256)"))

	RedeemTokenSig       = crypto.Keccak256Hash([]byte("RedeemToken(address,uint64,bytes32,string,string,bytes32,bytes,string,uint256)"))
	ContractCallDataArgs = abi.Arguments{
		{Type: stringType},
		{Type: stringType},
		{Type: bytesType},
	}
	SwitchPhaseSig = crypto.Keccak256Hash([]byte("SwitchPhase(bytes32,uint64,uint8,uint8)"))
)

func DecodeERC20TransferEvent(log *geth.Log) (types.EventTransfer, error) {
	if len(log.Topics) != 3 || log.Topics[0] != ERC20TransferSig {
		return types.EventTransfer{}, fmt.Errorf("log is not an ERC20 transfer")
	}

	uint256Type := funcs.Must(abi.NewType("uint256", "uint256", nil))

	to := common.BytesToAddress(log.Topics[2][:])

	arguments := abi.Arguments{
		{Type: uint256Type},
	}

	params, err := arguments.Unpack(log.Data)
	if err != nil {
		return types.EventTransfer{}, err
	}

	return types.EventTransfer{
		To:     types.Address(to),
		Amount: sdk.NewUintFromBigInt(params[0].(*big.Int)),
	}, nil
}

func DecodeERC20TokenDeploymentEvent(log *geth.Log) (types.EventTokenDeployed, error) {
	if len(log.Topics) != 1 || log.Topics[0] != ERC20TokenDeploymentSig {
		return types.EventTokenDeployed{}, fmt.Errorf("event is not for an ERC20 token deployment")
	}

	stringType := funcs.Must(abi.NewType("string", "string", nil))
	addressType := funcs.Must(abi.NewType("address", "address", nil))

	arguments := abi.Arguments{{Type: stringType}, {Type: addressType}}
	params, err := types.StrictDecode(arguments, log.Data)
	if err != nil {
		return types.EventTokenDeployed{}, err
	}

	return types.EventTokenDeployed{
		Symbol:       params[0].(string),
		TokenAddress: types.Address(params[1].(common.Address)),
	}, nil
}

func DecodeMultisigOperatorshipTransferredEvent(log *geth.Log) (types.EventMultisigOperatorshipTransferred, error) {
	if len(log.Topics) != 1 || log.Topics[0] != MultisigTransferOperatorshipSig {
		return types.EventMultisigOperatorshipTransferred{}, fmt.Errorf("event is not OperatorshipTransferred")
	}

	bytesType := funcs.Must(abi.NewType("bytes", "bytes", nil))
	newOperatorsData, err := types.StrictDecode(abi.Arguments{{Type: bytesType}}, log.Data)
	if err != nil {
		return types.EventMultisigOperatorshipTransferred{}, err
	}

	addressesType := funcs.Must(abi.NewType("address[]", "address[]", nil))
	uint256ArrayType := funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))
	uint256Type := funcs.Must(abi.NewType("uint256", "uint256", nil))

	arguments := abi.Arguments{{Type: addressesType}, {Type: uint256ArrayType}, {Type: uint256Type}}
	params, err := types.StrictDecode(arguments, newOperatorsData[0].([]byte))
	if err != nil {
		return types.EventMultisigOperatorshipTransferred{}, err
	}

	event := types.EventMultisigOperatorshipTransferred{
		NewOperators: slices.Map(params[0].([]common.Address), func(addr common.Address) types.Address { return types.Address(addr) }),
		NewWeights:   slices.Map(params[1].([]*big.Int), sdk.NewUintFromBigInt),
		NewThreshold: sdk.NewUintFromBigInt(params[2].(*big.Int)),
	}

	return event, nil
}

func DecodeEventContractCallWithToken(log *geth.Log) (types.EventContractCallWithToken, error) {
	stringType := funcs.Must(abi.NewType("string", "string", nil))
	bytesType := funcs.Must(abi.NewType("bytes", "bytes", nil))
	uint256Type := funcs.Must(abi.NewType("uint256", "uint256", nil))

	arguments := abi.Arguments{
		{Type: stringType},
		{Type: stringType},
		{Type: bytesType},
		{Type: stringType},
		{Type: uint256Type},
	}
	params, err := types.StrictDecode(arguments, log.Data)
	if err != nil {
		clog.Greenf("[Vald/decoders] failed to StrictDecode Data: %+v\n, Argurments: (string, string, bytes, string, uint256)", log.Data)
		return types.EventContractCallWithToken{}, err
	}
	payload, ok := params[2].([]byte)
	if !ok {
		return types.EventContractCallWithToken{}, fmt.Errorf("invalid payload")
	}

	clog.Greenf("[Vald/decoders] tx %x raw payload: %x", log.TxHash, payload)

	_, err = encode.DecodeContractCallWithTokenPayload(payload)
	if err != nil {
		clog.Redf("[Vald/decoders] Payload is invalid, err: %+v\n", err)
		return types.EventContractCallWithToken{}, fmt.Errorf("error decoding contract call payload: %w", err)
	}

	clog.Greenf("[Vald/decoders] Payload is valid")

	return types.EventContractCallWithToken{
		Sender:           types.Address(common.BytesToAddress(log.Topics[1].Bytes())),
		DestinationChain: nexus.ChainName(params[0].(string)),
		ContractAddress:  params[1].(string),
		PayloadHash:      exported.Hash(common.BytesToHash(log.Topics[2].Bytes())),
		Symbol:           params[3].(string),
		Amount:           sdk.NewUintFromBigInt(params[4].(*big.Int)),
		Payload:          payload,
	}, nil
}

func DecodeEventTokenSent(log *geth.Log) (types.EventTokenSent, error) {
	stringType := funcs.Must(abi.NewType("string", "string", nil))
	uint256Type := funcs.Must(abi.NewType("uint256", "uint256", nil))

	arguments := abi.Arguments{
		{Type: stringType},
		{Type: stringType},
		{Type: stringType},
		{Type: uint256Type},
	}
	params, err := types.StrictDecode(arguments, log.Data)
	if err != nil {
		return types.EventTokenSent{}, err
	}

	return types.EventTokenSent{
		Sender:             types.Address(common.BytesToAddress(log.Topics[1].Bytes())).String(),
		DestinationChain:   nexus.ChainName(params[0].(string)),
		DestinationAddress: params[1].(string),
		Asset:              sdk.NewCoin(params[2].(string), sdk.NewIntFromBigInt(params[3].(*big.Int))),
	}, nil
}

func DecodeEventContractCall(log *geth.Log) (types.EventContractCall, error) {
	stringType := funcs.Must(abi.NewType("string", "string", nil))
	bytesType := funcs.Must(abi.NewType("bytes", "bytes", nil))

	arguments := abi.Arguments{
		{Type: stringType},
		{Type: stringType},
		{Type: bytesType},
	}
	params, err := types.StrictDecode(arguments, log.Data)
	if err != nil {
		return types.EventContractCall{}, err
	}

	return types.EventContractCall{
		Sender:           types.Address(common.BytesToAddress(log.Topics[1].Bytes())),
		DestinationChain: nexus.ChainName(params[0].(string)),
		ContractAddress:  params[1].(string),
		PayloadHash:      exported.Hash(common.BytesToHash(log.Topics[2].Bytes())),
	}, nil
}

func DecodeEventSwitchPhase(log *geth.Log) (*covTypes.SwitchedPhaseConfirmed, error) {
	if len(log.Topics) < 3 || log.Topics[0] != SwitchPhaseSig {
		return nil, fmt.Errorf("event is not SwitchPhase")
	}

	// event SwitchPhase(bytes32 indexed custodianGroupId, uint64 indexed sequence, Phase from, Phase to);

	custodianGroupId := common.BytesToHash(log.Topics[1].Bytes())
	sequence := sdk.NewUintFromBigInt(new(big.Int).SetBytes(log.Topics[2].Bytes()))
	// enumType := funcs.Must(abi.NewType("enum Phase", "enum Phase", []abi.ArgumentMarshaling{
	// 	{Name: "Preparing", Type: "uint8"},
	// 	{Name: "Executing", Type: "uint8"},
	// }))
	// arguments := abi.Arguments{
	// 	{Type: enumType},
	// 	{Type: enumType},
	// }
	arguments := abi.Arguments{
		{Type: uint8Type},
		{Type: uint8Type},
	}
	params, err := types.StrictDecode(arguments, log.Data)
	if err != nil {
		return nil, err
	}

	from := params[0].(uint8)
	to := params[1].(uint8)

	// We can switch from Preparing to Executing and vice versa, so condition is from != to
	// if from == to {
	// 	return nil, fmt.Errorf("invalid phase transition")
	// }

	return &covTypes.SwitchedPhaseConfirmed{
		CustodianGroupUID: exported.Hash(custodianGroupId),
		Sequence:          sequence.Uint64(),
		FromPhase:         covExported.Phase(from),
		ToPhase:           covExported.Phase(to),
	}, nil

}

func DecodeEventRedeemToken(log *geth.Log) (types.EventRedeemToken, error) {
	stringType := funcs.Must(abi.NewType("string", "string", nil))
	bytesType := funcs.Must(abi.NewType("bytes", "bytes", nil))
	uint256Type := funcs.Must(abi.NewType("uint256", "uint256", nil))
	bytes32Type := funcs.Must(abi.NewType("bytes32", "bytes32", nil))

	arguments := abi.Arguments{
		{Type: bytes32Type},
		{Type: stringType},
		{Type: stringType},
		{Type: bytesType},
		{Type: stringType},
		{Type: uint256Type},
	}
	params, err := types.StrictDecode(arguments, log.Data)
	if err != nil {
		clog.Greenf("[Vald/decoders] failed to StrictDecode Data: %+v\n, Argurments: (string, string, bytes, string, uint256)", log.Data)
		return types.EventRedeemToken{}, err
	}
	payload, ok := params[3].([]byte)
	if !ok {
		return types.EventRedeemToken{}, fmt.Errorf("invalid payload")
	}

	clog.Greenf("[Vald/decoders] tx %x raw payload: %x", log.TxHash, payload)

	// this function is used to check if the payload is valid or not
	_, err = encode.DecodeContractCallWithTokenPayload(payload)
	if err != nil {
		clog.Redf("[Vald/decoders] Payload is invalid, err: %+v\n", err)
		return types.EventRedeemToken{}, fmt.Errorf("error decoding contract call payload: %w", err)
	}

	clog.Greenf("[Vald/decoders] Payload is valid")

	return types.EventRedeemToken{
		Sender:                     types.Address(common.BytesToAddress(log.Topics[1].Bytes())),
		Sequence:                   log.Topics[2].Big().Uint64(),
		CustodianGroupId:           exported.Hash(params[0].([32]byte)),
		DestinationChain:           nexus.ChainName(params[1].(string)),
		DestinationContractAddress: params[2].(string),
		PayloadHash:                exported.Hash(common.BytesToHash(log.Topics[3].Bytes())),
		Symbol:                     params[4].(string),
		Amount:                     sdk.NewUintFromBigInt(params[5].(*big.Int)),
		Payload:                    payload,
	}, nil
}
