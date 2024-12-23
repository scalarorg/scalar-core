package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/scalarorg/scalar-core/utils/funcs"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var (
	stringType = funcs.Must(abi.NewType("string", "string", nil))
	bytesType  = funcs.Must(abi.NewType("bytes", "bytes", nil))
)

// Smart contract event signatures
var (
	ContractCallSig = crypto.Keccak256Hash([]byte("ContractCall(address,string,string,bytes32,bytes)"))
)

func (client *EthereumClient) decodeEventContractCall(log *geth.Log) (*chainsTypes.TxConfirmationEvent, error) {
	arguments := abi.Arguments{
		{Type: stringType},
		{Type: stringType},
		{Type: bytesType},
	}

	params, err := chainsTypes.StrictDecode(arguments, log.Data)
	if err != nil {
		return nil, err
	}

	return &chainsTypes.TxConfirmationEvent{
		Sender:           chainsTypes.Address(common.BytesToAddress(log.Topics[1].Bytes())).Hex(),
		DestinationChain: nexus.ChainName(params[0].(string)),
		Amount:           0,          // TODO: Fix hard coded
		Asset:            "ethereum", // TODO: Fix hard coded
		PayloadHash:      chainsTypes.Hash(common.BytesToHash(log.Topics[2].Bytes())),

		// DestinationContractAddress:  chainsTypes.Address(destinationContractAddress),
		// DestinationRecipientAddress: chainsTypes.Address(destinationRecipientAddress),
		DestinationContractAddress:  params[1].(string),
		DestinationRecipientAddress: "",
	}, nil
}
