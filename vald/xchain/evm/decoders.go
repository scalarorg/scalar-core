package evm

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/scalarorg/bitcoin-vault/go-utils/address"
	btcChain "github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/funcs"
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
	ContractCallSig      = crypto.Keccak256Hash([]byte("ContractCall(address,string,string,bytes32,bytes)"))
	ContractCallDataArgs = abi.Arguments{
		{Type: stringType},
		{Type: stringType},
		{Type: bytesType},
	}
)

func (client *EthereumClient) decodeSourceTxConfirmationEvent(event *types.EventConfirmSourceTxsStarted, log *geth.Log) (*chainsTypes.SourceTxConfirmationEvent, error) {
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

	destinationRecipientAddress, err := decodeAddress(chainID, recipientChainIdentifier, chainMetadata)
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

func decodeAddress(chain string, identifier []byte, metadata map[string]string) (string, error) {
	chainInfoBytes, err := utils.ChainInfoBytesFromString(chain)
	if err != nil {
		return "", fmt.Errorf("error decoding chain info bytes: %w", err)
	}

	chainType := chainInfoBytes.ChainType()

	if chainType == btcChain.ChainTypeBitcoin {
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

	if chainType == btcChain.ChainTypeEVM {
		address := common.BytesToAddress(identifier)
		return address.String(), nil
	}

	return "", fmt.Errorf("chain not supported")
}
