package scalarnet

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/scalarorg/scalar-core/utils/events"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/scalarnet/exported"
	"github.com/scalarorg/scalar-core/x/scalarnet/keeper"
	"github.com/scalarorg/scalar-core/x/scalarnet/types"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
)

// NewHandler returns the handler of the Cosmos module
func NewHandler(k keeper.Keeper, n types.Nexus, b types.BankKeeper, ibcK keeper.IBCKeeper) sdk.Handler {
	server := keeper.NewMsgServerImpl(k, n, b, ibcK)
	h := func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.LinkRequest:
			res, err := server.Link(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			if err == nil {
				result.Log = fmt.Sprintf("successfully linked deposit %s to recipient %s", res.DepositAddr, msg.RecipientAddr)
			}
			return result, err
		case *types.ConfirmDepositRequest:
			res, err := server.ConfirmDeposit(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			if err == nil {
				result.Log = fmt.Sprintf("successfully confirmed deposit to {%s}", msg.DepositAddress.String())
			}
			return result, err
		case *types.ExecutePendingTransfersRequest:
			res, err := server.ExecutePendingTransfers(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			if err == nil {
				result.Log = "successfully executed pending transfers"
			}
			return result, err
		case *types.AddCosmosBasedChainRequest:
			res, err := server.AddCosmosBasedChain(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			if err == nil {
				result.Log = fmt.Sprintf("successfully added chain %s", msg.CosmosChain)
			}
			return result, err
		case *types.RegisterAssetRequest:
			res, err := server.RegisterAsset(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			if err == nil {
				result.Log = fmt.Sprintf("successfully registered asset %s to chain %s", msg.Asset.Denom, msg.Chain)
			}
			return result, err
		case *types.RouteIBCTransfersRequest:
			res, err := server.RouteIBCTransfers(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			if err == nil {
				result.Log = "successfully routed IBC transfers"
			}
			return result, err
		case *types.RegisterFeeCollectorRequest:
			res, err := server.RegisterFeeCollector(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			return result, err
		case *types.RetryIBCTransferRequest:
			res, err := server.RetryIBCTransfer(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			return result, err
		case *types.RouteMessageRequest:
			res, err := server.RouteMessage(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			return result, err
		case *types.CallContractRequest:
			res, err := server.CallContract(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			if err == nil {
				result.Log = fmt.Sprintf("successfully enqueued contract call for contract %s on chain %s", msg.ContractAddress, msg.Chain)
			}
			return result, err
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg))
		}
	}

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		res, err := h(ctx, msg)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrScalarnet, err.Error())
		}
		return res, nil
	}
}

// NewProposalHandler returns the handler for the proposals of the scalarnet module
func NewProposalHandler(k keeper.Keeper, nexusK types.Nexus, accountK types.AccountKeeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CallContractsProposal:
			for _, contractCall := range c.ContractCalls {
				sender := nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: accountK.GetModuleAddress(govtypes.ModuleName).String()}

				destChain, ok := nexusK.GetChain(ctx, contractCall.Chain)
				if !ok {
					// Try forwarding it to wasm router if destination chain is not registered
					// Wasm chain names are always lower case, so normalize it for consistency in core
					destChainName := nexus.ChainName(strings.ToLower(contractCall.Chain.String()))
					destChain = nexus.Chain{Name: destChainName, SupportsForeignAssets: false, KeyType: tss.None, Module: wasm.ModuleName}
				}
				recipient := nexus.CrossChainAddress{Chain: destChain, Address: contractCall.ContractAddress}

				//  gateway expects keccak256 hashes for payloads
				payloadHash := crypto.Keccak256(contractCall.Payload)
				msgID, txID, nonce := nexusK.GenerateMessageID(ctx)
				msg := nexus.NewGeneralMessage(msgID, sender, recipient, payloadHash, txID, nonce, nil)

				events.Emit(ctx, &types.ContractCallSubmitted{
					MessageID:        msg.ID,
					Sender:           msg.GetSourceAddress(),
					SourceChain:      msg.GetSourceChain(),
					DestinationChain: msg.GetDestinationChain(),
					ContractAddress:  msg.GetDestinationAddress(),
					PayloadHash:      msg.PayloadHash,
					Payload:          contractCall.Payload,
				})

				if err := nexusK.SetNewMessage(ctx, msg); err != nil {
					return sdkerrors.Wrap(err, "failed to add general message")
				}

				k.Logger(ctx).Debug(fmt.Sprintf("successfully enqueued contract call for contract address %s on chain %s from sender %s with message id %s", recipient.Address, recipient.Chain.String(), sender.Address, msg.ID),
					types.AttributeKeyDestinationChain, recipient.Chain.String(),
					types.AttributeKeyDestinationAddress, recipient.Address,
					types.AttributeKeySourceAddress, sender.Address,
					types.AttributeKeyMessageID, msg.ID,
					types.AttributeKeyPayloadHash, hex.EncodeToString(payloadHash),
				)
			}

			return nil
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized scalarnet proposal content type: %T", c)
		}
	}
}
