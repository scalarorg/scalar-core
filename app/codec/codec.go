package codec

import (
	"fmt"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexustypes "github.com/scalarorg/scalar-core/x/nexus/types"
	permissiontypes "github.com/scalarorg/scalar-core/x/permission/types"
	rewardtypes "github.com/scalarorg/scalar-core/x/reward/types"
	scalarnettypes "github.com/scalarorg/scalar-core/x/scalarnet/types"
	snapshottypes "github.com/scalarorg/scalar-core/x/snapshot/types"
	tsstypes "github.com/scalarorg/scalar-core/x/tss/types"
	votetypes "github.com/scalarorg/scalar-core/x/vote/types"
)

type customRegistry interface {
	RegisterCustomTypeURL(iface interface{}, typeURL string, impl proto.Message)
}

// RegisterLegacyMsgInterfaces registers the msg codec before the package name
// refactor done in https://github.com/scalarorg/scalar-core/commit/2d5e35d7da4fb02ac55fb040fed420954d3be020
// to keep transaction query backwards compatible
func RegisterLegacyMsgInterfaces(registry cdctypes.InterfaceRegistry) {
	r, ok := registry.(customRegistry)
	if !ok {
		panic(fmt.Errorf("failed to convert registry type %T", registry))
	}

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".evm.v1beta1.CreateTransferOwnershipRequest", &chainsTypes.CreateTransferOwnershipRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.RegisterIBCPathRequest", &scalarnettypes.RegisterIBCPathRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.LinkRequest", &scalarnettypes.LinkRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.ConfirmDepositRequest", &scalarnettypes.ConfirmDepositRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.ExecutePendingTransfersRequest", &scalarnettypes.ExecutePendingTransfersRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.RegisterIBCPathRequest", &scalarnettypes.RegisterIBCPathRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.AddCosmosBasedChainRequest", &scalarnettypes.AddCosmosBasedChainRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.RegisterAssetRequest", &scalarnettypes.RegisterAssetRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.RouteIBCTransfersRequest", &scalarnettypes.RouteIBCTransfersRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/scalar.scalarnet.v1beta1.RegisterFeeCollectorRequest", &scalarnettypes.RegisterFeeCollectorRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.LinkRequest", &chainsTypes.LinkRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.ConfirmTokenRequest", &chainsTypes.ConfirmTokenRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.ConfirmDepositRequest", &chainsTypes.ConfirmDepositRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.ConfirmTransferKeyRequest", &chainsTypes.ConfirmTransferKeyRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.CreatePendingTransfersRequest", &chainsTypes.CreatePendingTransfersRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.CreateDeployTokenRequest", &chainsTypes.CreateDeployTokenRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.CreateBurnTokensRequest", &chainsTypes.CreateBurnTokensRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.CreateTransferOwnershipRequest", &chainsTypes.CreateTransferOwnershipRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.CreateTransferOperatorshipRequest", &chainsTypes.CreateTransferOperatorshipRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.SignCommandsRequest", &chainsTypes.SignCommandsRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.AddChainRequest", &chainsTypes.AddChainRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.SetGatewayRequest", &chainsTypes.SetGatewayRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.ConfirmSourceTxsRequest", &chainsTypes.ConfirmSourceTxsRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/chains.v1beta1.RetryFailedEventRequest", &chainsTypes.RetryFailedEventRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/nexus.v1beta1.RegisterChainMaintainerRequest", &nexustypes.RegisterChainMaintainerRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/nexus.v1beta1.DeregisterChainMaintainerRequest", &nexustypes.DeregisterChainMaintainerRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/nexus.v1beta1.ActivateChainRequest", &nexustypes.ActivateChainRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/nexus.v1beta1.DeactivateChainRequest", &nexustypes.DeactivateChainRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/nexus.v1beta1.RegisterAssetFeeRequest", &nexustypes.RegisterAssetFeeRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/permission.v1beta1.UpdateGovernanceKeyRequest", &permissiontypes.UpdateGovernanceKeyRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/permission.v1beta1.RegisterControllerRequest", &permissiontypes.RegisterControllerRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/permission.v1beta1.DeregisterControllerRequest", &permissiontypes.DeregisterControllerRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/reward.v1beta1.RefundMsgRequest", &rewardtypes.RefundMsgRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/snapshot.v1beta1.RegisterProxyRequest", &snapshottypes.RegisterProxyRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/snapshot.v1beta1.DeactivateProxyRequest", &snapshottypes.DeactivateProxyRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.HeartBeatRequest", &tsstypes.HeartBeatRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.StartKeygenRequest", &tsstypes.StartKeygenRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.ProcessKeygenTrafficRequest", &tsstypes.ProcessKeygenTrafficRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.ProcessSignTrafficRequest", &tsstypes.ProcessSignTrafficRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.RotateKeyRequest", &tsstypes.RotateKeyRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.VoteSigRequest", &tsstypes.VoteSigRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.VotePubKeyRequest", &tsstypes.VotePubKeyRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.RegisterExternalKeysRequest", &tsstypes.RegisterExternalKeysRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.SubmitMultisigPubKeysRequest", &tsstypes.SubmitMultisigPubKeysRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/tss.v1beta1.SubmitMultisigSignaturesRequest", &tsstypes.SubmitMultisigSignaturesRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.StartKeygenRequest", &tsstypes.StartKeygenRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.ProcessKeygenTrafficRequest", &tsstypes.ProcessKeygenTrafficRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.ProcessSignTrafficRequest", &tsstypes.ProcessSignTrafficRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.RotateKeyRequest", &tsstypes.RotateKeyRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.VoteSigRequest", &tsstypes.VoteSigRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.VotePubKeyRequest", &tsstypes.VotePubKeyRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.RegisterExternalKeysRequest", &tsstypes.RegisterExternalKeysRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.SubmitMultisigPubKeysRequest", &tsstypes.SubmitMultisigPubKeysRequest{})
	r.RegisterCustomTypeURL((*sdk.Msg)(nil), ".tss.v1beta1.SubmitMultisigSignaturesRequest", &tsstypes.SubmitMultisigSignaturesRequest{})

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/vote.v1beta1.VoteRequest", &votetypes.VoteRequest{})
}
