package evm

import (
	"context"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/scalarorg/bitcoin-vault/go-utils/address"
	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	grpc_client "github.com/scalarorg/scalar-core/vald/grpc-client"
	"github.com/scalarorg/scalar-core/x/chains/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var (
	stringType = funcs.Must(abi.NewType("string", "string", nil))
	bytesType  = funcs.Must(abi.NewType("bytes", "bytes", nil))
)

// Smart contract event signatures
var (
	ERC20TransferSig                = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	ERC20TokenDeploymentSig         = crypto.Keccak256Hash([]byte("TokenDeployed(string,address)"))
	MultisigTransferOperatorshipSig = crypto.Keccak256Hash([]byte("OperatorshipTransferred(bytes)"))
	ContractCallSig                 = crypto.Keccak256Hash([]byte("ContractCall(address,string,string,bytes32,bytes)"))
	ContractCallWithTokenSig        = crypto.Keccak256Hash([]byte("ContractCallWithToken(address,string,string,bytes32,bytes,string,uint256)"))
	TokenSentSig                    = crypto.Keccak256Hash([]byte("TokenSent(address,string,string,string,uint256)"))
	ContractCallDataArgs            = abi.Arguments{
		{Type: stringType},
		{Type: stringType},
		{Type: bytesType},
	}
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
		return types.EventContractCallWithToken{}, err
	}

	return types.EventContractCallWithToken{
		Sender:           types.Address(common.BytesToAddress(log.Topics[1].Bytes())),
		DestinationChain: nexus.ChainName(params[0].(string)),
		ContractAddress:  params[1].(string),
		PayloadHash:      types.Hash(common.BytesToHash(log.Topics[2].Bytes())),
		Symbol:           params[3].(string),
		Amount:           sdk.NewUintFromBigInt(params[4].(*big.Int)),
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
		PayloadHash:      types.Hash(common.BytesToHash(log.Topics[2].Bytes())),
	}, nil
}

func (client *EthereumClient) decodeSourceTxConfirmationEvent(log *geth.Log) (*chainsTypes.SourceTxConfirmationEvent, error) {
	params, err := chainsTypes.StrictDecode(ContractCallDataArgs, log.Data)
	if err != nil {
		return nil, err
	}

	payload, ok := params[2].([]byte)
	if !ok {
		return nil, fmt.Errorf("invalid payload")
	}

	sender, _, symbol, metadata, err := encode.DecodeTransferRemotePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("error decoding transfer remote payload: %w", err)
	}

	amount, recipientChainIdentifier, _, err := encode.DecodeTransferRemoteMetadataPayload(metadata)
	if err != nil {
		return nil, fmt.Errorf("error decoding transfer remote metadata payload: %w", err)
	}

	chainID := params[0].(string)
	if !utils.ValidateChainID(chainID) {
		return nil, fmt.Errorf("invalid chain id")
	}

	destinationChain := nexus.ChainName(chainID)

	destinationContractAddress := common.HexToAddress(params[1].(string)).Hex()

	payloadHash := chainsTypes.Hash(common.BytesToHash(log.Topics[2].Bytes()))

	queryClient := grpc_client.QueryManager.GetClient()

	chainParams, err := queryClient.Params(context.Background(), &chainsTypes.ParamsRequest{
		Chain: chainID,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting chain metadata: %w", err)
	}

	chainMetadata := chainParams.Params.Metadata

	destinationRecipientAddress, err := decodeAddress(destinationChain, recipientChainIdentifier, chainMetadata)
	if err != nil {
		return nil, fmt.Errorf("error decoding destination recipient address: %w", err)
	}

	cfEvent := &chainsTypes.SourceTxConfirmationEvent{
		Sender:                      sender.Hex(),
		DestinationChain:            destinationChain,
		Amount:                      amount,
		Asset:                       symbol,
		PayloadHash:                 payloadHash,
		Payload:                     payload,
		DestinationContractAddress:  destinationContractAddress,
		DestinationRecipientAddress: destinationRecipientAddress,
	}

	clog.Greenf("decoded event: %+v", cfEvent)

	return cfEvent, nil
}

func decodeAddress(chain nexus.ChainName, identifier []byte, metadata map[string]string) (string, error) {
	if chainsTypes.IsBitcoinChain(chain) {
		params := metadata["params"]
		if params == "" {
			return "", fmt.Errorf("params is required")
		}

		addr, err := address.ScriptPubKeyToAddress(identifier, params)
		if err != nil {
			return "", fmt.Errorf("error decoding address: %w", err)
		}
		return addr.String(), nil
	}

	if chainsTypes.IsEvmChain(chain) {
		address := common.BytesToAddress(identifier)
		return address.String(), nil
	}

	return "", fmt.Errorf("chain not supported")
}
