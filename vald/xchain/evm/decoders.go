package evm

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	utilsTypes "github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/utils/clog"
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

func (client *EthereumClient) decodeSourceTxConfirmationEvent(log *geth.Log) (*chainsTypes.SourceTxConfirmationEvent, error) {
	arguments := abi.Arguments{
		{Type: stringType},
		{Type: stringType},
		{Type: bytesType},
	}

	params, err := chainsTypes.StrictDecode(arguments, log.Data)
	if err != nil {
		return nil, err
	}

	payloadHash := chainsTypes.Hash(common.BytesToHash(log.Topics[2].Bytes()))
	sender := chainsTypes.Address(common.BytesToAddress(log.Topics[1].Bytes())).Hex()
	destinationChain := nexus.ChainName(params[0].(string))

	clog.Yellowf("[VALD] sender: %s, destinationChain: %s, payloadHash: %s", sender, destinationChain, payloadHash)

	// TODO: Currently, it is the sender of the unstake contract call which is the protocol contract, we need to change the encoding the sender to the encoded payload
	// TODO: Currently, it is the sender of the unstake contract call which is the protocol contract, we need to change the encoding the sender to the encoded payload

	// TODO: Also, we need to change the amount to the encoded payload

	// TODO: We need to calculate the DestinationRecipientAddress from the locking script of the payload following the destinination chain

	// TODO: We also need to encode the symbol of the token to the payload

	// TODO: because the payload was double encoded, we need to decode it again

	// payload = abi.encode(msg.sender, amount, encodedPayload)

	// -> abi.decode(payload, (address, uint256, bytes))

	// get metadata = params[2] as bytes

	// The destination contract address maybe empty

	// return &chainsTypes.TxConfirmationEvent{
	// 	Sender:                      sender, // TODO:Protocol contract address
	// 	DestinationChain:            destinationChain,
	// 	Amount:                      0,          // TODO: Fix hard coded
	// 	Asset:                       "ethereum", // TODO: Fix hard coded
	// 	PayloadHash:                 payloadHash,
	// 	Payload:                     params[2].([]byte),
	// 	DestinationContractAddress:  params[1].(string),
	// 	DestinationRecipientAddress: "",
	// }, nil

	mockEvent, err := genMockEvent()
	if err != nil {
		return nil, err
	}

	clog.Yellowf("[VALD] mockEvent: %+v", mockEvent)

	return mockEvent, nil

}

func genMockEvent() (*chainsTypes.SourceTxConfirmationEvent, error) {
	lockingScript, _ := hex.DecodeString("1234567890123456789012345678901234567890123456789012345678901234567890")
	amount := uint64(7)
	feeOpts := utilsTypes.FastestFee

	payload, hash, err := encode.CalculateUnstakingPayloadHash(lockingScript, amount, feeOpts)
	if err != nil {
		return nil, err
	}

	return &chainsTypes.SourceTxConfirmationEvent{
		Sender:                      "0x24a1dB57Fa3ecAFcbaD91d6Ef068439acEeAe090",
		DestinationChain:            "bitcoin|4",
		Amount:                      100_000,
		Asset:                       "pBTC",
		PayloadHash:                 chainsTypes.Hash(hash),
		Payload:                     payload,
		DestinationContractAddress:  "",
		DestinationRecipientAddress: "tb1q2rwweg2c48y8966qt4fzj0f4zyg9wty7tykzwg",
	}, nil
}
