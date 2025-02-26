<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [scalar/auxiliary/v1beta1/events.proto](#scalar/auxiliary/v1beta1/events.proto)
    - [BatchedMessageFailed](#scalar.auxiliary.v1beta1.BatchedMessageFailed)
  
- [scalar/auxiliary/v1beta1/genesis.proto](#scalar/auxiliary/v1beta1/genesis.proto)
    - [GenesisState](#scalar.auxiliary.v1beta1.GenesisState)
  
- [scalar/permission/exported/v1beta1/types.proto](#scalar/permission/exported/v1beta1/types.proto)
    - [Role](#scalar.permission.exported.v1beta1.Role)
  
    - [File-level Extensions](#scalar/permission/exported/v1beta1/types.proto-extensions)
  
- [scalar/auxiliary/v1beta1/tx.proto](#scalar/auxiliary/v1beta1/tx.proto)
    - [BatchRequest](#scalar.auxiliary.v1beta1.BatchRequest)
    - [BatchResponse](#scalar.auxiliary.v1beta1.BatchResponse)
    - [BatchResponse.Response](#scalar.auxiliary.v1beta1.BatchResponse.Response)
  
- [scalar/auxiliary/v1beta1/service.proto](#scalar/auxiliary/v1beta1/service.proto)
    - [MsgService](#scalar.auxiliary.v1beta1.MsgService)
  
- [scalar/chains/btc/v1beta1/types.proto](#scalar/chains/btc/v1beta1/types.proto)
    - [BtcToken](#scalar.chains.btc.v1beta1.BtcToken)
  
- [scalar/chains/v1beta1/types.proto](#scalar/chains/v1beta1/types.proto)
    - [Asset](#scalar.chains.v1beta1.Asset)
    - [BurnerInfo](#scalar.chains.v1beta1.BurnerInfo)
    - [Command](#scalar.chains.v1beta1.Command)
    - [CommandBatchMetadata](#scalar.chains.v1beta1.CommandBatchMetadata)
    - [Gateway](#scalar.chains.v1beta1.Gateway)
    - [PollCompleted](#scalar.chains.v1beta1.PollCompleted)
    - [PollExpired](#scalar.chains.v1beta1.PollExpired)
    - [PollFailed](#scalar.chains.v1beta1.PollFailed)
    - [PollMapping](#scalar.chains.v1beta1.PollMapping)
    - [PollMetadata](#scalar.chains.v1beta1.PollMetadata)
    - [Proof](#scalar.chains.v1beta1.Proof)
    - [SigMetadata](#scalar.chains.v1beta1.SigMetadata)
    - [SourceTx](#scalar.chains.v1beta1.SourceTx)
    - [TokenDetails](#scalar.chains.v1beta1.TokenDetails)
    - [TransferKey](#scalar.chains.v1beta1.TransferKey)
  
    - [BatchedCommandsStatus](#scalar.chains.v1beta1.BatchedCommandsStatus)
    - [CommandType](#scalar.chains.v1beta1.CommandType)
    - [DepositStatus](#scalar.chains.v1beta1.DepositStatus)
    - [SigType](#scalar.chains.v1beta1.SigType)
    - [SourceTxStatus](#scalar.chains.v1beta1.SourceTxStatus)
    - [Status](#scalar.chains.v1beta1.Status)
  
- [scalar/utils/v1beta1/threshold.proto](#scalar/utils/v1beta1/threshold.proto)
    - [Threshold](#scalar.utils.v1beta1.Threshold)
  
- [scalar/tss/exported/v1beta1/types.proto](#scalar/tss/exported/v1beta1/types.proto)
    - [KeyRequirement](#scalar.tss.exported.v1beta1.KeyRequirement)
    - [SigKeyPair](#scalar.tss.exported.v1beta1.SigKeyPair)
  
    - [KeyRole](#scalar.tss.exported.v1beta1.KeyRole)
    - [KeyShareDistributionPolicy](#scalar.tss.exported.v1beta1.KeyShareDistributionPolicy)
    - [KeyType](#scalar.tss.exported.v1beta1.KeyType)
  
- [scalar/snapshot/exported/v1beta1/types.proto](#scalar/snapshot/exported/v1beta1/types.proto)
    - [Participant](#scalar.snapshot.exported.v1beta1.Participant)
    - [Snapshot](#scalar.snapshot.exported.v1beta1.Snapshot)
    - [Snapshot.ParticipantsEntry](#scalar.snapshot.exported.v1beta1.Snapshot.ParticipantsEntry)
  
- [scalar/vote/exported/v1beta1/types.proto](#scalar/vote/exported/v1beta1/types.proto)
    - [PollKey](#scalar.vote.exported.v1beta1.PollKey)
    - [PollMetadata](#scalar.vote.exported.v1beta1.PollMetadata)
    - [PollParticipants](#scalar.vote.exported.v1beta1.PollParticipants)
  
    - [PollState](#scalar.vote.exported.v1beta1.PollState)
  
- [scalar/chains/v1beta1/events.proto](#scalar/chains/v1beta1/events.proto)
    - [BurnCommand](#scalar.chains.v1beta1.BurnCommand)
    - [ChainAdded](#scalar.chains.v1beta1.ChainAdded)
    - [ChainEventCompleted](#scalar.chains.v1beta1.ChainEventCompleted)
    - [ChainEventConfirmed](#scalar.chains.v1beta1.ChainEventConfirmed)
    - [ChainEventFailed](#scalar.chains.v1beta1.ChainEventFailed)
    - [ChainEventRetryFailed](#scalar.chains.v1beta1.ChainEventRetryFailed)
    - [CommandBatchAborted](#scalar.chains.v1beta1.CommandBatchAborted)
    - [CommandBatchSigned](#scalar.chains.v1beta1.CommandBatchSigned)
    - [ConfirmDepositStarted](#scalar.chains.v1beta1.ConfirmDepositStarted)
    - [ConfirmKeyTransferStarted](#scalar.chains.v1beta1.ConfirmKeyTransferStarted)
    - [ConfirmTokenStarted](#scalar.chains.v1beta1.ConfirmTokenStarted)
    - [ContractCallApproved](#scalar.chains.v1beta1.ContractCallApproved)
    - [ContractCallFailed](#scalar.chains.v1beta1.ContractCallFailed)
    - [Event](#scalar.chains.v1beta1.Event)
    - [EventConfirmSourceTxsStarted](#scalar.chains.v1beta1.EventConfirmSourceTxsStarted)
    - [EventContractCall](#scalar.chains.v1beta1.EventContractCall)
    - [EventContractCallWithMintApproved](#scalar.chains.v1beta1.EventContractCallWithMintApproved)
    - [EventContractCallWithToken](#scalar.chains.v1beta1.EventContractCallWithToken)
    - [EventMultisigOperatorshipTransferred](#scalar.chains.v1beta1.EventMultisigOperatorshipTransferred)
    - [EventMultisigOwnershipTransferred](#scalar.chains.v1beta1.EventMultisigOwnershipTransferred)
    - [EventTokenDeployed](#scalar.chains.v1beta1.EventTokenDeployed)
    - [EventTokenSent](#scalar.chains.v1beta1.EventTokenSent)
    - [EventTransfer](#scalar.chains.v1beta1.EventTransfer)
    - [MintCommand](#scalar.chains.v1beta1.MintCommand)
    - [NoEventsConfirmed](#scalar.chains.v1beta1.NoEventsConfirmed)
    - [SourceTxConfirmationEvent](#scalar.chains.v1beta1.SourceTxConfirmationEvent)
    - [VoteEvents](#scalar.chains.v1beta1.VoteEvents)
  
    - [Event.Status](#scalar.chains.v1beta1.Event.Status)
  
- [scalar/chains/v1beta1/params.proto](#scalar/chains/v1beta1/params.proto)
    - [Params](#scalar.chains.v1beta1.Params)
    - [Params.MetadataEntry](#scalar.chains.v1beta1.Params.MetadataEntry)
  
- [scalar/chains/v1beta1/tokens.proto](#scalar/chains/v1beta1/tokens.proto)
    - [ERC20Deposit](#scalar.chains.v1beta1.ERC20Deposit)
    - [ERC20TokenMetadata](#scalar.chains.v1beta1.ERC20TokenMetadata)
  
- [scalar/utils/v1beta1/queuer.proto](#scalar/utils/v1beta1/queuer.proto)
    - [QueueState](#scalar.utils.v1beta1.QueueState)
    - [QueueState.Item](#scalar.utils.v1beta1.QueueState.Item)
    - [QueueState.ItemsEntry](#scalar.utils.v1beta1.QueueState.ItemsEntry)
  
- [scalar/chains/v1beta1/genesis.proto](#scalar/chains/v1beta1/genesis.proto)
    - [GenesisState](#scalar.chains.v1beta1.GenesisState)
    - [GenesisState.Chain](#scalar.chains.v1beta1.GenesisState.Chain)
  
- [scalar/chains/v1beta1/tx.proto](#scalar/chains/v1beta1/tx.proto)
    - [AddChainRequest](#scalar.chains.v1beta1.AddChainRequest)
    - [AddChainResponse](#scalar.chains.v1beta1.AddChainResponse)
    - [ConfirmDepositRequest](#scalar.chains.v1beta1.ConfirmDepositRequest)
    - [ConfirmDepositResponse](#scalar.chains.v1beta1.ConfirmDepositResponse)
    - [ConfirmSourceTxsRequest](#scalar.chains.v1beta1.ConfirmSourceTxsRequest)
    - [ConfirmSourceTxsResponse](#scalar.chains.v1beta1.ConfirmSourceTxsResponse)
    - [ConfirmTokenRequest](#scalar.chains.v1beta1.ConfirmTokenRequest)
    - [ConfirmTokenResponse](#scalar.chains.v1beta1.ConfirmTokenResponse)
    - [ConfirmTransferKeyRequest](#scalar.chains.v1beta1.ConfirmTransferKeyRequest)
    - [ConfirmTransferKeyResponse](#scalar.chains.v1beta1.ConfirmTransferKeyResponse)
    - [CreateBurnTokensRequest](#scalar.chains.v1beta1.CreateBurnTokensRequest)
    - [CreateBurnTokensResponse](#scalar.chains.v1beta1.CreateBurnTokensResponse)
    - [CreateDeployTokenRequest](#scalar.chains.v1beta1.CreateDeployTokenRequest)
    - [CreateDeployTokenResponse](#scalar.chains.v1beta1.CreateDeployTokenResponse)
    - [CreatePendingTransfersRequest](#scalar.chains.v1beta1.CreatePendingTransfersRequest)
    - [CreatePendingTransfersResponse](#scalar.chains.v1beta1.CreatePendingTransfersResponse)
    - [CreateTransferOperatorshipRequest](#scalar.chains.v1beta1.CreateTransferOperatorshipRequest)
    - [CreateTransferOperatorshipResponse](#scalar.chains.v1beta1.CreateTransferOperatorshipResponse)
    - [CreateTransferOwnershipRequest](#scalar.chains.v1beta1.CreateTransferOwnershipRequest)
    - [CreateTransferOwnershipResponse](#scalar.chains.v1beta1.CreateTransferOwnershipResponse)
    - [LinkRequest](#scalar.chains.v1beta1.LinkRequest)
    - [LinkResponse](#scalar.chains.v1beta1.LinkResponse)
    - [RetryFailedEventRequest](#scalar.chains.v1beta1.RetryFailedEventRequest)
    - [RetryFailedEventResponse](#scalar.chains.v1beta1.RetryFailedEventResponse)
    - [SetGatewayRequest](#scalar.chains.v1beta1.SetGatewayRequest)
    - [SetGatewayResponse](#scalar.chains.v1beta1.SetGatewayResponse)
    - [SignBtcCommandsRequest](#scalar.chains.v1beta1.SignBtcCommandsRequest)
    - [SignCommandsRequest](#scalar.chains.v1beta1.SignCommandsRequest)
    - [SignCommandsResponse](#scalar.chains.v1beta1.SignCommandsResponse)
    - [SignPsbtCommandRequest](#scalar.chains.v1beta1.SignPsbtCommandRequest)
    - [SignPsbtCommandResponse](#scalar.chains.v1beta1.SignPsbtCommandResponse)
  
- [scalar/covenant/exported/v1beta1/types.proto](#scalar/covenant/exported/v1beta1/types.proto)
    - [ListOfTapScriptSigsMap](#scalar.covenant.exported.v1beta1.ListOfTapScriptSigsMap)
    - [TapScriptSig](#scalar.covenant.exported.v1beta1.TapScriptSig)
    - [TapScriptSigsEntry](#scalar.covenant.exported.v1beta1.TapScriptSigsEntry)
    - [TapScriptSigsList](#scalar.covenant.exported.v1beta1.TapScriptSigsList)
    - [TapScriptSigsMap](#scalar.covenant.exported.v1beta1.TapScriptSigsMap)
  
    - [KeyState](#scalar.covenant.exported.v1beta1.KeyState)
    - [PsbtState](#scalar.covenant.exported.v1beta1.PsbtState)
  
- [scalar/multisig/exported/v1beta1/types.proto](#scalar/multisig/exported/v1beta1/types.proto)
    - [KeyState](#scalar.multisig.exported.v1beta1.KeyState)
    - [MultisigState](#scalar.multisig.exported.v1beta1.MultisigState)
  
- [scalar/multisig/v1beta1/types.proto](#scalar/multisig/v1beta1/types.proto)
    - [Key](#scalar.multisig.v1beta1.Key)
    - [Key.PubKeysEntry](#scalar.multisig.v1beta1.Key.PubKeysEntry)
    - [KeyEpoch](#scalar.multisig.v1beta1.KeyEpoch)
    - [KeygenSession](#scalar.multisig.v1beta1.KeygenSession)
    - [KeygenSession.IsPubKeyReceivedEntry](#scalar.multisig.v1beta1.KeygenSession.IsPubKeyReceivedEntry)
    - [MultiSig](#scalar.multisig.v1beta1.MultiSig)
    - [MultiSig.SigsEntry](#scalar.multisig.v1beta1.MultiSig.SigsEntry)
    - [SigningSession](#scalar.multisig.v1beta1.SigningSession)
  
- [scalar/covenant/v1beta1/types.proto](#scalar/covenant/v1beta1/types.proto)
    - [Custodian](#scalar.covenant.v1beta1.Custodian)
    - [CustodianGroup](#scalar.covenant.v1beta1.CustodianGroup)
    - [PsbtMultiSig](#scalar.covenant.v1beta1.PsbtMultiSig)
    - [PsbtMultiSig.ParticipantListTapScriptSigsEntry](#scalar.covenant.v1beta1.PsbtMultiSig.ParticipantListTapScriptSigsEntry)
    - [SigningSession](#scalar.covenant.v1beta1.SigningSession)
  
    - [Status](#scalar.covenant.v1beta1.Status)
  
- [scalar/protocol/exported/v1beta1/types.proto](#scalar/protocol/exported/v1beta1/types.proto)
    - [MinorAddress](#scalar.protocol.exported.v1beta1.MinorAddress)
    - [ProtocolAttributes](#scalar.protocol.exported.v1beta1.ProtocolAttributes)
    - [ProtocolInfo](#scalar.protocol.exported.v1beta1.ProtocolInfo)
    - [SupportedChain](#scalar.protocol.exported.v1beta1.SupportedChain)
  
    - [LiquidityModel](#scalar.protocol.exported.v1beta1.LiquidityModel)
    - [Status](#scalar.protocol.exported.v1beta1.Status)
  
- [scalar/chains/v1beta1/query.proto](#scalar/chains/v1beta1/query.proto)
    - [BatchedCommandsRequest](#scalar.chains.v1beta1.BatchedCommandsRequest)
    - [BatchedCommandsResponse](#scalar.chains.v1beta1.BatchedCommandsResponse)
    - [BurnerInfoRequest](#scalar.chains.v1beta1.BurnerInfoRequest)
    - [BurnerInfoResponse](#scalar.chains.v1beta1.BurnerInfoResponse)
    - [BytecodeRequest](#scalar.chains.v1beta1.BytecodeRequest)
    - [BytecodeResponse](#scalar.chains.v1beta1.BytecodeResponse)
    - [ChainsRequest](#scalar.chains.v1beta1.ChainsRequest)
    - [ChainsResponse](#scalar.chains.v1beta1.ChainsResponse)
    - [CommandRequest](#scalar.chains.v1beta1.CommandRequest)
    - [CommandResponse](#scalar.chains.v1beta1.CommandResponse)
    - [CommandResponse.ParamsEntry](#scalar.chains.v1beta1.CommandResponse.ParamsEntry)
    - [ConfirmationHeightRequest](#scalar.chains.v1beta1.ConfirmationHeightRequest)
    - [ConfirmationHeightResponse](#scalar.chains.v1beta1.ConfirmationHeightResponse)
    - [DepositStateRequest](#scalar.chains.v1beta1.DepositStateRequest)
    - [DepositStateResponse](#scalar.chains.v1beta1.DepositStateResponse)
    - [ERC20TokensRequest](#scalar.chains.v1beta1.ERC20TokensRequest)
    - [ERC20TokensResponse](#scalar.chains.v1beta1.ERC20TokensResponse)
    - [ERC20TokensResponse.Token](#scalar.chains.v1beta1.ERC20TokensResponse.Token)
    - [EventRequest](#scalar.chains.v1beta1.EventRequest)
    - [EventResponse](#scalar.chains.v1beta1.EventResponse)
    - [GatewayAddressRequest](#scalar.chains.v1beta1.GatewayAddressRequest)
    - [GatewayAddressResponse](#scalar.chains.v1beta1.GatewayAddressResponse)
    - [KeyAddressRequest](#scalar.chains.v1beta1.KeyAddressRequest)
    - [KeyAddressResponse](#scalar.chains.v1beta1.KeyAddressResponse)
    - [KeyAddressResponse.WeightedAddress](#scalar.chains.v1beta1.KeyAddressResponse.WeightedAddress)
    - [ParamsRequest](#scalar.chains.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.chains.v1beta1.ParamsResponse)
    - [PendingCommandsRequest](#scalar.chains.v1beta1.PendingCommandsRequest)
    - [PendingCommandsResponse](#scalar.chains.v1beta1.PendingCommandsResponse)
    - [QueryBurnerAddressResponse](#scalar.chains.v1beta1.QueryBurnerAddressResponse)
    - [QueryCommandResponse](#scalar.chains.v1beta1.QueryCommandResponse)
    - [QueryCommandResponse.ParamsEntry](#scalar.chains.v1beta1.QueryCommandResponse.ParamsEntry)
    - [QueryDepositStateParams](#scalar.chains.v1beta1.QueryDepositStateParams)
    - [QueryTokenAddressResponse](#scalar.chains.v1beta1.QueryTokenAddressResponse)
    - [TokenInfoRequest](#scalar.chains.v1beta1.TokenInfoRequest)
    - [TokenInfoResponse](#scalar.chains.v1beta1.TokenInfoResponse)
  
    - [ChainStatus](#scalar.chains.v1beta1.ChainStatus)
    - [TokenType](#scalar.chains.v1beta1.TokenType)
  
- [scalar/chains/v1beta1/service.proto](#scalar/chains/v1beta1/service.proto)
    - [MsgService](#scalar.chains.v1beta1.MsgService)
    - [QueryService](#scalar.chains.v1beta1.QueryService)
  
- [scalar/covenant/v1beta1/events.proto](#scalar/covenant/v1beta1/events.proto)
    - [KeyRotated](#scalar.covenant.v1beta1.KeyRotated)
    - [SigningPsbtCompleted](#scalar.covenant.v1beta1.SigningPsbtCompleted)
    - [SigningPsbtExpired](#scalar.covenant.v1beta1.SigningPsbtExpired)
    - [SigningPsbtStarted](#scalar.covenant.v1beta1.SigningPsbtStarted)
    - [SigningPsbtStarted.PubKeysEntry](#scalar.covenant.v1beta1.SigningPsbtStarted.PubKeysEntry)
    - [TapScriptSigsSubmitted](#scalar.covenant.v1beta1.TapScriptSigsSubmitted)
  
- [scalar/covenant/v1beta1/params.proto](#scalar/covenant/v1beta1/params.proto)
    - [Params](#scalar.covenant.v1beta1.Params)
  
- [scalar/covenant/v1beta1/genesis.proto](#scalar/covenant/v1beta1/genesis.proto)
    - [GenesisState](#scalar.covenant.v1beta1.GenesisState)
  
- [scalar/multisig/v1beta1/params.proto](#scalar/multisig/v1beta1/params.proto)
    - [Params](#scalar.multisig.v1beta1.Params)
  
- [scalar/multisig/v1beta1/query.proto](#scalar/multisig/v1beta1/query.proto)
    - [KeyIDRequest](#scalar.multisig.v1beta1.KeyIDRequest)
    - [KeyIDResponse](#scalar.multisig.v1beta1.KeyIDResponse)
    - [KeyRequest](#scalar.multisig.v1beta1.KeyRequest)
    - [KeyResponse](#scalar.multisig.v1beta1.KeyResponse)
    - [KeygenParticipant](#scalar.multisig.v1beta1.KeygenParticipant)
    - [KeygenSessionRequest](#scalar.multisig.v1beta1.KeygenSessionRequest)
    - [KeygenSessionResponse](#scalar.multisig.v1beta1.KeygenSessionResponse)
    - [NextKeyIDRequest](#scalar.multisig.v1beta1.NextKeyIDRequest)
    - [NextKeyIDResponse](#scalar.multisig.v1beta1.NextKeyIDResponse)
    - [ParamsRequest](#scalar.multisig.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.multisig.v1beta1.ParamsResponse)
  
- [scalar/covenant/v1beta1/query.proto](#scalar/covenant/v1beta1/query.proto)
    - [CustodiansRequest](#scalar.covenant.v1beta1.CustodiansRequest)
    - [CustodiansResponse](#scalar.covenant.v1beta1.CustodiansResponse)
    - [GroupsRequest](#scalar.covenant.v1beta1.GroupsRequest)
    - [GroupsResponse](#scalar.covenant.v1beta1.GroupsResponse)
    - [KeyRequest](#scalar.covenant.v1beta1.KeyRequest)
    - [KeyResponse](#scalar.covenant.v1beta1.KeyResponse)
    - [ParamsRequest](#scalar.covenant.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.covenant.v1beta1.ParamsResponse)
  
- [scalar/covenant/v1beta1/tx.proto](#scalar/covenant/v1beta1/tx.proto)
    - [AddCustodianToGroupRequest](#scalar.covenant.v1beta1.AddCustodianToGroupRequest)
    - [CreateCustodianGroupRequest](#scalar.covenant.v1beta1.CreateCustodianGroupRequest)
    - [CreateCustodianGroupResponse](#scalar.covenant.v1beta1.CreateCustodianGroupResponse)
    - [CreateCustodianRequest](#scalar.covenant.v1beta1.CreateCustodianRequest)
    - [CreateCustodianResponse](#scalar.covenant.v1beta1.CreateCustodianResponse)
    - [CustodianToGroupResponse](#scalar.covenant.v1beta1.CustodianToGroupResponse)
    - [RemoveCustodianFromGroupRequest](#scalar.covenant.v1beta1.RemoveCustodianFromGroupRequest)
    - [RotateKeyRequest](#scalar.covenant.v1beta1.RotateKeyRequest)
    - [RotateKeyResponse](#scalar.covenant.v1beta1.RotateKeyResponse)
    - [SubmitTapScriptSigsRequest](#scalar.covenant.v1beta1.SubmitTapScriptSigsRequest)
    - [SubmitTapScriptSigsResponse](#scalar.covenant.v1beta1.SubmitTapScriptSigsResponse)
    - [UpdateCustodianGroupRequest](#scalar.covenant.v1beta1.UpdateCustodianGroupRequest)
    - [UpdateCustodianGroupResponse](#scalar.covenant.v1beta1.UpdateCustodianGroupResponse)
    - [UpdateCustodianRequest](#scalar.covenant.v1beta1.UpdateCustodianRequest)
    - [UpdateCustodianResponse](#scalar.covenant.v1beta1.UpdateCustodianResponse)
  
- [scalar/covenant/v1beta1/service.proto](#scalar/covenant/v1beta1/service.proto)
    - [MsgService](#scalar.covenant.v1beta1.MsgService)
    - [QueryService](#scalar.covenant.v1beta1.QueryService)
  
- [scalar/multisig/v1beta1/events.proto](#scalar/multisig/v1beta1/events.proto)
    - [KeyAssigned](#scalar.multisig.v1beta1.KeyAssigned)
    - [KeyRotated](#scalar.multisig.v1beta1.KeyRotated)
    - [KeygenCompleted](#scalar.multisig.v1beta1.KeygenCompleted)
    - [KeygenExpired](#scalar.multisig.v1beta1.KeygenExpired)
    - [KeygenOptIn](#scalar.multisig.v1beta1.KeygenOptIn)
    - [KeygenOptOut](#scalar.multisig.v1beta1.KeygenOptOut)
    - [KeygenStarted](#scalar.multisig.v1beta1.KeygenStarted)
    - [PubKeySubmitted](#scalar.multisig.v1beta1.PubKeySubmitted)
    - [SignatureSubmitted](#scalar.multisig.v1beta1.SignatureSubmitted)
    - [SigningCompleted](#scalar.multisig.v1beta1.SigningCompleted)
    - [SigningExpired](#scalar.multisig.v1beta1.SigningExpired)
    - [SigningStarted](#scalar.multisig.v1beta1.SigningStarted)
    - [SigningStarted.PubKeysEntry](#scalar.multisig.v1beta1.SigningStarted.PubKeysEntry)
  
- [scalar/multisig/v1beta1/genesis.proto](#scalar/multisig/v1beta1/genesis.proto)
    - [GenesisState](#scalar.multisig.v1beta1.GenesisState)
  
- [scalar/multisig/v1beta1/tx.proto](#scalar/multisig/v1beta1/tx.proto)
    - [KeygenOptInRequest](#scalar.multisig.v1beta1.KeygenOptInRequest)
    - [KeygenOptInResponse](#scalar.multisig.v1beta1.KeygenOptInResponse)
    - [KeygenOptOutRequest](#scalar.multisig.v1beta1.KeygenOptOutRequest)
    - [KeygenOptOutResponse](#scalar.multisig.v1beta1.KeygenOptOutResponse)
    - [RotateKeyRequest](#scalar.multisig.v1beta1.RotateKeyRequest)
    - [RotateKeyResponse](#scalar.multisig.v1beta1.RotateKeyResponse)
    - [StartKeygenRequest](#scalar.multisig.v1beta1.StartKeygenRequest)
    - [StartKeygenResponse](#scalar.multisig.v1beta1.StartKeygenResponse)
    - [SubmitPubKeyRequest](#scalar.multisig.v1beta1.SubmitPubKeyRequest)
    - [SubmitPubKeyResponse](#scalar.multisig.v1beta1.SubmitPubKeyResponse)
    - [SubmitSignatureRequest](#scalar.multisig.v1beta1.SubmitSignatureRequest)
    - [SubmitSignatureResponse](#scalar.multisig.v1beta1.SubmitSignatureResponse)
  
- [scalar/multisig/v1beta1/service.proto](#scalar/multisig/v1beta1/service.proto)
    - [MsgService](#scalar.multisig.v1beta1.MsgService)
    - [QueryService](#scalar.multisig.v1beta1.QueryService)
  
- [scalar/nexus/exported/v1beta1/types.proto](#scalar/nexus/exported/v1beta1/types.proto)
    - [Asset](#scalar.nexus.exported.v1beta1.Asset)
    - [Chain](#scalar.nexus.exported.v1beta1.Chain)
    - [CrossChainAddress](#scalar.nexus.exported.v1beta1.CrossChainAddress)
    - [CrossChainTransfer](#scalar.nexus.exported.v1beta1.CrossChainTransfer)
    - [FeeInfo](#scalar.nexus.exported.v1beta1.FeeInfo)
    - [GeneralMessage](#scalar.nexus.exported.v1beta1.GeneralMessage)
    - [TransferFee](#scalar.nexus.exported.v1beta1.TransferFee)
    - [WasmMessage](#scalar.nexus.exported.v1beta1.WasmMessage)
  
    - [GeneralMessage.Status](#scalar.nexus.exported.v1beta1.GeneralMessage.Status)
    - [TransferDirection](#scalar.nexus.exported.v1beta1.TransferDirection)
    - [TransferState](#scalar.nexus.exported.v1beta1.TransferState)
  
- [scalar/nexus/v1beta1/events.proto](#scalar/nexus/v1beta1/events.proto)
    - [FeeDeducted](#scalar.nexus.v1beta1.FeeDeducted)
    - [InsufficientFee](#scalar.nexus.v1beta1.InsufficientFee)
    - [MessageExecuted](#scalar.nexus.v1beta1.MessageExecuted)
    - [MessageFailed](#scalar.nexus.v1beta1.MessageFailed)
    - [MessageProcessing](#scalar.nexus.v1beta1.MessageProcessing)
    - [MessageReceived](#scalar.nexus.v1beta1.MessageReceived)
    - [RateLimitUpdated](#scalar.nexus.v1beta1.RateLimitUpdated)
    - [WasmMessageRouted](#scalar.nexus.v1beta1.WasmMessageRouted)
  
- [scalar/nexus/v1beta1/params.proto](#scalar/nexus/v1beta1/params.proto)
    - [Params](#scalar.nexus.v1beta1.Params)
  
- [scalar/utils/v1beta1/bitmap.proto](#scalar/utils/v1beta1/bitmap.proto)
    - [Bitmap](#scalar.utils.v1beta1.Bitmap)
    - [CircularBuffer](#scalar.utils.v1beta1.CircularBuffer)
  
- [scalar/nexus/v1beta1/types.proto](#scalar/nexus/v1beta1/types.proto)
    - [ChainState](#scalar.nexus.v1beta1.ChainState)
    - [LinkedAddresses](#scalar.nexus.v1beta1.LinkedAddresses)
    - [MaintainerState](#scalar.nexus.v1beta1.MaintainerState)
    - [RateLimit](#scalar.nexus.v1beta1.RateLimit)
    - [TransferEpoch](#scalar.nexus.v1beta1.TransferEpoch)
  
- [scalar/nexus/v1beta1/genesis.proto](#scalar/nexus/v1beta1/genesis.proto)
    - [GenesisState](#scalar.nexus.v1beta1.GenesisState)
  
- [scalar/nexus/v1beta1/query.proto](#scalar/nexus/v1beta1/query.proto)
    - [AssetsRequest](#scalar.nexus.v1beta1.AssetsRequest)
    - [AssetsResponse](#scalar.nexus.v1beta1.AssetsResponse)
    - [ChainMaintainersRequest](#scalar.nexus.v1beta1.ChainMaintainersRequest)
    - [ChainMaintainersResponse](#scalar.nexus.v1beta1.ChainMaintainersResponse)
    - [ChainStateRequest](#scalar.nexus.v1beta1.ChainStateRequest)
    - [ChainStateResponse](#scalar.nexus.v1beta1.ChainStateResponse)
    - [ChainsByAssetRequest](#scalar.nexus.v1beta1.ChainsByAssetRequest)
    - [ChainsByAssetResponse](#scalar.nexus.v1beta1.ChainsByAssetResponse)
    - [ChainsRequest](#scalar.nexus.v1beta1.ChainsRequest)
    - [ChainsResponse](#scalar.nexus.v1beta1.ChainsResponse)
    - [FeeInfoRequest](#scalar.nexus.v1beta1.FeeInfoRequest)
    - [FeeInfoResponse](#scalar.nexus.v1beta1.FeeInfoResponse)
    - [LatestDepositAddressRequest](#scalar.nexus.v1beta1.LatestDepositAddressRequest)
    - [LatestDepositAddressResponse](#scalar.nexus.v1beta1.LatestDepositAddressResponse)
    - [MessageRequest](#scalar.nexus.v1beta1.MessageRequest)
    - [MessageResponse](#scalar.nexus.v1beta1.MessageResponse)
    - [ParamsRequest](#scalar.nexus.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.nexus.v1beta1.ParamsResponse)
    - [RecipientAddressRequest](#scalar.nexus.v1beta1.RecipientAddressRequest)
    - [RecipientAddressResponse](#scalar.nexus.v1beta1.RecipientAddressResponse)
    - [TransferFeeRequest](#scalar.nexus.v1beta1.TransferFeeRequest)
    - [TransferFeeResponse](#scalar.nexus.v1beta1.TransferFeeResponse)
    - [TransferRateLimit](#scalar.nexus.v1beta1.TransferRateLimit)
    - [TransferRateLimitRequest](#scalar.nexus.v1beta1.TransferRateLimitRequest)
    - [TransferRateLimitResponse](#scalar.nexus.v1beta1.TransferRateLimitResponse)
    - [TransfersForChainRequest](#scalar.nexus.v1beta1.TransfersForChainRequest)
    - [TransfersForChainResponse](#scalar.nexus.v1beta1.TransfersForChainResponse)
  
    - [ChainStatus](#scalar.nexus.v1beta1.ChainStatus)
  
- [scalar/nexus/v1beta1/tx.proto](#scalar/nexus/v1beta1/tx.proto)
    - [ActivateChainRequest](#scalar.nexus.v1beta1.ActivateChainRequest)
    - [ActivateChainResponse](#scalar.nexus.v1beta1.ActivateChainResponse)
    - [DeactivateChainRequest](#scalar.nexus.v1beta1.DeactivateChainRequest)
    - [DeactivateChainResponse](#scalar.nexus.v1beta1.DeactivateChainResponse)
    - [DeregisterChainMaintainerRequest](#scalar.nexus.v1beta1.DeregisterChainMaintainerRequest)
    - [DeregisterChainMaintainerResponse](#scalar.nexus.v1beta1.DeregisterChainMaintainerResponse)
    - [RegisterAssetFeeRequest](#scalar.nexus.v1beta1.RegisterAssetFeeRequest)
    - [RegisterAssetFeeResponse](#scalar.nexus.v1beta1.RegisterAssetFeeResponse)
    - [RegisterChainMaintainerRequest](#scalar.nexus.v1beta1.RegisterChainMaintainerRequest)
    - [RegisterChainMaintainerResponse](#scalar.nexus.v1beta1.RegisterChainMaintainerResponse)
    - [SetTransferRateLimitRequest](#scalar.nexus.v1beta1.SetTransferRateLimitRequest)
    - [SetTransferRateLimitResponse](#scalar.nexus.v1beta1.SetTransferRateLimitResponse)
  
- [scalar/nexus/v1beta1/service.proto](#scalar/nexus/v1beta1/service.proto)
    - [MsgService](#scalar.nexus.v1beta1.MsgService)
    - [QueryService](#scalar.nexus.v1beta1.QueryService)
  
- [scalar/permission/v1beta1/types.proto](#scalar/permission/v1beta1/types.proto)
    - [GovAccount](#scalar.permission.v1beta1.GovAccount)
  
- [scalar/permission/v1beta1/params.proto](#scalar/permission/v1beta1/params.proto)
    - [Params](#scalar.permission.v1beta1.Params)
  
- [scalar/permission/v1beta1/genesis.proto](#scalar/permission/v1beta1/genesis.proto)
    - [GenesisState](#scalar.permission.v1beta1.GenesisState)
  
- [scalar/permission/v1beta1/query.proto](#scalar/permission/v1beta1/query.proto)
    - [ParamsRequest](#scalar.permission.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.permission.v1beta1.ParamsResponse)
    - [QueryGovernanceKeyRequest](#scalar.permission.v1beta1.QueryGovernanceKeyRequest)
    - [QueryGovernanceKeyResponse](#scalar.permission.v1beta1.QueryGovernanceKeyResponse)
  
- [scalar/permission/v1beta1/tx.proto](#scalar/permission/v1beta1/tx.proto)
    - [DeregisterControllerRequest](#scalar.permission.v1beta1.DeregisterControllerRequest)
    - [DeregisterControllerResponse](#scalar.permission.v1beta1.DeregisterControllerResponse)
    - [RegisterControllerRequest](#scalar.permission.v1beta1.RegisterControllerRequest)
    - [RegisterControllerResponse](#scalar.permission.v1beta1.RegisterControllerResponse)
    - [UpdateGovernanceKeyRequest](#scalar.permission.v1beta1.UpdateGovernanceKeyRequest)
    - [UpdateGovernanceKeyResponse](#scalar.permission.v1beta1.UpdateGovernanceKeyResponse)
  
- [scalar/permission/v1beta1/service.proto](#scalar/permission/v1beta1/service.proto)
    - [Msg](#scalar.permission.v1beta1.Msg)
    - [Query](#scalar.permission.v1beta1.Query)
  
- [scalar/protocol/v1beta1/types.proto](#scalar/protocol/v1beta1/types.proto)
    - [Protocol](#scalar.protocol.v1beta1.Protocol)
    - [ProtocolDetails](#scalar.protocol.v1beta1.ProtocolDetails)
  
- [scalar/protocol/v1beta1/genesis.proto](#scalar/protocol/v1beta1/genesis.proto)
    - [GenesisState](#scalar.protocol.v1beta1.GenesisState)
  
- [scalar/protocol/v1beta1/params.proto](#scalar/protocol/v1beta1/params.proto)
    - [Params](#scalar.protocol.v1beta1.Params)
  
- [scalar/protocol/v1beta1/query.proto](#scalar/protocol/v1beta1/query.proto)
    - [ProtocolRequest](#scalar.protocol.v1beta1.ProtocolRequest)
    - [ProtocolResponse](#scalar.protocol.v1beta1.ProtocolResponse)
    - [ProtocolsRequest](#scalar.protocol.v1beta1.ProtocolsRequest)
    - [ProtocolsResponse](#scalar.protocol.v1beta1.ProtocolsResponse)
  
- [scalar/protocol/v1beta1/tx.proto](#scalar/protocol/v1beta1/tx.proto)
    - [AddSupportedChainRequest](#scalar.protocol.v1beta1.AddSupportedChainRequest)
    - [AddSupportedChainResponse](#scalar.protocol.v1beta1.AddSupportedChainResponse)
    - [CreateProtocolRequest](#scalar.protocol.v1beta1.CreateProtocolRequest)
    - [CreateProtocolResponse](#scalar.protocol.v1beta1.CreateProtocolResponse)
    - [UpdateProtocolRequest](#scalar.protocol.v1beta1.UpdateProtocolRequest)
    - [UpdateProtocolResponse](#scalar.protocol.v1beta1.UpdateProtocolResponse)
    - [UpdateSupportedChainRequest](#scalar.protocol.v1beta1.UpdateSupportedChainRequest)
    - [UpdateSupportedChainResponse](#scalar.protocol.v1beta1.UpdateSupportedChainResponse)
  
- [scalar/protocol/v1beta1/service.proto](#scalar/protocol/v1beta1/service.proto)
    - [Msg](#scalar.protocol.v1beta1.Msg)
    - [Query](#scalar.protocol.v1beta1.Query)
  
- [scalar/reward/v1beta1/params.proto](#scalar/reward/v1beta1/params.proto)
    - [Params](#scalar.reward.v1beta1.Params)
  
- [scalar/reward/v1beta1/types.proto](#scalar/reward/v1beta1/types.proto)
    - [Pool](#scalar.reward.v1beta1.Pool)
    - [Pool.Reward](#scalar.reward.v1beta1.Pool.Reward)
    - [Refund](#scalar.reward.v1beta1.Refund)
  
- [scalar/reward/v1beta1/genesis.proto](#scalar/reward/v1beta1/genesis.proto)
    - [GenesisState](#scalar.reward.v1beta1.GenesisState)
  
- [scalar/reward/v1beta1/query.proto](#scalar/reward/v1beta1/query.proto)
    - [InflationRateRequest](#scalar.reward.v1beta1.InflationRateRequest)
    - [InflationRateResponse](#scalar.reward.v1beta1.InflationRateResponse)
    - [ParamsRequest](#scalar.reward.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.reward.v1beta1.ParamsResponse)
  
- [scalar/reward/v1beta1/tx.proto](#scalar/reward/v1beta1/tx.proto)
    - [RefundMsgRequest](#scalar.reward.v1beta1.RefundMsgRequest)
    - [RefundMsgResponse](#scalar.reward.v1beta1.RefundMsgResponse)
  
- [scalar/reward/v1beta1/service.proto](#scalar/reward/v1beta1/service.proto)
    - [MsgService](#scalar.reward.v1beta1.MsgService)
    - [QueryService](#scalar.reward.v1beta1.QueryService)
  
- [scalar/scalarnet/v1beta1/events.proto](#scalar/scalarnet/v1beta1/events.proto)
    - [ContractCallSubmitted](#scalar.scalarnet.v1beta1.ContractCallSubmitted)
    - [ContractCallWithTokenSubmitted](#scalar.scalarnet.v1beta1.ContractCallWithTokenSubmitted)
    - [FeeCollected](#scalar.scalarnet.v1beta1.FeeCollected)
    - [FeePaid](#scalar.scalarnet.v1beta1.FeePaid)
    - [IBCTransferCompleted](#scalar.scalarnet.v1beta1.IBCTransferCompleted)
    - [IBCTransferFailed](#scalar.scalarnet.v1beta1.IBCTransferFailed)
    - [IBCTransferRetried](#scalar.scalarnet.v1beta1.IBCTransferRetried)
    - [IBCTransferSent](#scalar.scalarnet.v1beta1.IBCTransferSent)
    - [ScalarTransferCompleted](#scalar.scalarnet.v1beta1.ScalarTransferCompleted)
    - [TokenSent](#scalar.scalarnet.v1beta1.TokenSent)
  
- [scalar/scalarnet/v1beta1/params.proto](#scalar/scalarnet/v1beta1/params.proto)
    - [CallContractProposalMinDeposit](#scalar.scalarnet.v1beta1.CallContractProposalMinDeposit)
    - [Params](#scalar.scalarnet.v1beta1.Params)
  
- [scalar/scalarnet/v1beta1/types.proto](#scalar/scalarnet/v1beta1/types.proto)
    - [Asset](#scalar.scalarnet.v1beta1.Asset)
    - [CosmosChain](#scalar.scalarnet.v1beta1.CosmosChain)
    - [Fee](#scalar.scalarnet.v1beta1.Fee)
    - [IBCTransfer](#scalar.scalarnet.v1beta1.IBCTransfer)
  
    - [IBCTransfer.Status](#scalar.scalarnet.v1beta1.IBCTransfer.Status)
  
- [scalar/scalarnet/v1beta1/genesis.proto](#scalar/scalarnet/v1beta1/genesis.proto)
    - [GenesisState](#scalar.scalarnet.v1beta1.GenesisState)
    - [GenesisState.SeqIdMappingEntry](#scalar.scalarnet.v1beta1.GenesisState.SeqIdMappingEntry)
  
- [scalar/scalarnet/v1beta1/proposal.proto](#scalar/scalarnet/v1beta1/proposal.proto)
    - [CallContractsProposal](#scalar.scalarnet.v1beta1.CallContractsProposal)
    - [ContractCall](#scalar.scalarnet.v1beta1.ContractCall)
  
- [scalar/scalarnet/v1beta1/query.proto](#scalar/scalarnet/v1beta1/query.proto)
    - [ChainByIBCPathRequest](#scalar.scalarnet.v1beta1.ChainByIBCPathRequest)
    - [ChainByIBCPathResponse](#scalar.scalarnet.v1beta1.ChainByIBCPathResponse)
    - [IBCPathRequest](#scalar.scalarnet.v1beta1.IBCPathRequest)
    - [IBCPathResponse](#scalar.scalarnet.v1beta1.IBCPathResponse)
    - [ParamsRequest](#scalar.scalarnet.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.scalarnet.v1beta1.ParamsResponse)
    - [PendingIBCTransferCountRequest](#scalar.scalarnet.v1beta1.PendingIBCTransferCountRequest)
    - [PendingIBCTransferCountResponse](#scalar.scalarnet.v1beta1.PendingIBCTransferCountResponse)
    - [PendingIBCTransferCountResponse.TransfersByChainEntry](#scalar.scalarnet.v1beta1.PendingIBCTransferCountResponse.TransfersByChainEntry)
  
- [scalar/scalarnet/v1beta1/tx.proto](#scalar/scalarnet/v1beta1/tx.proto)
    - [AddCosmosBasedChainRequest](#scalar.scalarnet.v1beta1.AddCosmosBasedChainRequest)
    - [AddCosmosBasedChainResponse](#scalar.scalarnet.v1beta1.AddCosmosBasedChainResponse)
    - [CallContractRequest](#scalar.scalarnet.v1beta1.CallContractRequest)
    - [CallContractResponse](#scalar.scalarnet.v1beta1.CallContractResponse)
    - [ConfirmDepositRequest](#scalar.scalarnet.v1beta1.ConfirmDepositRequest)
    - [ConfirmDepositResponse](#scalar.scalarnet.v1beta1.ConfirmDepositResponse)
    - [ExecutePendingTransfersRequest](#scalar.scalarnet.v1beta1.ExecutePendingTransfersRequest)
    - [ExecutePendingTransfersResponse](#scalar.scalarnet.v1beta1.ExecutePendingTransfersResponse)
    - [LinkRequest](#scalar.scalarnet.v1beta1.LinkRequest)
    - [LinkResponse](#scalar.scalarnet.v1beta1.LinkResponse)
    - [RegisterAssetRequest](#scalar.scalarnet.v1beta1.RegisterAssetRequest)
    - [RegisterAssetResponse](#scalar.scalarnet.v1beta1.RegisterAssetResponse)
    - [RegisterFeeCollectorRequest](#scalar.scalarnet.v1beta1.RegisterFeeCollectorRequest)
    - [RegisterFeeCollectorResponse](#scalar.scalarnet.v1beta1.RegisterFeeCollectorResponse)
    - [RegisterIBCPathRequest](#scalar.scalarnet.v1beta1.RegisterIBCPathRequest)
    - [RegisterIBCPathResponse](#scalar.scalarnet.v1beta1.RegisterIBCPathResponse)
    - [RetryIBCTransferRequest](#scalar.scalarnet.v1beta1.RetryIBCTransferRequest)
    - [RetryIBCTransferResponse](#scalar.scalarnet.v1beta1.RetryIBCTransferResponse)
    - [RouteIBCTransfersRequest](#scalar.scalarnet.v1beta1.RouteIBCTransfersRequest)
    - [RouteIBCTransfersResponse](#scalar.scalarnet.v1beta1.RouteIBCTransfersResponse)
    - [RouteMessageRequest](#scalar.scalarnet.v1beta1.RouteMessageRequest)
    - [RouteMessageResponse](#scalar.scalarnet.v1beta1.RouteMessageResponse)
  
- [scalar/scalarnet/v1beta1/service.proto](#scalar/scalarnet/v1beta1/service.proto)
    - [MsgService](#scalar.scalarnet.v1beta1.MsgService)
    - [QueryService](#scalar.scalarnet.v1beta1.QueryService)
  
- [scalar/snapshot/v1beta1/params.proto](#scalar/snapshot/v1beta1/params.proto)
    - [Params](#scalar.snapshot.v1beta1.Params)
  
- [scalar/snapshot/v1beta1/types.proto](#scalar/snapshot/v1beta1/types.proto)
    - [ProxiedValidator](#scalar.snapshot.v1beta1.ProxiedValidator)
  
- [scalar/snapshot/v1beta1/genesis.proto](#scalar/snapshot/v1beta1/genesis.proto)
    - [GenesisState](#scalar.snapshot.v1beta1.GenesisState)
  
- [scalar/snapshot/v1beta1/query.proto](#scalar/snapshot/v1beta1/query.proto)
    - [ParamsRequest](#scalar.snapshot.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.snapshot.v1beta1.ParamsResponse)
    - [QueryValidatorsResponse](#scalar.snapshot.v1beta1.QueryValidatorsResponse)
    - [QueryValidatorsResponse.TssIllegibilityInfo](#scalar.snapshot.v1beta1.QueryValidatorsResponse.TssIllegibilityInfo)
    - [QueryValidatorsResponse.Validator](#scalar.snapshot.v1beta1.QueryValidatorsResponse.Validator)
  
- [scalar/snapshot/v1beta1/tx.proto](#scalar/snapshot/v1beta1/tx.proto)
    - [DeactivateProxyRequest](#scalar.snapshot.v1beta1.DeactivateProxyRequest)
    - [DeactivateProxyResponse](#scalar.snapshot.v1beta1.DeactivateProxyResponse)
    - [RegisterProxyRequest](#scalar.snapshot.v1beta1.RegisterProxyRequest)
    - [RegisterProxyResponse](#scalar.snapshot.v1beta1.RegisterProxyResponse)
  
- [scalar/snapshot/v1beta1/service.proto](#scalar/snapshot/v1beta1/service.proto)
    - [MsgService](#scalar.snapshot.v1beta1.MsgService)
    - [QueryService](#scalar.snapshot.v1beta1.QueryService)
  
- [scalar/tss/tofnd/v1beta1/common.proto](#scalar/tss/tofnd/v1beta1/common.proto)
    - [KeyPresenceRequest](#tofnd.KeyPresenceRequest)
    - [KeyPresenceResponse](#tofnd.KeyPresenceResponse)
  
    - [KeyPresenceResponse.Response](#tofnd.KeyPresenceResponse.Response)
  
- [scalar/tss/tofnd/v1beta1/multisig.proto](#scalar/tss/tofnd/v1beta1/multisig.proto)
    - [KeygenRequest](#tofnd.KeygenRequest)
    - [KeygenResponse](#tofnd.KeygenResponse)
    - [SignRequest](#tofnd.SignRequest)
    - [SignResponse](#tofnd.SignResponse)
  
    - [Multisig](#tofnd.Multisig)
  
- [scalar/tss/tofnd/v1beta1/tofnd.proto](#scalar/tss/tofnd/v1beta1/tofnd.proto)
    - [KeygenInit](#tofnd.KeygenInit)
    - [KeygenOutput](#tofnd.KeygenOutput)
    - [MessageIn](#tofnd.MessageIn)
    - [MessageOut](#tofnd.MessageOut)
    - [MessageOut.CriminalList](#tofnd.MessageOut.CriminalList)
    - [MessageOut.CriminalList.Criminal](#tofnd.MessageOut.CriminalList.Criminal)
    - [MessageOut.KeygenResult](#tofnd.MessageOut.KeygenResult)
    - [MessageOut.SignResult](#tofnd.MessageOut.SignResult)
    - [RecoverRequest](#tofnd.RecoverRequest)
    - [RecoverResponse](#tofnd.RecoverResponse)
    - [SignInit](#tofnd.SignInit)
    - [TrafficIn](#tofnd.TrafficIn)
    - [TrafficOut](#tofnd.TrafficOut)
  
    - [MessageOut.CriminalList.Criminal.CrimeType](#tofnd.MessageOut.CriminalList.Criminal.CrimeType)
    - [RecoverResponse.Response](#tofnd.RecoverResponse.Response)
  
- [scalar/tss/v1beta1/params.proto](#scalar/tss/v1beta1/params.proto)
    - [Params](#scalar.tss.v1beta1.Params)
  
- [scalar/tss/v1beta1/types.proto](#scalar/tss/v1beta1/types.proto)
    - [ExternalKeys](#scalar.tss.v1beta1.ExternalKeys)
    - [KeyInfo](#scalar.tss.v1beta1.KeyInfo)
    - [KeyRecoveryInfo](#scalar.tss.v1beta1.KeyRecoveryInfo)
    - [KeyRecoveryInfo.PrivateEntry](#scalar.tss.v1beta1.KeyRecoveryInfo.PrivateEntry)
    - [KeygenVoteData](#scalar.tss.v1beta1.KeygenVoteData)
    - [MultisigInfo](#scalar.tss.v1beta1.MultisigInfo)
    - [MultisigInfo.Info](#scalar.tss.v1beta1.MultisigInfo.Info)
    - [ValidatorStatus](#scalar.tss.v1beta1.ValidatorStatus)
  
- [scalar/tss/v1beta1/genesis.proto](#scalar/tss/v1beta1/genesis.proto)
    - [GenesisState](#scalar.tss.v1beta1.GenesisState)
  
- [scalar/tss/v1beta1/query.proto](#scalar/tss/v1beta1/query.proto)
    - [ParamsRequest](#scalar.tss.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.tss.v1beta1.ParamsResponse)
  
- [scalar/tss/v1beta1/tx.proto](#scalar/tss/v1beta1/tx.proto)
    - [HeartBeatRequest](#scalar.tss.v1beta1.HeartBeatRequest)
    - [HeartBeatResponse](#scalar.tss.v1beta1.HeartBeatResponse)
    - [ProcessKeygenTrafficRequest](#scalar.tss.v1beta1.ProcessKeygenTrafficRequest)
    - [ProcessKeygenTrafficResponse](#scalar.tss.v1beta1.ProcessKeygenTrafficResponse)
    - [ProcessSignTrafficRequest](#scalar.tss.v1beta1.ProcessSignTrafficRequest)
    - [ProcessSignTrafficResponse](#scalar.tss.v1beta1.ProcessSignTrafficResponse)
    - [RegisterExternalKeysRequest](#scalar.tss.v1beta1.RegisterExternalKeysRequest)
    - [RegisterExternalKeysRequest.ExternalKey](#scalar.tss.v1beta1.RegisterExternalKeysRequest.ExternalKey)
    - [RegisterExternalKeysResponse](#scalar.tss.v1beta1.RegisterExternalKeysResponse)
    - [RotateKeyRequest](#scalar.tss.v1beta1.RotateKeyRequest)
    - [RotateKeyResponse](#scalar.tss.v1beta1.RotateKeyResponse)
    - [StartKeygenRequest](#scalar.tss.v1beta1.StartKeygenRequest)
    - [StartKeygenResponse](#scalar.tss.v1beta1.StartKeygenResponse)
    - [SubmitMultisigPubKeysRequest](#scalar.tss.v1beta1.SubmitMultisigPubKeysRequest)
    - [SubmitMultisigPubKeysResponse](#scalar.tss.v1beta1.SubmitMultisigPubKeysResponse)
    - [SubmitMultisigSignaturesRequest](#scalar.tss.v1beta1.SubmitMultisigSignaturesRequest)
    - [SubmitMultisigSignaturesResponse](#scalar.tss.v1beta1.SubmitMultisigSignaturesResponse)
    - [VotePubKeyRequest](#scalar.tss.v1beta1.VotePubKeyRequest)
    - [VotePubKeyResponse](#scalar.tss.v1beta1.VotePubKeyResponse)
    - [VoteSigRequest](#scalar.tss.v1beta1.VoteSigRequest)
    - [VoteSigResponse](#scalar.tss.v1beta1.VoteSigResponse)
  
- [scalar/tss/v1beta1/service.proto](#scalar/tss/v1beta1/service.proto)
    - [MsgService](#scalar.tss.v1beta1.MsgService)
    - [QueryService](#scalar.tss.v1beta1.QueryService)
  
- [scalar/vote/v1beta1/events.proto](#scalar/vote/v1beta1/events.proto)
    - [Voted](#scalar.vote.v1beta1.Voted)
  
- [scalar/vote/v1beta1/params.proto](#scalar/vote/v1beta1/params.proto)
    - [Params](#scalar.vote.v1beta1.Params)
  
- [scalar/vote/v1beta1/genesis.proto](#scalar/vote/v1beta1/genesis.proto)
    - [GenesisState](#scalar.vote.v1beta1.GenesisState)
  
- [scalar/vote/v1beta1/query.proto](#scalar/vote/v1beta1/query.proto)
    - [ParamsRequest](#scalar.vote.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.vote.v1beta1.ParamsResponse)
  
- [scalar/vote/v1beta1/types.proto](#scalar/vote/v1beta1/types.proto)
    - [TalliedVote](#scalar.vote.v1beta1.TalliedVote)
    - [TalliedVote.IsVoterLateEntry](#scalar.vote.v1beta1.TalliedVote.IsVoterLateEntry)
  
- [scalar/vote/v1beta1/tx.proto](#scalar/vote/v1beta1/tx.proto)
    - [VoteRequest](#scalar.vote.v1beta1.VoteRequest)
    - [VoteResponse](#scalar.vote.v1beta1.VoteResponse)
  
- [scalar/vote/v1beta1/service.proto](#scalar/vote/v1beta1/service.proto)
    - [MsgService](#scalar.vote.v1beta1.MsgService)
    - [QueryService](#scalar.vote.v1beta1.QueryService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="scalar/auxiliary/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/auxiliary/v1beta1/events.proto



<a name="scalar.auxiliary.v1beta1.BatchedMessageFailed"></a>

### BatchedMessageFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [int32](#int32) |  |  |
| `error` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/auxiliary/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/auxiliary/v1beta1/genesis.proto



<a name="scalar.auxiliary.v1beta1.GenesisState"></a>

### GenesisState
GenesisState represents the genesis state





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/permission/exported/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/permission/exported/v1beta1/types.proto


 <!-- end messages -->


<a name="scalar.permission.exported.v1beta1.Role"></a>

### Role


| Name | Number | Description |
| ---- | ------ | ----------- |
| ROLE_UNSPECIFIED | 0 |  |
| ROLE_UNRESTRICTED | 1 |  |
| ROLE_CHAIN_MANAGEMENT | 2 |  |
| ROLE_ACCESS_CONTROL | 3 |  |


 <!-- end enums -->


<a name="scalar/permission/exported/v1beta1/types.proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| `permission_role` | Role | .google.protobuf.MessageOptions | 50000 | 50000-99999 reserved for use withing individual organizations |

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/auxiliary/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/auxiliary/v1beta1/tx.proto



<a name="scalar.auxiliary.v1beta1.BatchRequest"></a>

### BatchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="scalar.auxiliary.v1beta1.BatchResponse"></a>

### BatchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `responses` | [BatchResponse.Response](#scalar.auxiliary.v1beta1.BatchResponse.Response) | repeated |  |






<a name="scalar.auxiliary.v1beta1.BatchResponse.Response"></a>

### BatchResponse.Response



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `result` | [cosmos.base.abci.v1beta1.Result](#cosmos.base.abci.v1beta1.Result) |  |  |
| `err` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/auxiliary/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/auxiliary/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.auxiliary.v1beta1.MsgService"></a>

### MsgService
Msg defines the nexus Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Batch` | [BatchRequest](#scalar.auxiliary.v1beta1.BatchRequest) | [BatchResponse](#scalar.auxiliary.v1beta1.BatchResponse) |  | POST|/scalar/auxiliary/batch|

 <!-- end services -->



<a name="scalar/chains/btc/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/btc/v1beta1/types.proto



<a name="scalar.chains.btc.v1beta1.BtcToken"></a>

### BtcToken






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/chains/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/v1beta1/types.proto



<a name="scalar.chains.v1beta1.Asset"></a>

### Asset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.BurnerInfo"></a>

### BurnerInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `burner_address` | [bytes](#bytes) |  |  |
| `token_address` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |
| `salt` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.Command"></a>

### Command



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [bytes](#bytes) |  |  |
| `params` | [bytes](#bytes) |  |  |
| `key_id` | [string](#string) |  |  |
| `max_gas_cost` | [uint32](#uint32) |  |  |
| `type` | [CommandType](#scalar.chains.v1beta1.CommandType) |  |  |
| `payload` | [bytes](#bytes) |  | This field is used as extra data for the command, metadata is encoded in the payload, it can be fee information, etc. |






<a name="scalar.chains.v1beta1.CommandBatchMetadata"></a>

### CommandBatchMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [bytes](#bytes) |  |  |
| `command_ids` | [bytes](#bytes) | repeated |  |
| `data` | [bytes](#bytes) |  |  |
| `sig_hash` | [bytes](#bytes) |  |  |
| `status` | [BatchedCommandsStatus](#scalar.chains.v1beta1.BatchedCommandsStatus) |  |  |
| `key_id` | [string](#string) |  |  |
| `prev_batched_commands_id` | [bytes](#bytes) |  |  |
| `signature` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `extra_data` | [bytes](#bytes) | repeated | Store payload of each command to create psbt |






<a name="scalar.chains.v1beta1.Gateway"></a>

### Gateway



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.PollCompleted"></a>

### PollCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.chains.v1beta1.PollExpired"></a>

### PollExpired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.chains.v1beta1.PollFailed"></a>

### PollFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.chains.v1beta1.PollMapping"></a>

### PollMapping



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.chains.v1beta1.PollMetadata"></a>

### PollMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.Proof"></a>

### Proof



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [string](#string) | repeated |  |
| `weights` | [string](#string) | repeated |  |
| `threshold` | [string](#string) |  |  |
| `signatures` | [string](#string) | repeated |  |






<a name="scalar.chains.v1beta1.SigMetadata"></a>

### SigMetadata
SigMetadata stores necessary information for external apps to map signature
results to chains relay transaction types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [SigType](#scalar.chains.v1beta1.SigType) |  |  |
| `chain` | [string](#string) |  |  |
| `command_batch_id` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.SourceTx"></a>

### SourceTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `amount` | [bytes](#bytes) |  |  |
| `asset` | [string](#string) |  | TODO: change to asset type: sats, runes, btc, etc |
| `destination_chain` | [string](#string) |  |  |
| `destination_recipient_address` | [bytes](#bytes) |  |  |
| `log_index` | [uint64](#uint64) |  |  |






<a name="scalar.chains.v1beta1.TokenDetails"></a>

### TokenDetails



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_name` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `decimals` | [uint32](#uint32) |  |  |
| `capacity` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.TransferKey"></a>

### TransferKey
TransferKey contains information for a transfer operatorship


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `next_key_id` | [string](#string) |  |  |





 <!-- end messages -->


<a name="scalar.chains.v1beta1.BatchedCommandsStatus"></a>

### BatchedCommandsStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| BATCHED_COMMANDS_STATUS_UNSPECIFIED | 0 |  |
| BATCHED_COMMANDS_STATUS_SIGNING | 1 |  |
| BATCHED_COMMANDS_STATUS_ABORTED | 2 |  |
| BATCHED_COMMANDS_STATUS_SIGNED | 3 |  |



<a name="scalar.chains.v1beta1.CommandType"></a>

### CommandType


| Name | Number | Description |
| ---- | ------ | ----------- |
| COMMAND_TYPE_UNSPECIFIED | 0 |  |
| COMMAND_TYPE_MINT_TOKEN | 1 |  |
| COMMAND_TYPE_DEPLOY_TOKEN | 2 |  |
| COMMAND_TYPE_BURN_TOKEN | 3 |  |
| COMMAND_TYPE_TRANSFER_OPERATORSHIP | 4 |  |
| COMMAND_TYPE_APPROVE_CONTRACT_CALL_WITH_MINT | 5 |  |
| COMMAND_TYPE_APPROVE_CONTRACT_CALL | 6 |  |



<a name="scalar.chains.v1beta1.DepositStatus"></a>

### DepositStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| DEPOSIT_STATUS_UNSPECIFIED | 0 |  |
| DEPOSIT_STATUS_PENDING | 1 |  |
| DEPOSIT_STATUS_CONFIRMED | 2 |  |
| DEPOSIT_STATUS_BURNED | 3 |  |



<a name="scalar.chains.v1beta1.SigType"></a>

### SigType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SIG_TYPE_UNSPECIFIED | 0 |  |
| SIG_TYPE_TX | 1 |  |
| SIG_TYPE_COMMAND | 2 |  |



<a name="scalar.chains.v1beta1.SourceTxStatus"></a>

### SourceTxStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STAKING_TX_STATUS_UNSPECIFIED | 0 |  |
| STAKING_TX_STATUS_PENDING | 1 |  |
| STAKING_TX_STATUS_CONFIRMED | 2 |  |
| STAKING_TX_STATUS_COMPLETED | 3 |  |



<a name="scalar.chains.v1beta1.Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | these enum values are used for bitwise operations, therefore they need to be powers of 2 |
| STATUS_INITIALIZED | 1 |  |
| STATUS_PENDING | 2 |  |
| STATUS_CONFIRMED | 4 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/utils/v1beta1/threshold.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/utils/v1beta1/threshold.proto



<a name="scalar.utils.v1beta1.Threshold"></a>

### Threshold



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `numerator` | [int64](#int64) |  | split threshold into Numerator and denominator to avoid floating point errors down the line |
| `denominator` | [int64](#int64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/tss/exported/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/exported/v1beta1/types.proto



<a name="scalar.tss.exported.v1beta1.KeyRequirement"></a>

### KeyRequirement
KeyRequirement defines requirements for keys


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_role` | [KeyRole](#scalar.tss.exported.v1beta1.KeyRole) |  |  |
| `key_type` | [KeyType](#scalar.tss.exported.v1beta1.KeyType) |  |  |
| `min_keygen_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `safety_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `key_share_distribution_policy` | [KeyShareDistributionPolicy](#scalar.tss.exported.v1beta1.KeyShareDistributionPolicy) |  |  |
| `max_total_share_count` | [int64](#int64) |  |  |
| `min_total_share_count` | [int64](#int64) |  |  |
| `keygen_voting_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `sign_voting_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `keygen_timeout` | [int64](#int64) |  |  |
| `sign_timeout` | [int64](#int64) |  |  |






<a name="scalar.tss.exported.v1beta1.SigKeyPair"></a>

### SigKeyPair
PubKeyInfo holds a pubkey and a signature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_key` | [bytes](#bytes) |  |  |
| `signature` | [bytes](#bytes) |  |  |





 <!-- end messages -->


<a name="scalar.tss.exported.v1beta1.KeyRole"></a>

### KeyRole


| Name | Number | Description |
| ---- | ------ | ----------- |
| KEY_ROLE_UNSPECIFIED | 0 |  |
| KEY_ROLE_MASTER_KEY | 1 |  |
| KEY_ROLE_SECONDARY_KEY | 2 |  |
| KEY_ROLE_EXTERNAL_KEY | 3 |  |



<a name="scalar.tss.exported.v1beta1.KeyShareDistributionPolicy"></a>

### KeyShareDistributionPolicy


| Name | Number | Description |
| ---- | ------ | ----------- |
| KEY_SHARE_DISTRIBUTION_POLICY_UNSPECIFIED | 0 |  |
| KEY_SHARE_DISTRIBUTION_POLICY_WEIGHTED_BY_STAKE | 1 |  |
| KEY_SHARE_DISTRIBUTION_POLICY_ONE_PER_VALIDATOR | 2 |  |



<a name="scalar.tss.exported.v1beta1.KeyType"></a>

### KeyType


| Name | Number | Description |
| ---- | ------ | ----------- |
| KEY_TYPE_UNSPECIFIED | 0 |  |
| KEY_TYPE_NONE | 1 |  |
| KEY_TYPE_THRESHOLD | 2 |  |
| KEY_TYPE_MULTISIG | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/snapshot/exported/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/snapshot/exported/v1beta1/types.proto



<a name="scalar.snapshot.exported.v1beta1.Participant"></a>

### Participant



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `weight` | [bytes](#bytes) |  |  |






<a name="scalar.snapshot.exported.v1beta1.Snapshot"></a>

### Snapshot



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `height` | [int64](#int64) |  |  |
| `participants` | [Snapshot.ParticipantsEntry](#scalar.snapshot.exported.v1beta1.Snapshot.ParticipantsEntry) | repeated |  |
| `bonded_weight` | [bytes](#bytes) |  |  |






<a name="scalar.snapshot.exported.v1beta1.Snapshot.ParticipantsEntry"></a>

### Snapshot.ParticipantsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [Participant](#scalar.snapshot.exported.v1beta1.Participant) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/vote/exported/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/vote/exported/v1beta1/types.proto



<a name="scalar.vote.exported.v1beta1.PollKey"></a>

### PollKey
PollKey represents the key data for a poll


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `id` | [string](#string) |  |  |






<a name="scalar.vote.exported.v1beta1.PollMetadata"></a>

### PollMetadata
PollMetadata represents a poll with write-in voting, i.e. the result of the
vote can have any data type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `expires_at` | [int64](#int64) |  |  |
| `result` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `voting_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `state` | [PollState](#scalar.vote.exported.v1beta1.PollState) |  |  |
| `min_voter_count` | [int64](#int64) |  |  |
| `reward_pool_name` | [string](#string) |  |  |
| `grace_period` | [int64](#int64) |  |  |
| `completed_at` | [int64](#int64) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `snapshot` | [scalar.snapshot.exported.v1beta1.Snapshot](#scalar.snapshot.exported.v1beta1.Snapshot) |  |  |
| `module` | [string](#string) |  |  |
| `module_metadata` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="scalar.vote.exported.v1beta1.PollParticipants"></a>

### PollParticipants
PollParticipants should be embedded in poll events in other modules


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poll_id` | [uint64](#uint64) |  |  |
| `participants` | [bytes](#bytes) | repeated |  |





 <!-- end messages -->


<a name="scalar.vote.exported.v1beta1.PollState"></a>

### PollState


| Name | Number | Description |
| ---- | ------ | ----------- |
| POLL_STATE_UNSPECIFIED | 0 |  |
| POLL_STATE_PENDING | 1 |  |
| POLL_STATE_COMPLETED | 2 |  |
| POLL_STATE_FAILED | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/chains/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/v1beta1/events.proto



<a name="scalar.chains.v1beta1.BurnCommand"></a>

### BurnCommand



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `command_id` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `deposit_address` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ChainAdded"></a>

### ChainAdded



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ChainEventCompleted"></a>

### ChainEventCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ChainEventConfirmed"></a>

### ChainEventConfirmed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ChainEventFailed"></a>

### ChainEventFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ChainEventRetryFailed"></a>

### ChainEventRetryFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.CommandBatchAborted"></a>

### CommandBatchAborted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `command_batch_id` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.CommandBatchSigned"></a>

### CommandBatchSigned



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `command_batch_id` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.ConfirmDepositStarted"></a>

### ConfirmDepositStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `deposit_address` | [bytes](#bytes) |  |  |
| `token_address` | [bytes](#bytes) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [scalar.vote.exported.v1beta1.PollParticipants](#scalar.vote.exported.v1beta1.PollParticipants) |  |  |
| `asset` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ConfirmKeyTransferStarted"></a>

### ConfirmKeyTransferStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `gateway_address` | [bytes](#bytes) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [scalar.vote.exported.v1beta1.PollParticipants](#scalar.vote.exported.v1beta1.PollParticipants) |  |  |






<a name="scalar.chains.v1beta1.ConfirmTokenStarted"></a>

### ConfirmTokenStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `gateway_address` | [bytes](#bytes) |  |  |
| `token_address` | [bytes](#bytes) |  |  |
| `token_details` | [TokenDetails](#scalar.chains.v1beta1.TokenDetails) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [scalar.vote.exported.v1beta1.PollParticipants](#scalar.vote.exported.v1beta1.PollParticipants) |  |  |






<a name="scalar.chains.v1beta1.ContractCallApproved"></a>

### ContractCallApproved



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `command_id` | [bytes](#bytes) |  |  |
| `sender` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.ContractCallFailed"></a>

### ContractCallFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `message_id` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `status` | [Event.Status](#scalar.chains.v1beta1.Event.Status) |  |  |
| `index` | [uint64](#uint64) |  |  |
| `token_sent` | [EventTokenSent](#scalar.chains.v1beta1.EventTokenSent) |  |  |
| `contract_call` | [EventContractCall](#scalar.chains.v1beta1.EventContractCall) |  |  |
| `contract_call_with_token` | [EventContractCallWithToken](#scalar.chains.v1beta1.EventContractCallWithToken) |  |  |
| `contract_call_with_mint_approved` | [EventContractCallWithMintApproved](#scalar.chains.v1beta1.EventContractCallWithMintApproved) |  |  |
| `transfer` | [EventTransfer](#scalar.chains.v1beta1.EventTransfer) |  |  |
| `token_deployed` | [EventTokenDeployed](#scalar.chains.v1beta1.EventTokenDeployed) |  |  |
| `multisig_operatorship_transferred` | [EventMultisigOperatorshipTransferred](#scalar.chains.v1beta1.EventMultisigOperatorshipTransferred) |  |  |
| `source_tx_confirmation_event` | [SourceTxConfirmationEvent](#scalar.chains.v1beta1.SourceTxConfirmationEvent) |  | for general chains |






<a name="scalar.chains.v1beta1.EventConfirmSourceTxsStarted"></a>

### EventConfirmSourceTxsStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poll_mappings` | [PollMapping](#scalar.chains.v1beta1.PollMapping) | repeated |  |
| `chain` | [string](#string) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [bytes](#bytes) | repeated |  |






<a name="scalar.chains.v1beta1.EventContractCall"></a>

### EventContractCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.EventContractCallWithMintApproved"></a>

### EventContractCallWithMintApproved



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `command_id` | [bytes](#bytes) |  |  |
| `sender` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.chains.v1beta1.EventContractCallWithToken"></a>

### EventContractCallWithToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `symbol` | [string](#string) |  |  |
| `amount` | [bytes](#bytes) |  |  |
| `payload` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.EventMultisigOperatorshipTransferred"></a>

### EventMultisigOperatorshipTransferred



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `new_operators` | [bytes](#bytes) | repeated |  |
| `new_threshold` | [bytes](#bytes) |  |  |
| `new_weights` | [bytes](#bytes) | repeated |  |






<a name="scalar.chains.v1beta1.EventMultisigOwnershipTransferred"></a>

### EventMultisigOwnershipTransferred



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pre_owners` | [bytes](#bytes) | repeated |  |
| `prev_threshold` | [bytes](#bytes) |  |  |
| `new_owners` | [bytes](#bytes) | repeated |  |
| `new_threshold` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.EventTokenDeployed"></a>

### EventTokenDeployed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `token_address` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.EventTokenSent"></a>

### EventTokenSent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `transfer_id` | [uint64](#uint64) |  |  |
| `command_id` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `destination_address` | [string](#string) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.chains.v1beta1.EventTransfer"></a>

### EventTransfer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `to` | [bytes](#bytes) |  |  |
| `amount` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.MintCommand"></a>

### MintCommand



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `transfer_id` | [uint64](#uint64) |  |  |
| `command_id` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `destination_address` | [string](#string) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.chains.v1beta1.NoEventsConfirmed"></a>

### NoEventsConfirmed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.chains.v1beta1.SourceTxConfirmationEvent"></a>

### SourceTxConfirmationEvent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `amount` | [uint64](#uint64) |  |  |
| `asset` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `payload` | [bytes](#bytes) |  |  |
| `destination_contract_address` | [string](#string) |  |  |
| `destination_recipient_address` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.VoteEvents"></a>

### VoteEvents



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `events` | [Event](#scalar.chains.v1beta1.Event) | repeated |  |





 <!-- end messages -->


<a name="scalar.chains.v1beta1.Event.Status"></a>

### Event.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_CONFIRMED | 1 |  |
| STATUS_COMPLETED | 2 |  |
| STATUS_FAILED | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/chains/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/v1beta1/params.proto



<a name="scalar.chains.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `network_kind` | [uint32](#uint32) |  |  |
| `token_code` | [bytes](#bytes) |  |  |
| `burnable` | [bytes](#bytes) |  |  |
| `revote_locking_period` | [int64](#int64) |  |  |
| `chain_id` | [bytes](#bytes) |  |  |
| `voting_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `min_voter_count` | [int64](#int64) |  |  |
| `commands_gas_limit` | [uint32](#uint32) |  |  |
| `voting_grace_period` | [int64](#int64) |  |  |
| `end_blocker_limit` | [int64](#int64) |  |  |
| `transfer_limit` | [uint64](#uint64) |  |  |
| `metadata` | [Params.MetadataEntry](#scalar.chains.v1beta1.Params.MetadataEntry) | repeated |  |






<a name="scalar.chains.v1beta1.Params.MetadataEntry"></a>

### Params.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/chains/v1beta1/tokens.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/v1beta1/tokens.proto



<a name="scalar.chains.v1beta1.ERC20Deposit"></a>

### ERC20Deposit
ERC20Deposit contains information for an ERC20 deposit


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `amount` | [bytes](#bytes) |  |  |
| `asset` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `burner_address` | [bytes](#bytes) |  |  |
| `log_index` | [uint64](#uint64) |  |  |






<a name="scalar.chains.v1beta1.ERC20TokenMetadata"></a>

### ERC20TokenMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset` | [string](#string) |  |  |
| `chain_id` | [bytes](#bytes) |  |  |
| `details` | [TokenDetails](#scalar.chains.v1beta1.TokenDetails) |  |  |
| `token_address` | [string](#string) |  |  |
| `tx_hash` | [string](#string) |  |  |
| `status` | [Status](#scalar.chains.v1beta1.Status) |  |  |
| `is_external` | [bool](#bool) |  |  |
| `burner_code` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/utils/v1beta1/queuer.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/utils/v1beta1/queuer.proto



<a name="scalar.utils.v1beta1.QueueState"></a>

### QueueState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `items` | [QueueState.ItemsEntry](#scalar.utils.v1beta1.QueueState.ItemsEntry) | repeated |  |






<a name="scalar.utils.v1beta1.QueueState.Item"></a>

### QueueState.Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |
| `value` | [bytes](#bytes) |  |  |






<a name="scalar.utils.v1beta1.QueueState.ItemsEntry"></a>

### QueueState.ItemsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [QueueState.Item](#scalar.utils.v1beta1.QueueState.Item) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/chains/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/v1beta1/genesis.proto



<a name="scalar.chains.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chains` | [GenesisState.Chain](#scalar.chains.v1beta1.GenesisState.Chain) | repeated |  |






<a name="scalar.chains.v1beta1.GenesisState.Chain"></a>

### GenesisState.Chain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.chains.v1beta1.Params) |  |  |
| `command_queue` | [scalar.utils.v1beta1.QueueState](#scalar.utils.v1beta1.QueueState) |  |  |
| `confirmed_source_txs` | [SourceTx](#scalar.chains.v1beta1.SourceTx) | repeated |  |
| `command_batches` | [CommandBatchMetadata](#scalar.chains.v1beta1.CommandBatchMetadata) | repeated |  |
| `gateway` | [Gateway](#scalar.chains.v1beta1.Gateway) |  |  |
| `tokens` | [ERC20TokenMetadata](#scalar.chains.v1beta1.ERC20TokenMetadata) | repeated |  |
| `events` | [Event](#scalar.chains.v1beta1.Event) | repeated |  |
| `confirmed_event_queue` | [scalar.utils.v1beta1.QueueState](#scalar.utils.v1beta1.QueueState) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/chains/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/v1beta1/tx.proto



<a name="scalar.chains.v1beta1.AddChainRequest"></a>

### AddChainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `name` | [string](#string) |  |  |
| `params` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.AddChainResponse"></a>

### AddChainResponse







<a name="scalar.chains.v1beta1.ConfirmDepositRequest"></a>

### ConfirmDepositRequest
MsgConfirmDeposit represents an erc20 deposit confirmation message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `amount` | [bytes](#bytes) |  | **Deprecated.**  |
| `burner_address` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.ConfirmDepositResponse"></a>

### ConfirmDepositResponse







<a name="scalar.chains.v1beta1.ConfirmSourceTxsRequest"></a>

### ConfirmSourceTxsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_ids` | [bytes](#bytes) | repeated |  |






<a name="scalar.chains.v1beta1.ConfirmSourceTxsResponse"></a>

### ConfirmSourceTxsResponse







<a name="scalar.chains.v1beta1.ConfirmTokenRequest"></a>

### ConfirmTokenRequest
MsgConfirmToken represents a token deploy confirmation message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `asset` | [Asset](#scalar.chains.v1beta1.Asset) |  |  |






<a name="scalar.chains.v1beta1.ConfirmTokenResponse"></a>

### ConfirmTokenResponse







<a name="scalar.chains.v1beta1.ConfirmTransferKeyRequest"></a>

### ConfirmTransferKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.ConfirmTransferKeyResponse"></a>

### ConfirmTransferKeyResponse







<a name="scalar.chains.v1beta1.CreateBurnTokensRequest"></a>

### CreateBurnTokensRequest
CreateBurnTokensRequest represents the message to create commands to burn
tokens with scalarGateway


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.CreateBurnTokensResponse"></a>

### CreateBurnTokensResponse







<a name="scalar.chains.v1beta1.CreateDeployTokenRequest"></a>

### CreateDeployTokenRequest
CreateDeployTokenRequest represents the message to create a deploy token
command for scalarGateway


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `asset` | [Asset](#scalar.chains.v1beta1.Asset) |  |  |
| `token_details` | [TokenDetails](#scalar.chains.v1beta1.TokenDetails) |  |  |
| `address` | [bytes](#bytes) |  |  |
| `daily_mint_limit` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.CreateDeployTokenResponse"></a>

### CreateDeployTokenResponse







<a name="scalar.chains.v1beta1.CreatePendingTransfersRequest"></a>

### CreatePendingTransfersRequest
CreatePendingTransfersRequest represents a message to trigger the creation of
commands handling all pending transfers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.CreatePendingTransfersResponse"></a>

### CreatePendingTransfersResponse







<a name="scalar.chains.v1beta1.CreateTransferOperatorshipRequest"></a>

### CreateTransferOperatorshipRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.CreateTransferOperatorshipResponse"></a>

### CreateTransferOperatorshipResponse







<a name="scalar.chains.v1beta1.CreateTransferOwnershipRequest"></a>

### CreateTransferOwnershipRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.CreateTransferOwnershipResponse"></a>

### CreateTransferOwnershipResponse







<a name="scalar.chains.v1beta1.LinkRequest"></a>

### LinkRequest
MsgLink represents the message that links a cross chain address to a burner
address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `recipient_addr` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |
| `recipient_chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.LinkResponse"></a>

### LinkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit_addr` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.RetryFailedEventRequest"></a>

### RetryFailedEventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.RetryFailedEventResponse"></a>

### RetryFailedEventResponse







<a name="scalar.chains.v1beta1.SetGatewayRequest"></a>

### SetGatewayRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `address` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.SetGatewayResponse"></a>

### SetGatewayResponse







<a name="scalar.chains.v1beta1.SignBtcCommandsRequest"></a>

### SignBtcCommandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.SignCommandsRequest"></a>

### SignCommandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.SignCommandsResponse"></a>

### SignCommandsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batched_commands_id` | [bytes](#bytes) |  |  |
| `command_count` | [uint32](#uint32) |  |  |






<a name="scalar.chains.v1beta1.SignPsbtCommandRequest"></a>

### SignPsbtCommandRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `psbt` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.SignPsbtCommandResponse"></a>

### SignPsbtCommandResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batched_commands_id` | [bytes](#bytes) |  |  |
| `command_count` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/exported/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/exported/v1beta1/types.proto



<a name="scalar.covenant.exported.v1beta1.ListOfTapScriptSigsMap"></a>

### ListOfTapScriptSigsMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inner` | [TapScriptSigsMap](#scalar.covenant.exported.v1beta1.TapScriptSigsMap) | repeated |  |






<a name="scalar.covenant.exported.v1beta1.TapScriptSig"></a>

### TapScriptSig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_x_only` | [bytes](#bytes) |  |  |
| `leaf_hash` | [bytes](#bytes) |  |  |
| `signature` | [bytes](#bytes) |  |  |






<a name="scalar.covenant.exported.v1beta1.TapScriptSigsEntry"></a>

### TapScriptSigsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [uint64](#uint64) |  |  |
| `sigs` | [TapScriptSigsList](#scalar.covenant.exported.v1beta1.TapScriptSigsList) |  |  |






<a name="scalar.covenant.exported.v1beta1.TapScriptSigsList"></a>

### TapScriptSigsList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `list` | [TapScriptSig](#scalar.covenant.exported.v1beta1.TapScriptSig) | repeated |  |






<a name="scalar.covenant.exported.v1beta1.TapScriptSigsMap"></a>

### TapScriptSigsMap
The reason we use a list instead of a map is because the map is not ensured
the deterministic order of the entries


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inner` | [TapScriptSigsEntry](#scalar.covenant.exported.v1beta1.TapScriptSigsEntry) | repeated |  |





 <!-- end messages -->


<a name="scalar.covenant.exported.v1beta1.KeyState"></a>

### KeyState


| Name | Number | Description |
| ---- | ------ | ----------- |
| KEY_STATE_UNSPECIFIED | 0 |  |
| KEY_STATE_ASSIGNED | 1 |  |
| KEY_STATE_ACTIVE | 2 |  |



<a name="scalar.covenant.exported.v1beta1.PsbtState"></a>

### PsbtState


| Name | Number | Description |
| ---- | ------ | ----------- |
| PSBT_STATE_UNSPECIFIED | 0 |  |
| PSBT_STATE_PENDING | 1 |  |
| PSBT_STATE_CREATING | 2 |  |
| PSBT_STATE_SIGNING | 3 |  |
| PSBT_STATE_COMPLETED | 4 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/multisig/exported/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/exported/v1beta1/types.proto


 <!-- end messages -->


<a name="scalar.multisig.exported.v1beta1.KeyState"></a>

### KeyState


| Name | Number | Description |
| ---- | ------ | ----------- |
| KEY_STATE_UNSPECIFIED | 0 |  |
| KEY_STATE_ASSIGNED | 1 |  |
| KEY_STATE_ACTIVE | 2 |  |



<a name="scalar.multisig.exported.v1beta1.MultisigState"></a>

### MultisigState


| Name | Number | Description |
| ---- | ------ | ----------- |
| MULTISIG_STATE_UNSPECIFIED | 0 |  |
| MULTISIG_STATE_PENDING | 1 |  |
| MULTISIG_STATE_COMPLETED | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/multisig/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/v1beta1/types.proto



<a name="scalar.multisig.v1beta1.Key"></a>

### Key



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `snapshot` | [scalar.snapshot.exported.v1beta1.Snapshot](#scalar.snapshot.exported.v1beta1.Snapshot) |  |  |
| `pub_keys` | [Key.PubKeysEntry](#scalar.multisig.v1beta1.Key.PubKeysEntry) | repeated |  |
| `signing_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `state` | [scalar.multisig.exported.v1beta1.KeyState](#scalar.multisig.exported.v1beta1.KeyState) |  |  |






<a name="scalar.multisig.v1beta1.Key.PubKeysEntry"></a>

### Key.PubKeysEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.KeyEpoch"></a>

### KeyEpoch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `epoch` | [uint64](#uint64) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeygenSession"></a>

### KeygenSession



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [Key](#scalar.multisig.v1beta1.Key) |  |  |
| `state` | [scalar.multisig.exported.v1beta1.MultisigState](#scalar.multisig.exported.v1beta1.MultisigState) |  |  |
| `keygen_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `expires_at` | [int64](#int64) |  |  |
| `completed_at` | [int64](#int64) |  |  |
| `is_pub_key_received` | [KeygenSession.IsPubKeyReceivedEntry](#scalar.multisig.v1beta1.KeygenSession.IsPubKeyReceivedEntry) | repeated |  |
| `grace_period` | [int64](#int64) |  |  |






<a name="scalar.multisig.v1beta1.KeygenSession.IsPubKeyReceivedEntry"></a>

### KeygenSession.IsPubKeyReceivedEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [bool](#bool) |  |  |






<a name="scalar.multisig.v1beta1.MultiSig"></a>

### MultiSig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `sigs` | [MultiSig.SigsEntry](#scalar.multisig.v1beta1.MultiSig.SigsEntry) | repeated |  |






<a name="scalar.multisig.v1beta1.MultiSig.SigsEntry"></a>

### MultiSig.SigsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.SigningSession"></a>

### SigningSession



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `multi_sig` | [MultiSig](#scalar.multisig.v1beta1.MultiSig) |  |  |
| `state` | [scalar.multisig.exported.v1beta1.MultisigState](#scalar.multisig.exported.v1beta1.MultisigState) |  |  |
| `key` | [Key](#scalar.multisig.v1beta1.Key) |  |  |
| `expires_at` | [int64](#int64) |  |  |
| `completed_at` | [int64](#int64) |  |  |
| `grace_period` | [int64](#int64) |  |  |
| `module` | [string](#string) |  |  |
| `module_metadata` | [google.protobuf.Any](#google.protobuf.Any) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/types.proto



<a name="scalar.covenant.v1beta1.Custodian"></a>

### Custodian
Custodian represents an individual custodian configuration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | e.g., "Custodian1" |
| `val_address` | [string](#string) |  | e.g., "scalarvaloper1..." |
| `bitcoin_pubkey` | [bytes](#bytes) |  | e.g., |
| `status` | [Status](#scalar.covenant.v1beta1.Status) |  | "0215da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e1641488"

Whether the custodian is active |
| `description` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.CustodianGroup"></a>

### CustodianGroup
CustodianGroup represents a group of custodians with their configuration
uid is used as identity of the group, btc_pubkey is change by list of
custodians


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `uid` | [string](#string) |  | the UID is unique, to distinguish between custodian groups |
| `name` | [string](#string) |  | e.g., "All" |
| `bitcoin_pubkey` | [bytes](#bytes) |  | e.g., |
| `quorum` | [uint32](#uint32) |  | "tb1p07q440mdl4uyywns325dk8pvjphwety3psp4zvkngtjf3z3hhr2sfar3hv"

quorum threshold e.g.,3 |
| `status` | [Status](#scalar.covenant.v1beta1.Status) |  | Whether the custodian is active |
| `description` | [string](#string) |  |  |
| `custodians` | [Custodian](#scalar.covenant.v1beta1.Custodian) | repeated |  |






<a name="scalar.covenant.v1beta1.PsbtMultiSig"></a>

### PsbtMultiSig



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |
| `multi_psbt` | [bytes](#bytes) | repeated |  |
| `participant_list_tap_script_sigs` | [PsbtMultiSig.ParticipantListTapScriptSigsEntry](#scalar.covenant.v1beta1.PsbtMultiSig.ParticipantListTapScriptSigsEntry) | repeated |  |
| `finalized_txs` | [bytes](#bytes) | repeated |  |






<a name="scalar.covenant.v1beta1.PsbtMultiSig.ParticipantListTapScriptSigsEntry"></a>

### PsbtMultiSig.ParticipantListTapScriptSigsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [scalar.covenant.exported.v1beta1.ListOfTapScriptSigsMap](#scalar.covenant.exported.v1beta1.ListOfTapScriptSigsMap) |  |  |






<a name="scalar.covenant.v1beta1.SigningSession"></a>

### SigningSession



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `psbt_multi_sig` | [PsbtMultiSig](#scalar.covenant.v1beta1.PsbtMultiSig) |  |  |
| `state` | [scalar.covenant.exported.v1beta1.PsbtState](#scalar.covenant.exported.v1beta1.PsbtState) |  |  |
| `key` | [scalar.multisig.v1beta1.Key](#scalar.multisig.v1beta1.Key) |  |  |
| `expires_at` | [int64](#int64) |  |  |
| `completed_at` | [int64](#int64) |  |  |
| `grace_period` | [int64](#int64) |  |  |
| `module` | [string](#string) |  |  |
| `module_metadata` | [google.protobuf.Any](#google.protobuf.Any) |  |  |





 <!-- end messages -->


<a name="scalar.covenant.v1beta1.Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_ACTIVATED | 1 |  |
| STATUS_DEACTIVATED | 2 |  |
| STATUS_PENDING | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/protocol/exported/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/exported/v1beta1/types.proto



<a name="scalar.protocol.exported.v1beta1.MinorAddress"></a>

### MinorAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain_name` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="scalar.protocol.exported.v1beta1.ProtocolAttributes"></a>

### ProtocolAttributes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `model` | [LiquidityModel](#scalar.protocol.exported.v1beta1.LiquidityModel) |  |  |






<a name="scalar.protocol.exported.v1beta1.ProtocolInfo"></a>

### ProtocolInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `custodians_group_uid` | [string](#string) |  | string key_id = 1 [ (gogoproto.customname) = "KeyID", (gogoproto.casttype) = "github.com/scalarorg/scalar-core/x/multisig/exported.KeyID" ]; bytes custodians_pubkey = 2 [ (gogoproto.customname) = "CustodiansPubkey" ]; |
| `liquidity_model` | [LiquidityModel](#scalar.protocol.exported.v1beta1.LiquidityModel) |  |  |
| `symbol` | [string](#string) |  |  |
| `origin_chain` | [string](#string) |  |  |
| `minor_addresses` | [MinorAddress](#scalar.protocol.exported.v1beta1.MinorAddress) | repeated |  |






<a name="scalar.protocol.exported.v1beta1.SupportedChain"></a>

### SupportedChain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `address` | [string](#string) |  | Asset address on the chain |





 <!-- end messages -->


<a name="scalar.protocol.exported.v1beta1.LiquidityModel"></a>

### LiquidityModel


| Name | Number | Description |
| ---- | ------ | ----------- |
| LIQUIDITY_MODEL_UNSPECIFIED | 0 |  |
| LIQUIDITY_MODEL_POOL | 1 |  |
| LIQUIDITY_MODEL_UPC | 2 |  |



<a name="scalar.protocol.exported.v1beta1.Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_ACTIVATED | 1 |  |
| STATUS_DEACTIVATED | 2 |  |
| STATUS_PENDING | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/chains/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/v1beta1/query.proto



<a name="scalar.chains.v1beta1.BatchedCommandsRequest"></a>

### BatchedCommandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `id` | [string](#string) |  | id defines an optional id for the commandsbatch. If not specified the latest will be returned |






<a name="scalar.chains.v1beta1.BatchedCommandsResponse"></a>

### BatchedCommandsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |
| `status` | [BatchedCommandsStatus](#scalar.chains.v1beta1.BatchedCommandsStatus) |  |  |
| `key_id` | [string](#string) |  |  |
| `execute_data` | [string](#string) |  |  |
| `prev_batched_commands_id` | [string](#string) |  |  |
| `command_ids` | [string](#string) | repeated |  |
| `proof` | [Proof](#scalar.chains.v1beta1.Proof) |  |  |






<a name="scalar.chains.v1beta1.BurnerInfoRequest"></a>

### BurnerInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.BurnerInfoResponse"></a>

### BurnerInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `burner_info` | [BurnerInfo](#scalar.chains.v1beta1.BurnerInfo) |  |  |






<a name="scalar.chains.v1beta1.BytecodeRequest"></a>

### BytecodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `contract` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.BytecodeResponse"></a>

### BytecodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bytecode` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ChainsRequest"></a>

### ChainsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [ChainStatus](#scalar.chains.v1beta1.ChainStatus) |  |  |






<a name="scalar.chains.v1beta1.ChainsResponse"></a>

### ChainsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chains` | [string](#string) | repeated |  |






<a name="scalar.chains.v1beta1.CommandRequest"></a>

### CommandRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `id` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.CommandResponse"></a>

### CommandResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `params` | [CommandResponse.ParamsEntry](#scalar.chains.v1beta1.CommandResponse.ParamsEntry) | repeated |  |
| `key_id` | [string](#string) |  |  |
| `max_gas_cost` | [uint32](#uint32) |  |  |






<a name="scalar.chains.v1beta1.CommandResponse.ParamsEntry"></a>

### CommandResponse.ParamsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ConfirmationHeightRequest"></a>

### ConfirmationHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ConfirmationHeightResponse"></a>

### ConfirmationHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [uint64](#uint64) |  |  |






<a name="scalar.chains.v1beta1.DepositStateRequest"></a>

### DepositStateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `params` | [QueryDepositStateParams](#scalar.chains.v1beta1.QueryDepositStateParams) |  |  |






<a name="scalar.chains.v1beta1.DepositStateResponse"></a>

### DepositStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [DepositStatus](#scalar.chains.v1beta1.DepositStatus) |  |  |






<a name="scalar.chains.v1beta1.ERC20TokensRequest"></a>

### ERC20TokensRequest
ERC20TokensRequest describes the chain for which the type of ERC20 tokens are
requested.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `type` | [TokenType](#scalar.chains.v1beta1.TokenType) |  |  |






<a name="scalar.chains.v1beta1.ERC20TokensResponse"></a>

### ERC20TokensResponse
ERC20TokensResponse describes the asset and symbol for all
ERC20 tokens requested for a chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tokens` | [ERC20TokensResponse.Token](#scalar.chains.v1beta1.ERC20TokensResponse.Token) | repeated |  |






<a name="scalar.chains.v1beta1.ERC20TokensResponse.Token"></a>

### ERC20TokensResponse.Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `event` | [Event](#scalar.chains.v1beta1.Event) |  |  |






<a name="scalar.chains.v1beta1.GatewayAddressRequest"></a>

### GatewayAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.GatewayAddressResponse"></a>

### GatewayAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.KeyAddressRequest"></a>

### KeyAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.KeyAddressResponse"></a>

### KeyAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |
| `addresses` | [KeyAddressResponse.WeightedAddress](#scalar.chains.v1beta1.KeyAddressResponse.WeightedAddress) | repeated |  |
| `threshold` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.KeyAddressResponse.WeightedAddress"></a>

### KeyAddressResponse.WeightedAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `weight` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.chains.v1beta1.Params) |  |  |






<a name="scalar.chains.v1beta1.PendingCommandsRequest"></a>

### PendingCommandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.PendingCommandsResponse"></a>

### PendingCommandsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commands` | [QueryCommandResponse](#scalar.chains.v1beta1.QueryCommandResponse) | repeated |  |






<a name="scalar.chains.v1beta1.QueryBurnerAddressResponse"></a>

### QueryBurnerAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.QueryCommandResponse"></a>

### QueryCommandResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `params` | [QueryCommandResponse.ParamsEntry](#scalar.chains.v1beta1.QueryCommandResponse.ParamsEntry) | repeated |  |
| `key_id` | [string](#string) |  |  |
| `max_gas_cost` | [uint32](#uint32) |  |  |
| `payload` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.QueryCommandResponse.ParamsEntry"></a>

### QueryCommandResponse.ParamsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.QueryDepositStateParams"></a>

### QueryDepositStateParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `burner_address` | [bytes](#bytes) |  |  |






<a name="scalar.chains.v1beta1.QueryTokenAddressResponse"></a>

### QueryTokenAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `confirmed` | [bool](#bool) |  |  |






<a name="scalar.chains.v1beta1.TokenInfoRequest"></a>

### TokenInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="scalar.chains.v1beta1.TokenInfoResponse"></a>

### TokenInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset` | [string](#string) |  |  |
| `details` | [TokenDetails](#scalar.chains.v1beta1.TokenDetails) |  |  |
| `address` | [string](#string) |  |  |
| `confirmed` | [bool](#bool) |  |  |
| `is_external` | [bool](#bool) |  |  |
| `burner_code_hash` | [string](#string) |  |  |





 <!-- end messages -->


<a name="scalar.chains.v1beta1.ChainStatus"></a>

### ChainStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHAIN_STATUS_UNSPECIFIED | 0 |  |
| CHAIN_STATUS_ACTIVATED | 1 |  |
| CHAIN_STATUS_DEACTIVATED | 2 |  |



<a name="scalar.chains.v1beta1.TokenType"></a>

### TokenType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TOKEN_TYPE_UNSPECIFIED | 0 |  |
| TOKEN_TYPE_INTERNAL | 1 |  |
| TOKEN_TYPE_EXTERNAL | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/chains/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/chains/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.chains.v1beta1.MsgService"></a>

### MsgService
Msg defines the btc Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ConfirmSourceTxs` | [ConfirmSourceTxsRequest](#scalar.chains.v1beta1.ConfirmSourceTxsRequest) | [ConfirmSourceTxsResponse](#scalar.chains.v1beta1.ConfirmSourceTxsResponse) |  | POST|/scalar/chains/v1beta1/confirm_source_txs|
| `SetGateway` | [SetGatewayRequest](#scalar.chains.v1beta1.SetGatewayRequest) | [SetGatewayResponse](#scalar.chains.v1beta1.SetGatewayResponse) |  | POST|/scalar/chains/v1beta1/set_gateway|
| `Link` | [LinkRequest](#scalar.chains.v1beta1.LinkRequest) | [LinkResponse](#scalar.chains.v1beta1.LinkResponse) |  | POST|/scalar/chains/v1beta1/link|
| `ConfirmToken` | [ConfirmTokenRequest](#scalar.chains.v1beta1.ConfirmTokenRequest) | [ConfirmTokenResponse](#scalar.chains.v1beta1.ConfirmTokenResponse) |  | POST|/scalar/chains/v1beta1/confirm_token|
| `ConfirmDeposit` | [ConfirmDepositRequest](#scalar.chains.v1beta1.ConfirmDepositRequest) | [ConfirmDepositResponse](#scalar.chains.v1beta1.ConfirmDepositResponse) |  | POST|/scalar/chains/v1beta1/confirm_deposit|
| `ConfirmTransferKey` | [ConfirmTransferKeyRequest](#scalar.chains.v1beta1.ConfirmTransferKeyRequest) | [ConfirmTransferKeyResponse](#scalar.chains.v1beta1.ConfirmTransferKeyResponse) |  | POST|/scalar/chains/confirm_transfer_key|
| `CreateDeployToken` | [CreateDeployTokenRequest](#scalar.chains.v1beta1.CreateDeployTokenRequest) | [CreateDeployTokenResponse](#scalar.chains.v1beta1.CreateDeployTokenResponse) |  | POST|/scalar/chains/v1beta1/create_deploy_token|
| `CreateBurnTokens` | [CreateBurnTokensRequest](#scalar.chains.v1beta1.CreateBurnTokensRequest) | [CreateBurnTokensResponse](#scalar.chains.v1beta1.CreateBurnTokensResponse) |  | POST|/scalar/chains/v1beta1/create_burn_tokens|
| `CreatePendingTransfers` | [CreatePendingTransfersRequest](#scalar.chains.v1beta1.CreatePendingTransfersRequest) | [CreatePendingTransfersResponse](#scalar.chains.v1beta1.CreatePendingTransfersResponse) |  | POST|/scalar/chains/v1beta1/create_pending_transfers|
| `CreateTransferOperatorship` | [CreateTransferOperatorshipRequest](#scalar.chains.v1beta1.CreateTransferOperatorshipRequest) | [CreateTransferOperatorshipResponse](#scalar.chains.v1beta1.CreateTransferOperatorshipResponse) |  | POST|/scalar/chains/v1beta1/create_transfer_operatorship|
| `SignCommands` | [SignCommandsRequest](#scalar.chains.v1beta1.SignCommandsRequest) | [SignCommandsResponse](#scalar.chains.v1beta1.SignCommandsResponse) |  | POST|/scalar/chains/v1beta1/sign_commands|
| `SignBtcCommand` | [SignBtcCommandsRequest](#scalar.chains.v1beta1.SignBtcCommandsRequest) | [SignCommandsResponse](#scalar.chains.v1beta1.SignCommandsResponse) |  | POST|/scalar/chains/v1beta1/sign_btc_commands|
| `SignPsbtCommand` | [SignPsbtCommandRequest](#scalar.chains.v1beta1.SignPsbtCommandRequest) | [SignPsbtCommandResponse](#scalar.chains.v1beta1.SignPsbtCommandResponse) |  | POST|/scalar/chains/v1beta1/sign_btc_commands|
| `AddChain` | [AddChainRequest](#scalar.chains.v1beta1.AddChainRequest) | [AddChainResponse](#scalar.chains.v1beta1.AddChainResponse) |  | POST|/scalar/chains/v1beta1/add_chain|
| `RetryFailedEvent` | [RetryFailedEventRequest](#scalar.chains.v1beta1.RetryFailedEventRequest) | [RetryFailedEventResponse](#scalar.chains.v1beta1.RetryFailedEventResponse) |  | POST|/scalar/chains/v1beta1/retry-failed-event|


<a name="scalar.chains.v1beta1.QueryService"></a>

### QueryService
QueryService defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `BatchedCommands` | [BatchedCommandsRequest](#scalar.chains.v1beta1.BatchedCommandsRequest) | [BatchedCommandsResponse](#scalar.chains.v1beta1.BatchedCommandsResponse) | BatchedCommands queries the batched commands for a specified chain and BatchedCommandsID if no BatchedCommandsID is specified, then it returns the latest batched commands | GET|/scalar/chains/v1beta1/batched_commands/{chain}/{id}|
| `BurnerInfo` | [BurnerInfoRequest](#scalar.chains.v1beta1.BurnerInfoRequest) | [BurnerInfoResponse](#scalar.chains.v1beta1.BurnerInfoResponse) |  | GET|/scalar/chains/v1beta1/burner_info|
| `ConfirmationHeight` | [ConfirmationHeightRequest](#scalar.chains.v1beta1.ConfirmationHeightRequest) | [ConfirmationHeightResponse](#scalar.chains.v1beta1.ConfirmationHeightResponse) | ConfirmationHeight queries the confirmation height for the specified chain | GET|/scalar/chains/v1beta1/confirmation_height/{chain}|
| `PendingCommands` | [PendingCommandsRequest](#scalar.chains.v1beta1.PendingCommandsRequest) | [PendingCommandsResponse](#scalar.chains.v1beta1.PendingCommandsResponse) | PendingCommands queries the pending commands for the specified chain | GET|/scalar/chains/v1beta1/pending_commands/{chain}|
| `Chains` | [ChainsRequest](#scalar.chains.v1beta1.ChainsRequest) | [ChainsResponse](#scalar.chains.v1beta1.ChainsResponse) | Chains queries the available chains | GET|/scalar/chains/v1beta1/chains|
| `Command` | [CommandRequest](#scalar.chains.v1beta1.CommandRequest) | [CommandResponse](#scalar.chains.v1beta1.CommandResponse) | Command queries the command of a chain provided the command id | GET|/scalar/chains/v1beta1/command_request|
| `KeyAddress` | [KeyAddressRequest](#scalar.chains.v1beta1.KeyAddressRequest) | [KeyAddressResponse](#scalar.chains.v1beta1.KeyAddressResponse) | KeyAddress queries the address of key of a chain | GET|/scalar/chains/v1beta1/key_address/{chain}|
| `GatewayAddress` | [GatewayAddressRequest](#scalar.chains.v1beta1.GatewayAddressRequest) | [GatewayAddressResponse](#scalar.chains.v1beta1.GatewayAddressResponse) | GatewayAddress queries the address of scalar gateway at the specified chain | GET|/scalar/chains/v1beta1/gateway_address/{chain}|
| `Bytecode` | [BytecodeRequest](#scalar.chains.v1beta1.BytecodeRequest) | [BytecodeResponse](#scalar.chains.v1beta1.BytecodeResponse) | Bytecode queries the bytecode of a specified gateway at the specified chain | GET|/scalar/chains/v1beta1/bytecode/{chain}/{contract}|
| `Event` | [EventRequest](#scalar.chains.v1beta1.EventRequest) | [EventResponse](#scalar.chains.v1beta1.EventResponse) | Event queries an event at the specified chain | GET|/scalar/chains/v1beta1/event/{chain}/{event_id}|
| `ERC20Tokens` | [ERC20TokensRequest](#scalar.chains.v1beta1.ERC20TokensRequest) | [ERC20TokensResponse](#scalar.chains.v1beta1.ERC20TokensResponse) | ERC20Tokens queries the ERC20 tokens registered for a chain | GET|/scalar/chains/v1beta1/erc20_tokens/{chain}|
| `TokenInfo` | [TokenInfoRequest](#scalar.chains.v1beta1.TokenInfoRequest) | [TokenInfoResponse](#scalar.chains.v1beta1.TokenInfoResponse) | TokenInfo queries the token info for a registered ERC20 Token | GET|/scalar/chains/v1beta1/token_info/{chain}|
| `Params` | [ParamsRequest](#scalar.chains.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.chains.v1beta1.ParamsResponse) |  | GET|/scalar/chains/v1beta1/params/{chain}|

 <!-- end services -->



<a name="scalar/covenant/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/events.proto



<a name="scalar.covenant.v1beta1.KeyRotated"></a>

### KeyRotated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.SigningPsbtCompleted"></a>

### SigningPsbtCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |






<a name="scalar.covenant.v1beta1.SigningPsbtExpired"></a>

### SigningPsbtExpired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |






<a name="scalar.covenant.v1beta1.SigningPsbtStarted"></a>

### SigningPsbtStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `chain` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |
| `key_id` | [string](#string) |  |  |
| `pub_keys` | [SigningPsbtStarted.PubKeysEntry](#scalar.covenant.v1beta1.SigningPsbtStarted.PubKeysEntry) | repeated |  |
| `requesting_module` | [string](#string) |  |  |
| `multi_psbt` | [bytes](#bytes) | repeated |  |






<a name="scalar.covenant.v1beta1.SigningPsbtStarted.PubKeysEntry"></a>

### SigningPsbtStarted.PubKeysEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [bytes](#bytes) |  |  |






<a name="scalar.covenant.v1beta1.TapScriptSigsSubmitted"></a>

### TapScriptSigsSubmitted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |
| `participant` | [bytes](#bytes) |  |  |
| `list_of_tap_script_sigs_map` | [scalar.covenant.exported.v1beta1.TapScriptSigsMap](#scalar.covenant.exported.v1beta1.TapScriptSigsMap) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/params.proto



<a name="scalar.covenant.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signing_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `signing_timeout` | [int64](#int64) |  |  |
| `signing_grace_period` | [int64](#int64) |  |  |
| `active_epoch_count` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/genesis.proto



<a name="scalar.covenant.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.covenant.v1beta1.Params) |  |  |
| `signing_sessions` | [SigningSession](#scalar.covenant.v1beta1.SigningSession) | repeated |  |
| `custodians` | [Custodian](#scalar.covenant.v1beta1.Custodian) | repeated |  |
| `groups` | [CustodianGroup](#scalar.covenant.v1beta1.CustodianGroup) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/multisig/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/v1beta1/params.proto



<a name="scalar.multisig.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keygen_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `signing_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `keygen_timeout` | [int64](#int64) |  |  |
| `keygen_grace_period` | [int64](#int64) |  |  |
| `signing_timeout` | [int64](#int64) |  |  |
| `signing_grace_period` | [int64](#int64) |  |  |
| `active_epoch_count` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/multisig/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/v1beta1/query.proto



<a name="scalar.multisig.v1beta1.KeyIDRequest"></a>

### KeyIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeyIDResponse"></a>

### KeyIDResponse
KeyIDResponse contains the key ID of the key assigned to a given chain.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeyRequest"></a>

### KeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeyResponse"></a>

### KeyResponse
KeyResponse contains the key corresponding to a given key id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |
| `state` | [scalar.multisig.exported.v1beta1.KeyState](#scalar.multisig.exported.v1beta1.KeyState) |  |  |
| `started_at` | [int64](#int64) |  |  |
| `started_at_timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `threshold_weight` | [bytes](#bytes) |  |  |
| `bonded_weight` | [bytes](#bytes) |  |  |
| `participants` | [KeygenParticipant](#scalar.multisig.v1beta1.KeygenParticipant) | repeated | Keygen participants in descending order by weight |






<a name="scalar.multisig.v1beta1.KeygenParticipant"></a>

### KeygenParticipant



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `weight` | [bytes](#bytes) |  |  |
| `pub_key` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeygenSessionRequest"></a>

### KeygenSessionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeygenSessionResponse"></a>

### KeygenSessionResponse
KeygenSessionResponse contains the keygen session info for a given key ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `started_at` | [int64](#int64) |  |  |
| `started_at_timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `expires_at` | [int64](#int64) |  |  |
| `completed_at` | [int64](#int64) |  |  |
| `grace_period` | [int64](#int64) |  |  |
| `state` | [scalar.multisig.exported.v1beta1.MultisigState](#scalar.multisig.exported.v1beta1.MultisigState) |  |  |
| `keygen_threshold_weight` | [bytes](#bytes) |  |  |
| `signing_threshold_weight` | [bytes](#bytes) |  |  |
| `bonded_weight` | [bytes](#bytes) |  |  |
| `participants` | [KeygenParticipant](#scalar.multisig.v1beta1.KeygenParticipant) | repeated | Keygen candidates in descending order by weight |






<a name="scalar.multisig.v1beta1.NextKeyIDRequest"></a>

### NextKeyIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.NextKeyIDResponse"></a>

### NextKeyIDResponse
NextKeyIDResponse contains the key ID for the next rotation on the given
chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.multisig.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.multisig.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/query.proto



<a name="scalar.covenant.v1beta1.CustodiansRequest"></a>

### CustodiansRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `pubkey` | [bytes](#bytes) |  |  |
| `status` | [Status](#scalar.covenant.v1beta1.Status) |  |  |






<a name="scalar.covenant.v1beta1.CustodiansResponse"></a>

### CustodiansResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `custodians` | [Custodian](#scalar.covenant.v1beta1.Custodian) | repeated |  |






<a name="scalar.covenant.v1beta1.GroupsRequest"></a>

### GroupsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `uid` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.GroupsResponse"></a>

### GroupsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `groups` | [CustodianGroup](#scalar.covenant.v1beta1.CustodianGroup) | repeated |  |






<a name="scalar.covenant.v1beta1.KeyRequest"></a>

### KeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.KeyResponse"></a>

### KeyResponse
KeyResponse contains the key corresponding to a given key id.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |
| `state` | [scalar.multisig.exported.v1beta1.KeyState](#scalar.multisig.exported.v1beta1.KeyState) |  |  |
| `started_at` | [int64](#int64) |  |  |
| `started_at_timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `threshold_weight` | [bytes](#bytes) |  |  |
| `bonded_weight` | [bytes](#bytes) |  |  |
| `participants` | [scalar.multisig.v1beta1.KeygenParticipant](#scalar.multisig.v1beta1.KeygenParticipant) | repeated | Keygen participants in descending order by weight |






<a name="scalar.covenant.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.covenant.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.covenant.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/tx.proto



<a name="scalar.covenant.v1beta1.AddCustodianToGroupRequest"></a>

### AddCustodianToGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `guid` | [string](#string) |  | CustodianGroup uid |
| `custodian_pubkey` | [bytes](#bytes) |  |  |
| `description` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.CreateCustodianGroupRequest"></a>

### CreateCustodianGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `uid` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `quorum` | [uint32](#uint32) |  |  |
| `status` | [Status](#scalar.covenant.v1beta1.Status) |  |  |
| `description` | [string](#string) |  |  |
| `custodian` | [bytes](#bytes) | repeated |  |






<a name="scalar.covenant.v1beta1.CreateCustodianGroupResponse"></a>

### CreateCustodianGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group` | [CustodianGroup](#scalar.covenant.v1beta1.CustodianGroup) |  |  |






<a name="scalar.covenant.v1beta1.CreateCustodianRequest"></a>

### CreateCustodianRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `name` | [string](#string) |  |  |
| `bitcoin_pubkey` | [bytes](#bytes) |  |  |
| `status` | [Status](#scalar.covenant.v1beta1.Status) |  |  |
| `description` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.CreateCustodianResponse"></a>

### CreateCustodianResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `custodian` | [Custodian](#scalar.covenant.v1beta1.Custodian) |  |  |






<a name="scalar.covenant.v1beta1.CustodianToGroupResponse"></a>

### CustodianToGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group` | [CustodianGroup](#scalar.covenant.v1beta1.CustodianGroup) |  |  |






<a name="scalar.covenant.v1beta1.RemoveCustodianFromGroupRequest"></a>

### RemoveCustodianFromGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `guid` | [string](#string) |  | CustodianGroup uid |
| `custodian_pubkey` | [bytes](#bytes) |  |  |
| `description` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.RotateKeyRequest"></a>

### RotateKeyRequest
Rotate key for custodian group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.RotateKeyResponse"></a>

### RotateKeyResponse







<a name="scalar.covenant.v1beta1.SubmitTapScriptSigsRequest"></a>

### SubmitTapScriptSigsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |
| `list_of_tap_script_sigs_map` | [scalar.covenant.exported.v1beta1.TapScriptSigsMap](#scalar.covenant.exported.v1beta1.TapScriptSigsMap) | repeated |  |






<a name="scalar.covenant.v1beta1.SubmitTapScriptSigsResponse"></a>

### SubmitTapScriptSigsResponse







<a name="scalar.covenant.v1beta1.UpdateCustodianGroupRequest"></a>

### UpdateCustodianGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `uid` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `quorum` | [uint32](#uint32) |  |  |
| `status` | [Status](#scalar.covenant.v1beta1.Status) |  |  |
| `description` | [string](#string) |  |  |
| `custodian` | [bytes](#bytes) | repeated |  |






<a name="scalar.covenant.v1beta1.UpdateCustodianGroupResponse"></a>

### UpdateCustodianGroupResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group` | [CustodianGroup](#scalar.covenant.v1beta1.CustodianGroup) |  |  |






<a name="scalar.covenant.v1beta1.UpdateCustodianRequest"></a>

### UpdateCustodianRequest
Pubkey used as key for lookup custodian to update other values


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `bitcoin_pubkey` | [bytes](#bytes) |  |  |
| `name` | [string](#string) |  |  |
| `status` | [Status](#scalar.covenant.v1beta1.Status) |  |  |
| `description` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.UpdateCustodianResponse"></a>

### UpdateCustodianResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `custodian` | [Custodian](#scalar.covenant.v1beta1.Custodian) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.covenant.v1beta1.MsgService"></a>

### MsgService


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateCustodian` | [CreateCustodianRequest](#scalar.covenant.v1beta1.CreateCustodianRequest) | [CreateCustodianResponse](#scalar.covenant.v1beta1.CreateCustodianResponse) | Create custodian | POST|/scalar/covenant/v1beta1/create_custodian|
| `UpdateCustodian` | [UpdateCustodianRequest](#scalar.covenant.v1beta1.UpdateCustodianRequest) | [UpdateCustodianResponse](#scalar.covenant.v1beta1.UpdateCustodianResponse) | Update custodian | POST|/scalar/covenant/v1beta1/update_custodian|
| `CreateCustodianGroup` | [CreateCustodianGroupRequest](#scalar.covenant.v1beta1.CreateCustodianGroupRequest) | [CreateCustodianGroupResponse](#scalar.covenant.v1beta1.CreateCustodianGroupResponse) | Create custodian group | POST|/scalar/covenant/v1beta1/create_custodian_group|
| `UpdateCustodianGroup` | [UpdateCustodianGroupRequest](#scalar.covenant.v1beta1.UpdateCustodianGroupRequest) | [UpdateCustodianGroupResponse](#scalar.covenant.v1beta1.UpdateCustodianGroupResponse) | Update Custodian group | POST|/scalar/covenant/v1beta1/update_custodian_group|
| `AddCustodianToGroup` | [AddCustodianToGroupRequest](#scalar.covenant.v1beta1.AddCustodianToGroupRequest) | [CustodianToGroupResponse](#scalar.covenant.v1beta1.CustodianToGroupResponse) | Add Custodian to custodian group recalculate taproot pubkey when adding custodian to custodian group | POST|/scalar/covenant/v1beta1/add_custodian_to_group|
| `RemoveCustodianFromGroup` | [RemoveCustodianFromGroupRequest](#scalar.covenant.v1beta1.RemoveCustodianFromGroupRequest) | [CustodianToGroupResponse](#scalar.covenant.v1beta1.CustodianToGroupResponse) | Remove Custodian from custodian group recalculate taproot address when deleting custodian from custodian group | POST|/scalar/covenant/v1beta1/remove_custodian_from_group|
| `RotateKey` | [RotateKeyRequest](#scalar.covenant.v1beta1.RotateKeyRequest) | [RotateKeyResponse](#scalar.covenant.v1beta1.RotateKeyResponse) |  | POST|/scalar/covenant/v1beta1/rotate_key|
| `SubmitTapScriptSigs` | [SubmitTapScriptSigsRequest](#scalar.covenant.v1beta1.SubmitTapScriptSigsRequest) | [SubmitTapScriptSigsResponse](#scalar.covenant.v1beta1.SubmitTapScriptSigsResponse) |  | POST|/scalar/covenant/v1beta1/submit_tap_script_sigs|


<a name="scalar.covenant.v1beta1.QueryService"></a>

### QueryService


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Custodians` | [CustodiansRequest](#scalar.covenant.v1beta1.CustodiansRequest) | [CustodiansResponse](#scalar.covenant.v1beta1.CustodiansResponse) | Get custodians | GET|/scalar/convenant/v1beta1/custodians|
| `Groups` | [GroupsRequest](#scalar.covenant.v1beta1.GroupsRequest) | [GroupsResponse](#scalar.covenant.v1beta1.GroupsResponse) | Get custodian groups | GET|/scalar/covenant/v1beta1/custodian_groups|
| `Params` | [ParamsRequest](#scalar.covenant.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.covenant.v1beta1.ParamsResponse) |  | GET|/scalar/covenant/v1beta1/params|

 <!-- end services -->



<a name="scalar/multisig/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/v1beta1/events.proto



<a name="scalar.multisig.v1beta1.KeyAssigned"></a>

### KeyAssigned



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeyRotated"></a>

### KeyRotated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeygenCompleted"></a>

### KeygenCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeygenExpired"></a>

### KeygenExpired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.KeygenOptIn"></a>

### KeygenOptIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `participant` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.KeygenOptOut"></a>

### KeygenOptOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `participant` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.KeygenStarted"></a>

### KeygenStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |
| `participants` | [bytes](#bytes) | repeated |  |






<a name="scalar.multisig.v1beta1.PubKeySubmitted"></a>

### PubKeySubmitted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |
| `participant` | [bytes](#bytes) |  |  |
| `pub_key` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.SignatureSubmitted"></a>

### SignatureSubmitted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |
| `participant` | [bytes](#bytes) |  |  |
| `signature` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.SigningCompleted"></a>

### SigningCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |






<a name="scalar.multisig.v1beta1.SigningExpired"></a>

### SigningExpired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |






<a name="scalar.multisig.v1beta1.SigningStarted"></a>

### SigningStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |
| `key_id` | [string](#string) |  |  |
| `pub_keys` | [SigningStarted.PubKeysEntry](#scalar.multisig.v1beta1.SigningStarted.PubKeysEntry) | repeated |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `requesting_module` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.SigningStarted.PubKeysEntry"></a>

### SigningStarted.PubKeysEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/multisig/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/v1beta1/genesis.proto



<a name="scalar.multisig.v1beta1.GenesisState"></a>

### GenesisState
GenesisState represents the genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.multisig.v1beta1.Params) |  |  |
| `keygen_sessions` | [KeygenSession](#scalar.multisig.v1beta1.KeygenSession) | repeated |  |
| `signing_sessions` | [SigningSession](#scalar.multisig.v1beta1.SigningSession) | repeated |  |
| `keys` | [Key](#scalar.multisig.v1beta1.Key) | repeated |  |
| `key_epochs` | [KeyEpoch](#scalar.multisig.v1beta1.KeyEpoch) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/multisig/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/v1beta1/tx.proto



<a name="scalar.multisig.v1beta1.KeygenOptInRequest"></a>

### KeygenOptInRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.KeygenOptInResponse"></a>

### KeygenOptInResponse







<a name="scalar.multisig.v1beta1.KeygenOptOutRequest"></a>

### KeygenOptOutRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.KeygenOptOutResponse"></a>

### KeygenOptOutResponse







<a name="scalar.multisig.v1beta1.RotateKeyRequest"></a>

### RotateKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.RotateKeyResponse"></a>

### RotateKeyResponse







<a name="scalar.multisig.v1beta1.StartKeygenRequest"></a>

### StartKeygenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.multisig.v1beta1.StartKeygenResponse"></a>

### StartKeygenResponse







<a name="scalar.multisig.v1beta1.SubmitPubKeyRequest"></a>

### SubmitPubKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |
| `pub_key` | [bytes](#bytes) |  |  |
| `signature` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.SubmitPubKeyResponse"></a>

### SubmitPubKeyResponse







<a name="scalar.multisig.v1beta1.SubmitSignatureRequest"></a>

### SubmitSignatureRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `sig_id` | [uint64](#uint64) |  |  |
| `signature` | [bytes](#bytes) |  |  |






<a name="scalar.multisig.v1beta1.SubmitSignatureResponse"></a>

### SubmitSignatureResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/multisig/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.multisig.v1beta1.MsgService"></a>

### MsgService
Msg defines the multisig Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `StartKeygen` | [StartKeygenRequest](#scalar.multisig.v1beta1.StartKeygenRequest) | [StartKeygenResponse](#scalar.multisig.v1beta1.StartKeygenResponse) |  | POST|/scalar/multisig/start_keygen|
| `SubmitPubKey` | [SubmitPubKeyRequest](#scalar.multisig.v1beta1.SubmitPubKeyRequest) | [SubmitPubKeyResponse](#scalar.multisig.v1beta1.SubmitPubKeyResponse) |  | POST|/scalar/multisig/submit_pub_key|
| `SubmitSignature` | [SubmitSignatureRequest](#scalar.multisig.v1beta1.SubmitSignatureRequest) | [SubmitSignatureResponse](#scalar.multisig.v1beta1.SubmitSignatureResponse) |  | POST|/scalar/multisig/submit_signature|
| `RotateKey` | [RotateKeyRequest](#scalar.multisig.v1beta1.RotateKeyRequest) | [RotateKeyResponse](#scalar.multisig.v1beta1.RotateKeyResponse) |  | POST|/scalar/multisig/rotate_key|
| `KeygenOptOut` | [KeygenOptOutRequest](#scalar.multisig.v1beta1.KeygenOptOutRequest) | [KeygenOptOutResponse](#scalar.multisig.v1beta1.KeygenOptOutResponse) |  | POST|/scalar/multisig/v1beta1/keygen_opt_out|
| `KeygenOptIn` | [KeygenOptInRequest](#scalar.multisig.v1beta1.KeygenOptInRequest) | [KeygenOptInResponse](#scalar.multisig.v1beta1.KeygenOptInResponse) |  | POST|/scalar/multisig/v1beta1/keygen_opt_in|


<a name="scalar.multisig.v1beta1.QueryService"></a>

### QueryService
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `KeyID` | [KeyIDRequest](#scalar.multisig.v1beta1.KeyIDRequest) | [KeyIDResponse](#scalar.multisig.v1beta1.KeyIDResponse) | KeyID returns the key ID of a key assigned to a given chain. If no key is assigned, it returns the grpc NOT_FOUND error. | GET|/scalar/multisig/v1beta1/key_id/{chain}|
| `NextKeyID` | [NextKeyIDRequest](#scalar.multisig.v1beta1.NextKeyIDRequest) | [NextKeyIDResponse](#scalar.multisig.v1beta1.NextKeyIDResponse) | NextKeyID returns the key ID assigned for the next rotation on a given chain. If no key rotation is in progress, it returns the grpc NOT_FOUND error. | GET|/scalar/multisig/v1beta1/next_key_id/{chain}|
| `Key` | [KeyRequest](#scalar.multisig.v1beta1.KeyRequest) | [KeyResponse](#scalar.multisig.v1beta1.KeyResponse) | Key returns the key corresponding to a given key ID. If no key is found, it returns the grpc NOT_FOUND error. | GET|/scalar/multisig/v1beta1/key|
| `KeygenSession` | [KeygenSessionRequest](#scalar.multisig.v1beta1.KeygenSessionRequest) | [KeygenSessionResponse](#scalar.multisig.v1beta1.KeygenSessionResponse) | KeygenSession returns the keygen session info for a given key ID. If no key is found, it returns the grpc NOT_FOUND error. | GET|/scalar/multisig/v1beta1/keygen_session|
| `Params` | [ParamsRequest](#scalar.multisig.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.multisig.v1beta1.ParamsResponse) |  | GET|/scalar/multisig/v1beta1/params|

 <!-- end services -->



<a name="scalar/nexus/exported/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/nexus/exported/v1beta1/types.proto



<a name="scalar.nexus.exported.v1beta1.Asset"></a>

### Asset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `is_native_asset` | [bool](#bool) |  |  |






<a name="scalar.nexus.exported.v1beta1.Chain"></a>

### Chain
Chain represents the properties of a registered blockchain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | The descriptor of the chain, e.g. "evm|11155111" |
| `supports_foreign_assets` | [bool](#bool) |  |  |
| `key_type` | [scalar.tss.exported.v1beta1.KeyType](#scalar.tss.exported.v1beta1.KeyType) |  |  |
| `module` | [string](#string) |  | the module has two types: chains and scalarnet |






<a name="scalar.nexus.exported.v1beta1.CrossChainAddress"></a>

### CrossChainAddress
CrossChainAddress represents a generalized address on any registered chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [Chain](#scalar.nexus.exported.v1beta1.Chain) |  |  |
| `address` | [string](#string) |  |  |






<a name="scalar.nexus.exported.v1beta1.CrossChainTransfer"></a>

### CrossChainTransfer
CrossChainTransfer represents a generalized transfer of some asset to a
registered blockchain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `recipient` | [CrossChainAddress](#scalar.nexus.exported.v1beta1.CrossChainAddress) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `id` | [uint64](#uint64) |  |  |
| `state` | [TransferState](#scalar.nexus.exported.v1beta1.TransferState) |  |  |
| `source_tx_hash` | [bytes](#bytes) |  |  |






<a name="scalar.nexus.exported.v1beta1.FeeInfo"></a>

### FeeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |
| `fee_rate` | [bytes](#bytes) |  |  |
| `min_fee` | [bytes](#bytes) |  |  |
| `max_fee` | [bytes](#bytes) |  |  |






<a name="scalar.nexus.exported.v1beta1.GeneralMessage"></a>

### GeneralMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `sender` | [CrossChainAddress](#scalar.nexus.exported.v1beta1.CrossChainAddress) |  |  |
| `recipient` | [CrossChainAddress](#scalar.nexus.exported.v1beta1.CrossChainAddress) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `status` | [GeneralMessage.Status](#scalar.nexus.exported.v1beta1.GeneralMessage.Status) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `source_tx_id` | [bytes](#bytes) |  |  |
| `source_tx_index` | [uint64](#uint64) |  |  |
| `payload` | [bytes](#bytes) |  | Additional data for the message, metadata is encoded in the payload, it can be fee information, etc. It will be used later when enqueuing the command and batch command. Currently, the main purpose is use to form the psbt for btc |






<a name="scalar.nexus.exported.v1beta1.TransferFee"></a>

### TransferFee
TransferFee represents accumulated fees generated by the network


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="scalar.nexus.exported.v1beta1.WasmMessage"></a>

### WasmMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `source_chain` | [string](#string) |  |  |
| `source_address` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `destination_address` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `source_tx_id` | [bytes](#bytes) |  |  |
| `source_tx_index` | [uint64](#uint64) |  |  |
| `sender` | [bytes](#bytes) |  |  |
| `id` | [string](#string) |  |  |





 <!-- end messages -->


<a name="scalar.nexus.exported.v1beta1.GeneralMessage.Status"></a>

### GeneralMessage.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_APPROVED | 1 |  |
| STATUS_PROCESSING | 2 |  |
| STATUS_EXECUTED | 3 |  |
| STATUS_FAILED | 4 |  |



<a name="scalar.nexus.exported.v1beta1.TransferDirection"></a>

### TransferDirection


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSFER_DIRECTION_UNSPECIFIED | 0 |  |
| TRANSFER_DIRECTION_FROM | 1 |  |
| TRANSFER_DIRECTION_TO | 2 |  |



<a name="scalar.nexus.exported.v1beta1.TransferState"></a>

### TransferState


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRANSFER_STATE_UNSPECIFIED | 0 |  |
| TRANSFER_STATE_PENDING | 1 |  |
| TRANSFER_STATE_ARCHIVED | 2 |  |
| TRANSFER_STATE_INSUFFICIENT_AMOUNT | 3 |  |
| TRANSFER_STATE_FAILED | 4 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/nexus/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/nexus/v1beta1/events.proto



<a name="scalar.nexus.v1beta1.FeeDeducted"></a>

### FeeDeducted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transfer_id` | [uint64](#uint64) |  |  |
| `recipient_chain` | [string](#string) |  |  |
| `recipient_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.nexus.v1beta1.InsufficientFee"></a>

### InsufficientFee



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transfer_id` | [uint64](#uint64) |  |  |
| `recipient_chain` | [string](#string) |  |  |
| `recipient_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.nexus.v1beta1.MessageExecuted"></a>

### MessageExecuted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `source_chain` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.MessageFailed"></a>

### MessageFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `source_chain` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.MessageProcessing"></a>

### MessageProcessing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `source_chain` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.MessageReceived"></a>

### MessageReceived



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `sender` | [scalar.nexus.exported.v1beta1.CrossChainAddress](#scalar.nexus.exported.v1beta1.CrossChainAddress) |  |  |
| `recipient` | [scalar.nexus.exported.v1beta1.CrossChainAddress](#scalar.nexus.exported.v1beta1.CrossChainAddress) |  |  |






<a name="scalar.nexus.v1beta1.RateLimitUpdated"></a>

### RateLimitUpdated



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `window` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="scalar.nexus.v1beta1.WasmMessageRouted"></a>

### WasmMessageRouted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message` | [scalar.nexus.exported.v1beta1.WasmMessage](#scalar.nexus.exported.v1beta1.WasmMessage) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/nexus/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/nexus/v1beta1/params.proto



<a name="scalar.nexus.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain_activation_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `chain_maintainer_missing_vote_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `chain_maintainer_incorrect_vote_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `chain_maintainer_check_window` | [int32](#int32) |  |  |
| `gateway` | [bytes](#bytes) |  |  |
| `end_blocker_limit` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/utils/v1beta1/bitmap.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/utils/v1beta1/bitmap.proto



<a name="scalar.utils.v1beta1.Bitmap"></a>

### Bitmap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `true_count_cache` | [CircularBuffer](#scalar.utils.v1beta1.CircularBuffer) |  |  |






<a name="scalar.utils.v1beta1.CircularBuffer"></a>

### CircularBuffer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cumulative_value` | [uint64](#uint64) | repeated |  |
| `index` | [int32](#int32) |  |  |
| `max_size` | [int32](#int32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/nexus/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/nexus/v1beta1/types.proto



<a name="scalar.nexus.v1beta1.ChainState"></a>

### ChainState
ChainState represents the state of a registered blockchain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [scalar.nexus.exported.v1beta1.Chain](#scalar.nexus.exported.v1beta1.Chain) |  |  |
| `activated` | [bool](#bool) |  |  |
| `assets` | [scalar.nexus.exported.v1beta1.Asset](#scalar.nexus.exported.v1beta1.Asset) | repeated |  |
| `maintainer_states` | [MaintainerState](#scalar.nexus.v1beta1.MaintainerState) | repeated | **Deprecated.**  |






<a name="scalar.nexus.v1beta1.LinkedAddresses"></a>

### LinkedAddresses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit_address` | [scalar.nexus.exported.v1beta1.CrossChainAddress](#scalar.nexus.exported.v1beta1.CrossChainAddress) |  |  |
| `recipient_address` | [scalar.nexus.exported.v1beta1.CrossChainAddress](#scalar.nexus.exported.v1beta1.CrossChainAddress) |  |  |






<a name="scalar.nexus.v1beta1.MaintainerState"></a>

### MaintainerState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `missing_votes` | [scalar.utils.v1beta1.Bitmap](#scalar.utils.v1beta1.Bitmap) |  |  |
| `incorrect_votes` | [scalar.utils.v1beta1.Bitmap](#scalar.utils.v1beta1.Bitmap) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.RateLimit"></a>

### RateLimit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `window` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="scalar.nexus.v1beta1.TransferEpoch"></a>

### TransferEpoch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `epoch` | [uint64](#uint64) |  |  |
| `direction` | [scalar.nexus.exported.v1beta1.TransferDirection](#scalar.nexus.exported.v1beta1.TransferDirection) |  | indicates whether the rate tracking is for transfers going |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/nexus/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/nexus/v1beta1/genesis.proto



<a name="scalar.nexus.v1beta1.GenesisState"></a>

### GenesisState
GenesisState represents the genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.nexus.v1beta1.Params) |  |  |
| `nonce` | [uint64](#uint64) |  |  |
| `chains` | [scalar.nexus.exported.v1beta1.Chain](#scalar.nexus.exported.v1beta1.Chain) | repeated |  |
| `chain_states` | [ChainState](#scalar.nexus.v1beta1.ChainState) | repeated |  |
| `linked_addresses` | [LinkedAddresses](#scalar.nexus.v1beta1.LinkedAddresses) | repeated |  |
| `transfers` | [scalar.nexus.exported.v1beta1.CrossChainTransfer](#scalar.nexus.exported.v1beta1.CrossChainTransfer) | repeated |  |
| `fee` | [scalar.nexus.exported.v1beta1.TransferFee](#scalar.nexus.exported.v1beta1.TransferFee) |  |  |
| `fee_infos` | [scalar.nexus.exported.v1beta1.FeeInfo](#scalar.nexus.exported.v1beta1.FeeInfo) | repeated |  |
| `rate_limits` | [RateLimit](#scalar.nexus.v1beta1.RateLimit) | repeated |  |
| `transfer_epochs` | [TransferEpoch](#scalar.nexus.v1beta1.TransferEpoch) | repeated |  |
| `messages` | [scalar.nexus.exported.v1beta1.GeneralMessage](#scalar.nexus.exported.v1beta1.GeneralMessage) | repeated |  |
| `message_nonce` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/nexus/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/nexus/v1beta1/query.proto



<a name="scalar.nexus.v1beta1.AssetsRequest"></a>

### AssetsRequest
AssetsRequest represents a message that queries the registered assets of a
chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.AssetsResponse"></a>

### AssetsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `assets` | [string](#string) | repeated |  |






<a name="scalar.nexus.v1beta1.ChainMaintainersRequest"></a>

### ChainMaintainersRequest
ChainMaintainersRequest represents a message that queries
the chain maintainers for the specified chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.ChainMaintainersResponse"></a>

### ChainMaintainersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `maintainers` | [bytes](#bytes) | repeated |  |






<a name="scalar.nexus.v1beta1.ChainStateRequest"></a>

### ChainStateRequest
ChainStateRequest represents a message that queries the state of a chain
registered on the network


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.ChainStateResponse"></a>

### ChainStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `state` | [ChainState](#scalar.nexus.v1beta1.ChainState) |  |  |






<a name="scalar.nexus.v1beta1.ChainsByAssetRequest"></a>

### ChainsByAssetRequest
ChainsByAssetRequest represents a message that queries the chains
that support an asset on the network


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.ChainsByAssetResponse"></a>

### ChainsByAssetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chains` | [string](#string) | repeated |  |






<a name="scalar.nexus.v1beta1.ChainsRequest"></a>

### ChainsRequest
ChainsRequest represents a message that queries the chains
registered on the network


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [ChainStatus](#scalar.nexus.v1beta1.ChainStatus) |  |  |






<a name="scalar.nexus.v1beta1.ChainsResponse"></a>

### ChainsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chains` | [string](#string) | repeated |  |






<a name="scalar.nexus.v1beta1.FeeInfoRequest"></a>

### FeeInfoRequest
FeeInfoRequest represents a message that queries the transfer fees associated
to an asset on a chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.FeeInfoResponse"></a>

### FeeInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fee_info` | [scalar.nexus.exported.v1beta1.FeeInfo](#scalar.nexus.exported.v1beta1.FeeInfo) |  |  |






<a name="scalar.nexus.v1beta1.LatestDepositAddressRequest"></a>

### LatestDepositAddressRequest
LatestDepositAddressRequest represents a message that queries a deposit
address by recipient address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `recipient_addr` | [string](#string) |  |  |
| `recipient_chain` | [string](#string) |  |  |
| `deposit_chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.LatestDepositAddressResponse"></a>

### LatestDepositAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit_addr` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.MessageRequest"></a>

### MessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.MessageResponse"></a>

### MessageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message` | [scalar.nexus.exported.v1beta1.GeneralMessage](#scalar.nexus.exported.v1beta1.GeneralMessage) |  |  |






<a name="scalar.nexus.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.nexus.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.nexus.v1beta1.Params) |  |  |






<a name="scalar.nexus.v1beta1.RecipientAddressRequest"></a>

### RecipientAddressRequest
RecipientAddressRequest represents a message that queries the registered
recipient address for a given deposit address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit_addr` | [string](#string) |  |  |
| `deposit_chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.RecipientAddressResponse"></a>

### RecipientAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `recipient_addr` | [string](#string) |  |  |
| `recipient_chain` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.TransferFeeRequest"></a>

### TransferFeeRequest
TransferFeeRequest represents a message that queries the fees charged by
the network for a cross-chain transfer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `source_chain` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.TransferFeeResponse"></a>

### TransferFeeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.nexus.v1beta1.TransferRateLimit"></a>

### TransferRateLimit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `limit` | [bytes](#bytes) |  |  |
| `window` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `incoming` | [bytes](#bytes) |  | **Deprecated.**  |
| `outgoing` | [bytes](#bytes) |  | **Deprecated.**  |
| `time_left` | [google.protobuf.Duration](#google.protobuf.Duration) |  | time_left indicates the time left in the rate limit window |
| `from` | [bytes](#bytes) |  |  |
| `to` | [bytes](#bytes) |  |  |






<a name="scalar.nexus.v1beta1.TransferRateLimitRequest"></a>

### TransferRateLimitRequest
TransferRateLimitRequest represents a message that queries the registered
transfer rate limit and current transfer amounts for a given chain and asset


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.TransferRateLimitResponse"></a>

### TransferRateLimitResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transfer_rate_limit` | [TransferRateLimit](#scalar.nexus.v1beta1.TransferRateLimit) |  |  |






<a name="scalar.nexus.v1beta1.TransfersForChainRequest"></a>

### TransfersForChainRequest
TransfersForChainRequest represents a message that queries the
transfers for the specified chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `state` | [scalar.nexus.exported.v1beta1.TransferState](#scalar.nexus.exported.v1beta1.TransferState) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="scalar.nexus.v1beta1.TransfersForChainResponse"></a>

### TransfersForChainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transfers` | [scalar.nexus.exported.v1beta1.CrossChainTransfer](#scalar.nexus.exported.v1beta1.CrossChainTransfer) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->


<a name="scalar.nexus.v1beta1.ChainStatus"></a>

### ChainStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHAIN_STATUS_UNSPECIFIED | 0 |  |
| CHAIN_STATUS_ACTIVATED | 1 |  |
| CHAIN_STATUS_DEACTIVATED | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/nexus/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/nexus/v1beta1/tx.proto



<a name="scalar.nexus.v1beta1.ActivateChainRequest"></a>

### ActivateChainRequest
ActivateChainRequest represents a message to activate chains


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chains` | [string](#string) | repeated |  |






<a name="scalar.nexus.v1beta1.ActivateChainResponse"></a>

### ActivateChainResponse







<a name="scalar.nexus.v1beta1.DeactivateChainRequest"></a>

### DeactivateChainRequest
DeactivateChainRequest represents a message to deactivate chains


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chains` | [string](#string) | repeated |  |






<a name="scalar.nexus.v1beta1.DeactivateChainResponse"></a>

### DeactivateChainResponse







<a name="scalar.nexus.v1beta1.DeregisterChainMaintainerRequest"></a>

### DeregisterChainMaintainerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chains` | [string](#string) | repeated |  |






<a name="scalar.nexus.v1beta1.DeregisterChainMaintainerResponse"></a>

### DeregisterChainMaintainerResponse







<a name="scalar.nexus.v1beta1.RegisterAssetFeeRequest"></a>

### RegisterAssetFeeRequest
RegisterAssetFeeRequest represents a message to register the transfer fee
info associated to an asset on a chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `fee_info` | [scalar.nexus.exported.v1beta1.FeeInfo](#scalar.nexus.exported.v1beta1.FeeInfo) |  |  |






<a name="scalar.nexus.v1beta1.RegisterAssetFeeResponse"></a>

### RegisterAssetFeeResponse







<a name="scalar.nexus.v1beta1.RegisterChainMaintainerRequest"></a>

### RegisterChainMaintainerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chains` | [string](#string) | repeated |  |






<a name="scalar.nexus.v1beta1.RegisterChainMaintainerResponse"></a>

### RegisterChainMaintainerResponse







<a name="scalar.nexus.v1beta1.SetTransferRateLimitRequest"></a>

### SetTransferRateLimitRequest
SetTransferRateLimitRequest represents a message to set rate limits on
transfers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `window` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="scalar.nexus.v1beta1.SetTransferRateLimitResponse"></a>

### SetTransferRateLimitResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/nexus/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/nexus/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.nexus.v1beta1.MsgService"></a>

### MsgService
Msg defines the nexus Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `RegisterChainMaintainer` | [RegisterChainMaintainerRequest](#scalar.nexus.v1beta1.RegisterChainMaintainerRequest) | [RegisterChainMaintainerResponse](#scalar.nexus.v1beta1.RegisterChainMaintainerResponse) |  | POST|/scalar/nexus/register_chain_maintainer|
| `DeregisterChainMaintainer` | [DeregisterChainMaintainerRequest](#scalar.nexus.v1beta1.DeregisterChainMaintainerRequest) | [DeregisterChainMaintainerResponse](#scalar.nexus.v1beta1.DeregisterChainMaintainerResponse) |  | POST|/scalar/nexus/deregister_chain_maintainer|
| `ActivateChain` | [ActivateChainRequest](#scalar.nexus.v1beta1.ActivateChainRequest) | [ActivateChainResponse](#scalar.nexus.v1beta1.ActivateChainResponse) |  | POST|/scalar/nexus/activate_chain|
| `DeactivateChain` | [DeactivateChainRequest](#scalar.nexus.v1beta1.DeactivateChainRequest) | [DeactivateChainResponse](#scalar.nexus.v1beta1.DeactivateChainResponse) |  | POST|/scalar/nexus/deactivate_chain|
| `RegisterAssetFee` | [RegisterAssetFeeRequest](#scalar.nexus.v1beta1.RegisterAssetFeeRequest) | [RegisterAssetFeeResponse](#scalar.nexus.v1beta1.RegisterAssetFeeResponse) |  | POST|/scalar/nexus/register_asset_fee|
| `SetTransferRateLimit` | [SetTransferRateLimitRequest](#scalar.nexus.v1beta1.SetTransferRateLimitRequest) | [SetTransferRateLimitResponse](#scalar.nexus.v1beta1.SetTransferRateLimitResponse) |  | POST|/scalar/nexus/set_transfer_rate_limit|


<a name="scalar.nexus.v1beta1.QueryService"></a>

### QueryService
QueryService defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `LatestDepositAddress` | [LatestDepositAddressRequest](#scalar.nexus.v1beta1.LatestDepositAddressRequest) | [LatestDepositAddressResponse](#scalar.nexus.v1beta1.LatestDepositAddressResponse) | LatestDepositAddress queries the a deposit address by recipient | GET|/scalar/nexus/v1beta1/latest_deposit_address/{recipient_addr}/{recipient_chain}/{deposit_chain}|
| `TransfersForChain` | [TransfersForChainRequest](#scalar.nexus.v1beta1.TransfersForChainRequest) | [TransfersForChainResponse](#scalar.nexus.v1beta1.TransfersForChainResponse) | TransfersForChain queries transfers by chain | GET|/scalar/nexus/v1beta1/transfers_for_chain/{chain}/{state}|
| `FeeInfo` | [FeeInfoRequest](#scalar.nexus.v1beta1.FeeInfoRequest) | [FeeInfoResponse](#scalar.nexus.v1beta1.FeeInfoResponse) | FeeInfo queries the fee info by chain and asset | GET|/scalar/nexus/v1beta1/fee_info/{chain}/{asset}GET|/scalar/nexus/v1beta1/fee|
| `TransferFee` | [TransferFeeRequest](#scalar.nexus.v1beta1.TransferFeeRequest) | [TransferFeeResponse](#scalar.nexus.v1beta1.TransferFeeResponse) | TransferFee queries the transfer fee by the source, destination chain, and amount. If amount is 0, the min fee is returned | GET|/scalar/nexus/v1beta1/transfer_fee/{source_chain}/{destination_chain}/{amount}GET|/scalar/nexus/v1beta1/transfer_fee|
| `Chains` | [ChainsRequest](#scalar.nexus.v1beta1.ChainsRequest) | [ChainsResponse](#scalar.nexus.v1beta1.ChainsResponse) | Chains queries the chains registered on the network | GET|/scalar/nexus/v1beta1/chains|
| `Assets` | [AssetsRequest](#scalar.nexus.v1beta1.AssetsRequest) | [AssetsResponse](#scalar.nexus.v1beta1.AssetsResponse) | Assets queries the assets registered for a chain | GET|/scalar/nexus/v1beta1/assets/{chain}|
| `ChainState` | [ChainStateRequest](#scalar.nexus.v1beta1.ChainStateRequest) | [ChainStateResponse](#scalar.nexus.v1beta1.ChainStateResponse) | ChainState queries the state of a registered chain on the network | GET|/scalar/nexus/v1beta1/chain_state/{chain}|
| `ChainsByAsset` | [ChainsByAssetRequest](#scalar.nexus.v1beta1.ChainsByAssetRequest) | [ChainsByAssetResponse](#scalar.nexus.v1beta1.ChainsByAssetResponse) | ChainsByAsset queries the chains that support an asset on the network | GET|/scalar/nexus/v1beta1/chains_by_asset/{asset}|
| `RecipientAddress` | [RecipientAddressRequest](#scalar.nexus.v1beta1.RecipientAddressRequest) | [RecipientAddressResponse](#scalar.nexus.v1beta1.RecipientAddressResponse) | RecipientAddress queries the recipient address for a given deposit address | GET|/scalar/nexus/v1beta1/recipient_address/{deposit_chain}/{deposit_addr}|
| `ChainMaintainers` | [ChainMaintainersRequest](#scalar.nexus.v1beta1.ChainMaintainersRequest) | [ChainMaintainersResponse](#scalar.nexus.v1beta1.ChainMaintainersResponse) | ChainMaintainers queries the chain maintainers for a given chain | GET|/scalar/nexus/v1beta1/chain_maintainers/{chain}|
| `TransferRateLimit` | [TransferRateLimitRequest](#scalar.nexus.v1beta1.TransferRateLimitRequest) | [TransferRateLimitResponse](#scalar.nexus.v1beta1.TransferRateLimitResponse) | TransferRateLimit queries the transfer rate limit for a given chain and asset. If a rate limit is not set, nil is returned. | GET|/scalar/nexus/v1beta1/transfer_rate_limit/{chain}/{asset}|
| `Message` | [MessageRequest](#scalar.nexus.v1beta1.MessageRequest) | [MessageResponse](#scalar.nexus.v1beta1.MessageResponse) |  | GET|/scalar/nexus/v1beta1/message|
| `Params` | [ParamsRequest](#scalar.nexus.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.nexus.v1beta1.ParamsResponse) |  | GET|/scalar/nexus/v1beta1/params|

 <!-- end services -->



<a name="scalar/permission/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/permission/v1beta1/types.proto



<a name="scalar.permission.v1beta1.GovAccount"></a>

### GovAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `role` | [scalar.permission.exported.v1beta1.Role](#scalar.permission.exported.v1beta1.Role) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/permission/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/permission/v1beta1/params.proto



<a name="scalar.permission.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/permission/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/permission/v1beta1/genesis.proto



<a name="scalar.permission.v1beta1.GenesisState"></a>

### GenesisState
GenesisState represents the genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.permission.v1beta1.Params) |  |  |
| `governance_key` | [cosmos.crypto.multisig.LegacyAminoPubKey](#cosmos.crypto.multisig.LegacyAminoPubKey) |  |  |
| `gov_accounts` | [GovAccount](#scalar.permission.v1beta1.GovAccount) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/permission/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/permission/v1beta1/query.proto



<a name="scalar.permission.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.permission.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.permission.v1beta1.Params) |  |  |






<a name="scalar.permission.v1beta1.QueryGovernanceKeyRequest"></a>

### QueryGovernanceKeyRequest
QueryGovernanceKeyRequest is the request type for the
Query/GovernanceKey RPC method






<a name="scalar.permission.v1beta1.QueryGovernanceKeyResponse"></a>

### QueryGovernanceKeyResponse
QueryGovernanceKeyResponse is the response type for the
Query/GovernanceKey RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `governance_key` | [cosmos.crypto.multisig.LegacyAminoPubKey](#cosmos.crypto.multisig.LegacyAminoPubKey) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/permission/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/permission/v1beta1/tx.proto



<a name="scalar.permission.v1beta1.DeregisterControllerRequest"></a>

### DeregisterControllerRequest
DeregisterController represents a message to deregister a controller account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `controller` | [bytes](#bytes) |  |  |






<a name="scalar.permission.v1beta1.DeregisterControllerResponse"></a>

### DeregisterControllerResponse







<a name="scalar.permission.v1beta1.RegisterControllerRequest"></a>

### RegisterControllerRequest
MsgRegisterController represents a message to register a controller account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `controller` | [bytes](#bytes) |  |  |






<a name="scalar.permission.v1beta1.RegisterControllerResponse"></a>

### RegisterControllerResponse







<a name="scalar.permission.v1beta1.UpdateGovernanceKeyRequest"></a>

### UpdateGovernanceKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `governance_key` | [cosmos.crypto.multisig.LegacyAminoPubKey](#cosmos.crypto.multisig.LegacyAminoPubKey) |  |  |






<a name="scalar.permission.v1beta1.UpdateGovernanceKeyResponse"></a>

### UpdateGovernanceKeyResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/permission/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/permission/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.permission.v1beta1.Msg"></a>

### Msg
Msg defines the gov Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `RegisterController` | [RegisterControllerRequest](#scalar.permission.v1beta1.RegisterControllerRequest) | [RegisterControllerResponse](#scalar.permission.v1beta1.RegisterControllerResponse) |  | POST|/scalar/permission/register_controller|
| `DeregisterController` | [DeregisterControllerRequest](#scalar.permission.v1beta1.DeregisterControllerRequest) | [DeregisterControllerResponse](#scalar.permission.v1beta1.DeregisterControllerResponse) |  | POST|/scalar/permission/deregister_controller|
| `UpdateGovernanceKey` | [UpdateGovernanceKeyRequest](#scalar.permission.v1beta1.UpdateGovernanceKeyRequest) | [UpdateGovernanceKeyResponse](#scalar.permission.v1beta1.UpdateGovernanceKeyResponse) |  | POST|/scalar/permission/update_governance_key|


<a name="scalar.permission.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GovernanceKey` | [QueryGovernanceKeyRequest](#scalar.permission.v1beta1.QueryGovernanceKeyRequest) | [QueryGovernanceKeyResponse](#scalar.permission.v1beta1.QueryGovernanceKeyResponse) | GovernanceKey returns the multisig governance key | GET|/scalar/permission/v1beta1/governance_key|
| `Params` | [ParamsRequest](#scalar.permission.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.permission.v1beta1.ParamsResponse) |  | GET|/scalar/permission/v1beta1/params|

 <!-- end services -->



<a name="scalar/protocol/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/types.proto



<a name="scalar.protocol.v1beta1.Protocol"></a>

### Protocol



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bitcoin_pubkey` | [bytes](#bytes) |  | BTC's pubkey |
| `scalar_pubkey` | [bytes](#bytes) |  | Scalar's pubkey |
| `scalar_address` | [bytes](#bytes) |  | Scalar's address |
| `name` | [string](#string) |  |  |
| `tag` | [bytes](#bytes) |  |  |
| `attributes` | [scalar.protocol.exported.v1beta1.ProtocolAttributes](#scalar.protocol.exported.v1beta1.ProtocolAttributes) |  |  |
| `status` | [scalar.protocol.exported.v1beta1.Status](#scalar.protocol.exported.v1beta1.Status) |  |  |
| `custodian_group_uid` | [string](#string) |  | scalar.covenant.v1beta1.CustodianGroup custodian_group = 8; |
| `asset` | [scalar.chains.v1beta1.Asset](#scalar.chains.v1beta1.Asset) |  | External asset |
| `chains` | [scalar.protocol.exported.v1beta1.SupportedChain](#scalar.protocol.exported.v1beta1.SupportedChain) | repeated | Other chains with internal asset |
| `avatar` | [bytes](#bytes) |  | Avatar of the protocol, base64 encoded |






<a name="scalar.protocol.v1beta1.ProtocolDetails"></a>

### ProtocolDetails



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bitcoin_pubkey` | [bytes](#bytes) |  | BTC's pubkey |
| `scalar_pubkey` | [bytes](#bytes) |  | Scalar's pubkey |
| `scalar_address` | [bytes](#bytes) |  | Scalar's address |
| `name` | [string](#string) |  |  |
| `tag` | [bytes](#bytes) |  |  |
| `attributes` | [scalar.protocol.exported.v1beta1.ProtocolAttributes](#scalar.protocol.exported.v1beta1.ProtocolAttributes) |  |  |
| `status` | [scalar.protocol.exported.v1beta1.Status](#scalar.protocol.exported.v1beta1.Status) |  |  |
| `custodian_group_uid` | [string](#string) |  |  |
| `asset` | [scalar.chains.v1beta1.Asset](#scalar.chains.v1beta1.Asset) |  | External asset |
| `chains` | [scalar.protocol.exported.v1beta1.SupportedChain](#scalar.protocol.exported.v1beta1.SupportedChain) | repeated | Other chains with internal asset |
| `avatar` | [bytes](#bytes) |  | Avatar of the protocol, base64 encoded |
| `custodian_group` | [scalar.covenant.v1beta1.CustodianGroup](#scalar.covenant.v1beta1.CustodianGroup) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/protocol/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/genesis.proto



<a name="scalar.protocol.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocols` | [Protocol](#scalar.protocol.v1beta1.Protocol) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/protocol/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/params.proto



<a name="scalar.protocol.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/protocol/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/query.proto



<a name="scalar.protocol.v1beta1.ProtocolRequest"></a>

### ProtocolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `origin_chain` | [string](#string) |  |  |
| `minor_chain` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="scalar.protocol.v1beta1.ProtocolResponse"></a>

### ProtocolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocol` | [ProtocolDetails](#scalar.protocol.v1beta1.ProtocolDetails) |  |  |






<a name="scalar.protocol.v1beta1.ProtocolsRequest"></a>

### ProtocolsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pubkey` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `status` | [scalar.protocol.exported.v1beta1.Status](#scalar.protocol.exported.v1beta1.Status) |  |  |






<a name="scalar.protocol.v1beta1.ProtocolsResponse"></a>

### ProtocolsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocols` | [ProtocolDetails](#scalar.protocol.v1beta1.ProtocolDetails) | repeated |  |
| `total` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/protocol/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/tx.proto



<a name="scalar.protocol.v1beta1.AddSupportedChainRequest"></a>

### AddSupportedChainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [scalar.protocol.exported.v1beta1.SupportedChain](#scalar.protocol.exported.v1beta1.SupportedChain) |  |  |






<a name="scalar.protocol.v1beta1.AddSupportedChainResponse"></a>

### AddSupportedChainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocol` | [Protocol](#scalar.protocol.v1beta1.Protocol) |  |  |






<a name="scalar.protocol.v1beta1.CreateProtocolRequest"></a>

### CreateProtocolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  | address |
| `bitcoin_pubkey` | [bytes](#bytes) |  | BTC's pubkey |
| `scalar_pubkey` | [bytes](#bytes) |  | Scalar's pubkey |
| `name` | [string](#string) |  | e.g., "protocol-1" |
| `tag` | [string](#string) |  | e.g., "pools" |
| `attributes` | [scalar.protocol.exported.v1beta1.ProtocolAttributes](#scalar.protocol.exported.v1beta1.ProtocolAttributes) |  |  |
| `custodian_group_uid` | [string](#string) |  |  |
| `asset` | [scalar.chains.v1beta1.Asset](#scalar.chains.v1beta1.Asset) |  | External asset |
| `avatar` | [bytes](#bytes) |  | Avatar of the protocol, base64 encoded |






<a name="scalar.protocol.v1beta1.CreateProtocolResponse"></a>

### CreateProtocolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocol` | [Protocol](#scalar.protocol.v1beta1.Protocol) |  |  |






<a name="scalar.protocol.v1beta1.UpdateProtocolRequest"></a>

### UpdateProtocolRequest
pubkey used as protocol unique id


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `name` | [string](#string) |  |  |
| `tag` | [string](#string) |  | e.g., "pools" |






<a name="scalar.protocol.v1beta1.UpdateProtocolResponse"></a>

### UpdateProtocolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocol` | [Protocol](#scalar.protocol.v1beta1.Protocol) |  |  |






<a name="scalar.protocol.v1beta1.UpdateSupportedChainRequest"></a>

### UpdateSupportedChainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain_family` | [string](#string) |  |  |
| `chain_id` | [uint64](#uint64) |  |  |
| `status` | [scalar.protocol.exported.v1beta1.Status](#scalar.protocol.exported.v1beta1.Status) |  |  |






<a name="scalar.protocol.v1beta1.UpdateSupportedChainResponse"></a>

### UpdateSupportedChainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocol` | [Protocol](#scalar.protocol.v1beta1.Protocol) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/protocol/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.protocol.v1beta1.Msg"></a>

### Msg


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateProtocol` | [CreateProtocolRequest](#scalar.protocol.v1beta1.CreateProtocolRequest) | [CreateProtocolResponse](#scalar.protocol.v1beta1.CreateProtocolResponse) | Create protocol | POST|/scalar/protocol/v1beta1/create_protocol|
| `UpdateProtocol` | [UpdateProtocolRequest](#scalar.protocol.v1beta1.UpdateProtocolRequest) | [UpdateProtocolResponse](#scalar.protocol.v1beta1.UpdateProtocolResponse) |  | POST|/scalar/protocol/v1beta1/update_protocol|
| `AddSupportedChain` | [AddSupportedChainRequest](#scalar.protocol.v1beta1.AddSupportedChainRequest) | [AddSupportedChainResponse](#scalar.protocol.v1beta1.AddSupportedChainResponse) | Add DestinationChain into protocol | POST|/scalar/protocol/v1beta1/add_supported_chain|
| `UpdateSupportedChain` | [UpdateSupportedChainRequest](#scalar.protocol.v1beta1.UpdateSupportedChainRequest) | [UpdateSupportedChainResponse](#scalar.protocol.v1beta1.UpdateSupportedChainResponse) | Delete DestinationChain from protocol | POST|/scalar/protocol/v1beta1/update_supported_chain|


<a name="scalar.protocol.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Protocols` | [ProtocolsRequest](#scalar.protocol.v1beta1.ProtocolsRequest) | [ProtocolsResponse](#scalar.protocol.v1beta1.ProtocolsResponse) | GetProtocols returns all Protocol | GET|/scalar/protocol/v1beta1|
| `Protocol` | [ProtocolRequest](#scalar.protocol.v1beta1.ProtocolRequest) | [ProtocolResponse](#scalar.protocol.v1beta1.ProtocolResponse) |  | GET|/scalar/protocol/v1beta1/protocol|

 <!-- end services -->



<a name="scalar/reward/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/reward/v1beta1/params.proto



<a name="scalar.reward.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `external_chain_voting_inflation_rate` | [bytes](#bytes) |  |  |
| `key_mgmt_relative_inflation_rate` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/reward/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/reward/v1beta1/types.proto



<a name="scalar.reward.v1beta1.Pool"></a>

### Pool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `rewards` | [Pool.Reward](#scalar.reward.v1beta1.Pool.Reward) | repeated |  |






<a name="scalar.reward.v1beta1.Pool.Reward"></a>

### Pool.Reward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [bytes](#bytes) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="scalar.reward.v1beta1.Refund"></a>

### Refund



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payer` | [bytes](#bytes) |  |  |
| `fees` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/reward/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/reward/v1beta1/genesis.proto



<a name="scalar.reward.v1beta1.GenesisState"></a>

### GenesisState
GenesisState represents the genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.reward.v1beta1.Params) |  |  |
| `pools` | [Pool](#scalar.reward.v1beta1.Pool) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/reward/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/reward/v1beta1/query.proto



<a name="scalar.reward.v1beta1.InflationRateRequest"></a>

### InflationRateRequest
InflationRateRequest represents a message that queries the scalar specific
inflation RPC method. Ideally, this would use ValAddress as the validator
field type. However, this makes it awkward for REST-based calls, because it
would expect a byte array as part of the url. So, the bech32 encoded address
string is used for this request instead.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [string](#string) |  |  |






<a name="scalar.reward.v1beta1.InflationRateResponse"></a>

### InflationRateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation_rate` | [bytes](#bytes) |  |  |






<a name="scalar.reward.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.reward.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.reward.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/reward/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/reward/v1beta1/tx.proto



<a name="scalar.reward.v1beta1.RefundMsgRequest"></a>

### RefundMsgRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `inner_message` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="scalar.reward.v1beta1.RefundMsgResponse"></a>

### RefundMsgResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  |  |
| `log` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/reward/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/reward/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.reward.v1beta1.MsgService"></a>

### MsgService
Msg defines the scalarnet Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `RefundMsg` | [RefundMsgRequest](#scalar.reward.v1beta1.RefundMsgRequest) | [RefundMsgResponse](#scalar.reward.v1beta1.RefundMsgResponse) |  | POST|/scalar/reward/refund_message|


<a name="scalar.reward.v1beta1.QueryService"></a>

### QueryService
QueryService defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `InflationRate` | [InflationRateRequest](#scalar.reward.v1beta1.InflationRateRequest) | [InflationRateResponse](#scalar.reward.v1beta1.InflationRateResponse) |  | GET|/scalar/reward/v1beta1/inflation_rate/{validator}GET|/scalar/reward/v1beta1/inflation_rate|
| `Params` | [ParamsRequest](#scalar.reward.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.reward.v1beta1.ParamsResponse) |  | GET|/scalar/reward/v1beta1/params|

 <!-- end services -->



<a name="scalar/scalarnet/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/scalarnet/v1beta1/events.proto



<a name="scalar.scalarnet.v1beta1.ContractCallSubmitted"></a>

### ContractCallSubmitted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message_id` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |
| `source_chain` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload` | [bytes](#bytes) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |






<a name="scalar.scalarnet.v1beta1.ContractCallWithTokenSubmitted"></a>

### ContractCallWithTokenSubmitted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message_id` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |
| `source_chain` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload` | [bytes](#bytes) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.scalarnet.v1beta1.FeeCollected"></a>

### FeeCollected



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `collector` | [bytes](#bytes) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.scalarnet.v1beta1.FeePaid"></a>

### FeePaid



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message_id` | [string](#string) |  |  |
| `recipient` | [bytes](#bytes) |  |  |
| `fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `refund_recipient` | [string](#string) |  |  |
| `asset` | [string](#string) |  | registered asset name in nexus |
| `source_chain` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.IBCTransferCompleted"></a>

### IBCTransferCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `sequence` | [uint64](#uint64) |  |  |
| `port_id` | [string](#string) |  |  |
| `channel_id` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.IBCTransferFailed"></a>

### IBCTransferFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `sequence` | [uint64](#uint64) |  |  |
| `port_id` | [string](#string) |  |  |
| `channel_id` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.IBCTransferRetried"></a>

### IBCTransferRetried



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `receipient` | [string](#string) |  | **Deprecated.**  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `sequence` | [uint64](#uint64) |  |  |
| `port_id` | [string](#string) |  |  |
| `channel_id` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.IBCTransferSent"></a>

### IBCTransferSent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `receipient` | [string](#string) |  | **Deprecated.**  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `sequence` | [uint64](#uint64) |  |  |
| `port_id` | [string](#string) |  |  |
| `channel_id` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.ScalarTransferCompleted"></a>

### ScalarTransferCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `receipient` | [string](#string) |  | **Deprecated.**  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `recipient` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.TokenSent"></a>

### TokenSent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transfer_id` | [uint64](#uint64) |  |  |
| `sender` | [string](#string) |  |  |
| `source_chain` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `destination_address` | [string](#string) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/scalarnet/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/scalarnet/v1beta1/params.proto



<a name="scalar.scalarnet.v1beta1.CallContractProposalMinDeposit"></a>

### CallContractProposalMinDeposit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `min_deposits` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="scalar.scalarnet.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `route_timeout_window` | [uint64](#uint64) |  | IBC packet route timeout window |
| `transfer_limit` | [uint64](#uint64) |  |  |
| `end_blocker_limit` | [uint64](#uint64) |  |  |
| `version` | [uint32](#uint32) |  |  |
| `tag` | [bytes](#bytes) |  |  |
| `call_contracts_proposal_min_deposits` | [CallContractProposalMinDeposit](#scalar.scalarnet.v1beta1.CallContractProposalMinDeposit) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/scalarnet/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/scalarnet/v1beta1/types.proto



<a name="scalar.scalarnet.v1beta1.Asset"></a>

### Asset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `min_amount` | [bytes](#bytes) |  |  |






<a name="scalar.scalarnet.v1beta1.CosmosChain"></a>

### CosmosChain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `ibc_path` | [string](#string) |  |  |
| `assets` | [Asset](#scalar.scalarnet.v1beta1.Asset) | repeated | **Deprecated.**  |
| `addr_prefix` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.Fee"></a>

### Fee



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `recipient` | [bytes](#bytes) |  |  |
| `refund_recipient` | [bytes](#bytes) |  |  |






<a name="scalar.scalarnet.v1beta1.IBCTransfer"></a>

### IBCTransfer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `receiver` | [string](#string) |  |  |
| `token` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `port_id` | [string](#string) |  |  |
| `channel_id` | [string](#string) |  |  |
| `sequence` | [uint64](#uint64) |  | **Deprecated.**  |
| `id` | [uint64](#uint64) |  |  |
| `status` | [IBCTransfer.Status](#scalar.scalarnet.v1beta1.IBCTransfer.Status) |  |  |





 <!-- end messages -->


<a name="scalar.scalarnet.v1beta1.IBCTransfer.Status"></a>

### IBCTransfer.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_PENDING | 1 |  |
| STATUS_COMPLETED | 2 |  |
| STATUS_FAILED | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/scalarnet/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/scalarnet/v1beta1/genesis.proto



<a name="scalar.scalarnet.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.scalarnet.v1beta1.Params) |  |  |
| `collector_address` | [bytes](#bytes) |  |  |
| `chains` | [CosmosChain](#scalar.scalarnet.v1beta1.CosmosChain) | repeated |  |
| `transfer_queue` | [scalar.utils.v1beta1.QueueState](#scalar.utils.v1beta1.QueueState) |  |  |
| `ibc_transfers` | [IBCTransfer](#scalar.scalarnet.v1beta1.IBCTransfer) | repeated |  |
| `seq_id_mapping` | [GenesisState.SeqIdMappingEntry](#scalar.scalarnet.v1beta1.GenesisState.SeqIdMappingEntry) | repeated |  |






<a name="scalar.scalarnet.v1beta1.GenesisState.SeqIdMappingEntry"></a>

### GenesisState.SeqIdMappingEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/scalarnet/v1beta1/proposal.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/scalarnet/v1beta1/proposal.proto



<a name="scalar.scalarnet.v1beta1.CallContractsProposal"></a>

### CallContractsProposal
CallContractsProposal is a gov Content type for calling contracts on other
chains


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `contract_calls` | [ContractCall](#scalar.scalarnet.v1beta1.ContractCall) | repeated |  |






<a name="scalar.scalarnet.v1beta1.ContractCall"></a>

### ContractCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/scalarnet/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/scalarnet/v1beta1/query.proto



<a name="scalar.scalarnet.v1beta1.ChainByIBCPathRequest"></a>

### ChainByIBCPathRequest
ChainByIBCPathRequest represents a message that queries the chain that an IBC
path is registered to


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ibc_path` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.ChainByIBCPathResponse"></a>

### ChainByIBCPathResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.IBCPathRequest"></a>

### IBCPathRequest
IBCPathRequest represents a message that queries the IBC path registered for
a given chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.IBCPathResponse"></a>

### IBCPathResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ibc_path` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.scalarnet.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.scalarnet.v1beta1.Params) |  |  |






<a name="scalar.scalarnet.v1beta1.PendingIBCTransferCountRequest"></a>

### PendingIBCTransferCountRequest







<a name="scalar.scalarnet.v1beta1.PendingIBCTransferCountResponse"></a>

### PendingIBCTransferCountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transfers_by_chain` | [PendingIBCTransferCountResponse.TransfersByChainEntry](#scalar.scalarnet.v1beta1.PendingIBCTransferCountResponse.TransfersByChainEntry) | repeated |  |






<a name="scalar.scalarnet.v1beta1.PendingIBCTransferCountResponse.TransfersByChainEntry"></a>

### PendingIBCTransferCountResponse.TransfersByChainEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/scalarnet/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/scalarnet/v1beta1/tx.proto



<a name="scalar.scalarnet.v1beta1.AddCosmosBasedChainRequest"></a>

### AddCosmosBasedChainRequest
MsgAddCosmosBasedChain represents a message to register a cosmos based chain
to nexus


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [scalar.nexus.exported.v1beta1.Chain](#scalar.nexus.exported.v1beta1.Chain) |  | **Deprecated.** chain was deprecated in v0.27 |
| `addr_prefix` | [string](#string) |  |  |
| `native_assets` | [scalar.nexus.exported.v1beta1.Asset](#scalar.nexus.exported.v1beta1.Asset) | repeated | **Deprecated.** native_assets was deprecated in v0.27 |
| `cosmos_chain` | [string](#string) |  | TODO: Rename this to `chain` after v1beta1 -> v1 version bump |
| `ibc_path` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.AddCosmosBasedChainResponse"></a>

### AddCosmosBasedChainResponse







<a name="scalar.scalarnet.v1beta1.CallContractRequest"></a>

### CallContractRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload` | [bytes](#bytes) |  |  |
| `fee` | [Fee](#scalar.scalarnet.v1beta1.Fee) |  |  |






<a name="scalar.scalarnet.v1beta1.CallContractResponse"></a>

### CallContractResponse







<a name="scalar.scalarnet.v1beta1.ConfirmDepositRequest"></a>

### ConfirmDepositRequest
MsgConfirmDeposit represents a deposit confirmation message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `deposit_address` | [bytes](#bytes) |  |  |
| `denom` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.ConfirmDepositResponse"></a>

### ConfirmDepositResponse







<a name="scalar.scalarnet.v1beta1.ExecutePendingTransfersRequest"></a>

### ExecutePendingTransfersRequest
MsgExecutePendingTransfers represents a message to trigger transfer all
pending transfers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |






<a name="scalar.scalarnet.v1beta1.ExecutePendingTransfersResponse"></a>

### ExecutePendingTransfersResponse







<a name="scalar.scalarnet.v1beta1.LinkRequest"></a>

### LinkRequest
MsgLink represents a message to link a cross-chain address to an Scalar
address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `recipient_addr` | [string](#string) |  |  |
| `recipient_chain` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.LinkResponse"></a>

### LinkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit_addr` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.RegisterAssetRequest"></a>

### RegisterAssetRequest
RegisterAssetRequest represents a message to register an asset to a cosmos
based chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `asset` | [scalar.nexus.exported.v1beta1.Asset](#scalar.nexus.exported.v1beta1.Asset) |  |  |
| `limit` | [bytes](#bytes) |  |  |
| `window` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |






<a name="scalar.scalarnet.v1beta1.RegisterAssetResponse"></a>

### RegisterAssetResponse







<a name="scalar.scalarnet.v1beta1.RegisterFeeCollectorRequest"></a>

### RegisterFeeCollectorRequest
RegisterFeeCollectorRequest represents a message to register scalarnet fee
collector account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `fee_collector` | [bytes](#bytes) |  |  |






<a name="scalar.scalarnet.v1beta1.RegisterFeeCollectorResponse"></a>

### RegisterFeeCollectorResponse







<a name="scalar.scalarnet.v1beta1.RegisterIBCPathRequest"></a>

### RegisterIBCPathRequest
MSgRegisterIBCPath represents a message to register an IBC tracing path for
a cosmos chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `path` | [string](#string) |  |  |






<a name="scalar.scalarnet.v1beta1.RegisterIBCPathResponse"></a>

### RegisterIBCPathResponse







<a name="scalar.scalarnet.v1beta1.RetryIBCTransferRequest"></a>

### RetryIBCTransferRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  | **Deprecated.**  |
| `id` | [uint64](#uint64) |  |  |






<a name="scalar.scalarnet.v1beta1.RetryIBCTransferResponse"></a>

### RetryIBCTransferResponse







<a name="scalar.scalarnet.v1beta1.RouteIBCTransfersRequest"></a>

### RouteIBCTransfersRequest
RouteIBCTransfersRequest represents a message to route pending transfers to
cosmos based chains


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |






<a name="scalar.scalarnet.v1beta1.RouteIBCTransfersResponse"></a>

### RouteIBCTransfersResponse







<a name="scalar.scalarnet.v1beta1.RouteMessageRequest"></a>

### RouteMessageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `id` | [string](#string) |  |  |
| `payload` | [bytes](#bytes) |  |  |
| `feegranter` | [bytes](#bytes) |  |  |






<a name="scalar.scalarnet.v1beta1.RouteMessageResponse"></a>

### RouteMessageResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/scalarnet/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/scalarnet/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.scalarnet.v1beta1.MsgService"></a>

### MsgService
Msg defines the scalarnet Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Link` | [LinkRequest](#scalar.scalarnet.v1beta1.LinkRequest) | [LinkResponse](#scalar.scalarnet.v1beta1.LinkResponse) |  | POST|/scalar/scalarnet/link|
| `ConfirmDeposit` | [ConfirmDepositRequest](#scalar.scalarnet.v1beta1.ConfirmDepositRequest) | [ConfirmDepositResponse](#scalar.scalarnet.v1beta1.ConfirmDepositResponse) |  | POST|/scalar/scalarnet/confirm_deposit|
| `ExecutePendingTransfers` | [ExecutePendingTransfersRequest](#scalar.scalarnet.v1beta1.ExecutePendingTransfersRequest) | [ExecutePendingTransfersResponse](#scalar.scalarnet.v1beta1.ExecutePendingTransfersResponse) |  | POST|/scalar/scalarnet/execute_pending_transfers|
| `AddCosmosBasedChain` | [AddCosmosBasedChainRequest](#scalar.scalarnet.v1beta1.AddCosmosBasedChainRequest) | [AddCosmosBasedChainResponse](#scalar.scalarnet.v1beta1.AddCosmosBasedChainResponse) |  | POST|/scalar/scalarnet/add_cosmos_based_chain|
| `RegisterAsset` | [RegisterAssetRequest](#scalar.scalarnet.v1beta1.RegisterAssetRequest) | [RegisterAssetResponse](#scalar.scalarnet.v1beta1.RegisterAssetResponse) |  | POST|/scalar/scalarnet/register_asset|
| `RouteIBCTransfers` | [RouteIBCTransfersRequest](#scalar.scalarnet.v1beta1.RouteIBCTransfersRequest) | [RouteIBCTransfersResponse](#scalar.scalarnet.v1beta1.RouteIBCTransfersResponse) |  | POST|/scalar/scalarnet/route_ibc_transfers|
| `RegisterFeeCollector` | [RegisterFeeCollectorRequest](#scalar.scalarnet.v1beta1.RegisterFeeCollectorRequest) | [RegisterFeeCollectorResponse](#scalar.scalarnet.v1beta1.RegisterFeeCollectorResponse) |  | POST|/scalar/scalarnet/register_fee_collector|
| `RetryIBCTransfer` | [RetryIBCTransferRequest](#scalar.scalarnet.v1beta1.RetryIBCTransferRequest) | [RetryIBCTransferResponse](#scalar.scalarnet.v1beta1.RetryIBCTransferResponse) |  | POST|/scalar/scalarnet/retry_ibc_transfer|
| `RouteMessage` | [RouteMessageRequest](#scalar.scalarnet.v1beta1.RouteMessageRequest) | [RouteMessageResponse](#scalar.scalarnet.v1beta1.RouteMessageResponse) |  | POST|/scalar/scalarnet/route_message|
| `CallContract` | [CallContractRequest](#scalar.scalarnet.v1beta1.CallContractRequest) | [CallContractResponse](#scalar.scalarnet.v1beta1.CallContractResponse) |  | POST|/scalar/scalarnet/call_contract|


<a name="scalar.scalarnet.v1beta1.QueryService"></a>

### QueryService
QueryService defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PendingIBCTransferCount` | [PendingIBCTransferCountRequest](#scalar.scalarnet.v1beta1.PendingIBCTransferCountRequest) | [PendingIBCTransferCountResponse](#scalar.scalarnet.v1beta1.PendingIBCTransferCountResponse) | PendingIBCTransferCount queries the pending ibc transfers for all chains | GET|/scalar/scalarnet/v1beta1/ibc_transfer_count|
| `Params` | [ParamsRequest](#scalar.scalarnet.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.scalarnet.v1beta1.ParamsResponse) |  | GET|/scalar/scalarnet/v1beta1/params|
| `IBCPath` | [IBCPathRequest](#scalar.scalarnet.v1beta1.IBCPathRequest) | [IBCPathResponse](#scalar.scalarnet.v1beta1.IBCPathResponse) |  | GET|/scalar/scalarnet/v1beta1/ibc_path/{chain}|
| `ChainByIBCPath` | [ChainByIBCPathRequest](#scalar.scalarnet.v1beta1.ChainByIBCPathRequest) | [ChainByIBCPathResponse](#scalar.scalarnet.v1beta1.ChainByIBCPathResponse) |  | GET|/scalar/scalarnet/v1beta1/chain_by_ibc_path/{ibc_path}|

 <!-- end services -->



<a name="scalar/snapshot/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/snapshot/v1beta1/params.proto



<a name="scalar.snapshot.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_proxy_balance` | [int64](#int64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/snapshot/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/snapshot/v1beta1/types.proto



<a name="scalar.snapshot.v1beta1.ProxiedValidator"></a>

### ProxiedValidator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [bytes](#bytes) |  |  |
| `proxy` | [bytes](#bytes) |  |  |
| `active` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/snapshot/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/snapshot/v1beta1/genesis.proto



<a name="scalar.snapshot.v1beta1.GenesisState"></a>

### GenesisState
GenesisState represents the genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.snapshot.v1beta1.Params) |  |  |
| `proxied_validators` | [ProxiedValidator](#scalar.snapshot.v1beta1.ProxiedValidator) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/snapshot/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/snapshot/v1beta1/query.proto



<a name="scalar.snapshot.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.snapshot.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.snapshot.v1beta1.Params) |  |  |






<a name="scalar.snapshot.v1beta1.QueryValidatorsResponse"></a>

### QueryValidatorsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validators` | [QueryValidatorsResponse.Validator](#scalar.snapshot.v1beta1.QueryValidatorsResponse.Validator) | repeated |  |






<a name="scalar.snapshot.v1beta1.QueryValidatorsResponse.TssIllegibilityInfo"></a>

### QueryValidatorsResponse.TssIllegibilityInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tombstoned` | [bool](#bool) |  |  |
| `jailed` | [bool](#bool) |  |  |
| `missed_too_many_blocks` | [bool](#bool) |  |  |
| `no_proxy_registered` | [bool](#bool) |  |  |
| `tss_suspended` | [bool](#bool) |  |  |
| `proxy_insuficient_funds` | [bool](#bool) |  |  |
| `stale_tss_heartbeat` | [bool](#bool) |  |  |






<a name="scalar.snapshot.v1beta1.QueryValidatorsResponse.Validator"></a>

### QueryValidatorsResponse.Validator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `operator_address` | [string](#string) |  |  |
| `moniker` | [string](#string) |  |  |
| `tss_illegibility_info` | [QueryValidatorsResponse.TssIllegibilityInfo](#scalar.snapshot.v1beta1.QueryValidatorsResponse.TssIllegibilityInfo) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/snapshot/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/snapshot/v1beta1/tx.proto



<a name="scalar.snapshot.v1beta1.DeactivateProxyRequest"></a>

### DeactivateProxyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |






<a name="scalar.snapshot.v1beta1.DeactivateProxyResponse"></a>

### DeactivateProxyResponse







<a name="scalar.snapshot.v1beta1.RegisterProxyRequest"></a>

### RegisterProxyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `proxy_addr` | [bytes](#bytes) |  |  |






<a name="scalar.snapshot.v1beta1.RegisterProxyResponse"></a>

### RegisterProxyResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/snapshot/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/snapshot/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.snapshot.v1beta1.MsgService"></a>

### MsgService
Msg defines the snapshot Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `RegisterProxy` | [RegisterProxyRequest](#scalar.snapshot.v1beta1.RegisterProxyRequest) | [RegisterProxyResponse](#scalar.snapshot.v1beta1.RegisterProxyResponse) | RegisterProxy defines a method for registering a proxy account that can act in a validator account's stead. | POST|/scalar/snapshot/register_proxy|
| `DeactivateProxy` | [DeactivateProxyRequest](#scalar.snapshot.v1beta1.DeactivateProxyRequest) | [DeactivateProxyResponse](#scalar.snapshot.v1beta1.DeactivateProxyResponse) | DeactivateProxy defines a method for deregistering a proxy account. | POST|/scalar/snapshot/deactivate_proxy|


<a name="scalar.snapshot.v1beta1.QueryService"></a>

### QueryService
QueryService defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [ParamsRequest](#scalar.snapshot.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.snapshot.v1beta1.ParamsResponse) |  | GET|/scalar/snapshot/v1beta1/params|

 <!-- end services -->



<a name="scalar/tss/tofnd/v1beta1/common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/tofnd/v1beta1/common.proto
File copied from golang tofnd with minor tweaks


<a name="tofnd.KeyPresenceRequest"></a>

### KeyPresenceRequest
Key presence check types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_uid` | [string](#string) |  |  |
| `pub_key` | [bytes](#bytes) |  | SEC1-encoded compressed pub key bytes to find the right |






<a name="tofnd.KeyPresenceResponse"></a>

### KeyPresenceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `response` | [KeyPresenceResponse.Response](#tofnd.KeyPresenceResponse.Response) |  |  |





 <!-- end messages -->


<a name="tofnd.KeyPresenceResponse.Response"></a>

### KeyPresenceResponse.Response


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESPONSE_UNSPECIFIED | 0 |  |
| RESPONSE_PRESENT | 1 |  |
| RESPONSE_ABSENT | 2 |  |
| RESPONSE_FAIL | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/tss/tofnd/v1beta1/multisig.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/tofnd/v1beta1/multisig.proto
File copied from golang tofnd with minor tweaks


<a name="tofnd.KeygenRequest"></a>

### KeygenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_uid` | [string](#string) |  |  |
| `party_uid` | [string](#string) |  | used only for logging |






<a name="tofnd.KeygenResponse"></a>

### KeygenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_key` | [bytes](#bytes) |  | SEC1-encoded compressed curve point |
| `error` | [string](#string) |  | reply with an error message if keygen fails |






<a name="tofnd.SignRequest"></a>

### SignRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_uid` | [string](#string) |  |  |
| `msg_to_sign` | [bytes](#bytes) |  | 32-byte pre-hashed message digest |
| `party_uid` | [string](#string) |  | used only for logging |
| `pub_key` | [bytes](#bytes) |  | SEC1-encoded compressed pub key bytes to find the right |






<a name="tofnd.SignResponse"></a>

### SignResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signature` | [bytes](#bytes) |  | ASN.1 DER-encoded ECDSA signature |
| `error` | [string](#string) |  | reply with an error message if sign fails |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="tofnd.Multisig"></a>

### Multisig


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `KeyPresence` | [KeyPresenceRequest](#tofnd.KeyPresenceRequest) | [KeyPresenceResponse](#tofnd.KeyPresenceResponse) |  | |
| `Keygen` | [KeygenRequest](#tofnd.KeygenRequest) | [KeygenResponse](#tofnd.KeygenResponse) |  | |
| `Sign` | [SignRequest](#tofnd.SignRequest) | [SignResponse](#tofnd.SignResponse) |  | |

 <!-- end services -->



<a name="scalar/tss/tofnd/v1beta1/tofnd.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/tofnd/v1beta1/tofnd.proto
File copied from golang tofnd with minor tweaks


<a name="tofnd.KeygenInit"></a>

### KeygenInit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `new_key_uid` | [string](#string) |  |  |
| `party_uids` | [string](#string) | repeated |  |
| `party_share_counts` | [uint32](#uint32) | repeated |  |
| `my_party_index` | [uint32](#uint32) |  | parties[my_party_index] belongs to the server |
| `threshold` | [uint32](#uint32) |  |  |






<a name="tofnd.KeygenOutput"></a>

### KeygenOutput
Keygen's success response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_key` | [bytes](#bytes) |  | pub_key; common for all parties |
| `group_recover_info` | [bytes](#bytes) |  | recover info of all parties' shares; common for all parties |
| `private_recover_info` | [bytes](#bytes) |  | private recover info of this party's shares; unique for each party |






<a name="tofnd.MessageIn"></a>

### MessageIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keygen_init` | [KeygenInit](#tofnd.KeygenInit) |  | first message only, Keygen |
| `sign_init` | [SignInit](#tofnd.SignInit) |  | first message only, Sign |
| `traffic` | [TrafficIn](#tofnd.TrafficIn) |  | all subsequent messages |
| `abort` | [bool](#bool) |  | abort the protocol, ignore the bool value |






<a name="tofnd.MessageOut"></a>

### MessageOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `traffic` | [TrafficOut](#tofnd.TrafficOut) |  | all but final message |
| `keygen_result` | [MessageOut.KeygenResult](#tofnd.MessageOut.KeygenResult) |  | final message only, Keygen |
| `sign_result` | [MessageOut.SignResult](#tofnd.MessageOut.SignResult) |  | final message only, Sign |
| `need_recover` | [bool](#bool) |  | issue recover from client |






<a name="tofnd.MessageOut.CriminalList"></a>

### MessageOut.CriminalList
Keygen/Sign failure response message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `criminals` | [MessageOut.CriminalList.Criminal](#tofnd.MessageOut.CriminalList.Criminal) | repeated |  |






<a name="tofnd.MessageOut.CriminalList.Criminal"></a>

### MessageOut.CriminalList.Criminal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `party_uid` | [string](#string) |  |  |
| `crime_type` | [MessageOut.CriminalList.Criminal.CrimeType](#tofnd.MessageOut.CriminalList.Criminal.CrimeType) |  |  |






<a name="tofnd.MessageOut.KeygenResult"></a>

### MessageOut.KeygenResult
Keygen's response types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [KeygenOutput](#tofnd.KeygenOutput) |  | Success response |
| `criminals` | [MessageOut.CriminalList](#tofnd.MessageOut.CriminalList) |  | Faiilure response |






<a name="tofnd.MessageOut.SignResult"></a>

### MessageOut.SignResult
Sign's response types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signature` | [bytes](#bytes) |  | Success response |
| `criminals` | [MessageOut.CriminalList](#tofnd.MessageOut.CriminalList) |  | Failure response |






<a name="tofnd.RecoverRequest"></a>

### RecoverRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keygen_init` | [KeygenInit](#tofnd.KeygenInit) |  |  |
| `keygen_output` | [KeygenOutput](#tofnd.KeygenOutput) |  |  |






<a name="tofnd.RecoverResponse"></a>

### RecoverResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `response` | [RecoverResponse.Response](#tofnd.RecoverResponse.Response) |  |  |






<a name="tofnd.SignInit"></a>

### SignInit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `new_sig_uid` | [string](#string) |  |  |
| `key_uid` | [string](#string) |  |  |
| `party_uids` | [string](#string) | repeated | TODO replace this with a subset of indices? |
| `message_to_sign` | [bytes](#bytes) |  |  |






<a name="tofnd.TrafficIn"></a>

### TrafficIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_party_uid` | [string](#string) |  |  |
| `payload` | [bytes](#bytes) |  |  |
| `is_broadcast` | [bool](#bool) |  |  |






<a name="tofnd.TrafficOut"></a>

### TrafficOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `to_party_uid` | [string](#string) |  |  |
| `payload` | [bytes](#bytes) |  |  |
| `is_broadcast` | [bool](#bool) |  |  |





 <!-- end messages -->


<a name="tofnd.MessageOut.CriminalList.Criminal.CrimeType"></a>

### MessageOut.CriminalList.Criminal.CrimeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CRIME_TYPE_UNSPECIFIED | 0 |  |
| CRIME_TYPE_NON_MALICIOUS | 1 |  |
| CRIME_TYPE_MALICIOUS | 2 |  |



<a name="tofnd.RecoverResponse.Response"></a>

### RecoverResponse.Response


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESPONSE_UNSPECIFIED | 0 |  |
| RESPONSE_SUCCESS | 1 |  |
| RESPONSE_FAIL | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/tss/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/v1beta1/params.proto



<a name="scalar.tss.v1beta1.Params"></a>

### Params
Params is the parameter set for this module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_requirements` | [scalar.tss.exported.v1beta1.KeyRequirement](#scalar.tss.exported.v1beta1.KeyRequirement) | repeated | KeyRequirements defines the requirement for each key role |
| `suspend_duration_in_blocks` | [int64](#int64) |  | SuspendDurationInBlocks defines the number of blocks a validator is disallowed to participate in any TSS ceremony after committing a malicious behaviour during signing |
| `heartbeat_period_in_blocks` | [int64](#int64) |  | HeartBeatPeriodInBlocks defines the time period in blocks for tss to emit the event asking validators to send their heartbeats |
| `max_missed_blocks_per_window` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `unbonding_locking_key_rotation_count` | [int64](#int64) |  |  |
| `external_multisig_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `max_sign_queue_size` | [int64](#int64) |  |  |
| `max_simultaneous_sign_shares` | [int64](#int64) |  |  |
| `tss_signed_blocks_window` | [int64](#int64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/tss/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/v1beta1/types.proto



<a name="scalar.tss.v1beta1.ExternalKeys"></a>

### ExternalKeys



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `key_ids` | [string](#string) | repeated |  |






<a name="scalar.tss.v1beta1.KeyInfo"></a>

### KeyInfo
KeyInfo holds information about a key


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |
| `key_role` | [scalar.tss.exported.v1beta1.KeyRole](#scalar.tss.exported.v1beta1.KeyRole) |  |  |
| `key_type` | [scalar.tss.exported.v1beta1.KeyType](#scalar.tss.exported.v1beta1.KeyType) |  |  |






<a name="scalar.tss.v1beta1.KeyRecoveryInfo"></a>

### KeyRecoveryInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |
| `public` | [bytes](#bytes) |  |  |
| `private` | [KeyRecoveryInfo.PrivateEntry](#scalar.tss.v1beta1.KeyRecoveryInfo.PrivateEntry) | repeated |  |






<a name="scalar.tss.v1beta1.KeyRecoveryInfo.PrivateEntry"></a>

### KeyRecoveryInfo.PrivateEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [bytes](#bytes) |  |  |






<a name="scalar.tss.v1beta1.KeygenVoteData"></a>

### KeygenVoteData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_key` | [bytes](#bytes) |  |  |
| `group_recovery_info` | [bytes](#bytes) |  |  |






<a name="scalar.tss.v1beta1.MultisigInfo"></a>

### MultisigInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `timeout` | [int64](#int64) |  |  |
| `target_num` | [int64](#int64) |  |  |
| `infos` | [MultisigInfo.Info](#scalar.tss.v1beta1.MultisigInfo.Info) | repeated |  |






<a name="scalar.tss.v1beta1.MultisigInfo.Info"></a>

### MultisigInfo.Info



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `participant` | [bytes](#bytes) |  |  |
| `data` | [bytes](#bytes) | repeated |  |






<a name="scalar.tss.v1beta1.ValidatorStatus"></a>

### ValidatorStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [bytes](#bytes) |  |  |
| `suspended_until` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/tss/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/v1beta1/genesis.proto



<a name="scalar.tss.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.tss.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/tss/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/v1beta1/query.proto



<a name="scalar.tss.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.tss.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.tss.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/tss/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/v1beta1/tx.proto



<a name="scalar.tss.v1beta1.HeartBeatRequest"></a>

### HeartBeatRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `key_ids` | [string](#string) | repeated | **Deprecated.** Deprecated: this field will be removed in the next release

key_ids was deprecated in v1.0 |






<a name="scalar.tss.v1beta1.HeartBeatResponse"></a>

### HeartBeatResponse







<a name="scalar.tss.v1beta1.ProcessKeygenTrafficRequest"></a>

### ProcessKeygenTrafficRequest
ProcessKeygenTrafficRequest protocol message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `session_id` | [string](#string) |  |  |
| `payload` | [tofnd.TrafficOut](#tofnd.TrafficOut) |  |  |






<a name="scalar.tss.v1beta1.ProcessKeygenTrafficResponse"></a>

### ProcessKeygenTrafficResponse







<a name="scalar.tss.v1beta1.ProcessSignTrafficRequest"></a>

### ProcessSignTrafficRequest
ProcessSignTrafficRequest protocol message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `session_id` | [string](#string) |  |  |
| `payload` | [tofnd.TrafficOut](#tofnd.TrafficOut) |  |  |






<a name="scalar.tss.v1beta1.ProcessSignTrafficResponse"></a>

### ProcessSignTrafficResponse







<a name="scalar.tss.v1beta1.RegisterExternalKeysRequest"></a>

### RegisterExternalKeysRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `external_keys` | [RegisterExternalKeysRequest.ExternalKey](#scalar.tss.v1beta1.RegisterExternalKeysRequest.ExternalKey) | repeated |  |






<a name="scalar.tss.v1beta1.RegisterExternalKeysRequest.ExternalKey"></a>

### RegisterExternalKeysRequest.ExternalKey



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `pub_key` | [bytes](#bytes) |  |  |






<a name="scalar.tss.v1beta1.RegisterExternalKeysResponse"></a>

### RegisterExternalKeysResponse







<a name="scalar.tss.v1beta1.RotateKeyRequest"></a>

### RotateKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `key_role` | [scalar.tss.exported.v1beta1.KeyRole](#scalar.tss.exported.v1beta1.KeyRole) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.tss.v1beta1.RotateKeyResponse"></a>

### RotateKeyResponse







<a name="scalar.tss.v1beta1.StartKeygenRequest"></a>

### StartKeygenRequest
StartKeygenRequest indicate the start of keygen


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `key_info` | [KeyInfo](#scalar.tss.v1beta1.KeyInfo) |  |  |






<a name="scalar.tss.v1beta1.StartKeygenResponse"></a>

### StartKeygenResponse







<a name="scalar.tss.v1beta1.SubmitMultisigPubKeysRequest"></a>

### SubmitMultisigPubKeysRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `key_id` | [string](#string) |  |  |
| `sig_key_pairs` | [scalar.tss.exported.v1beta1.SigKeyPair](#scalar.tss.exported.v1beta1.SigKeyPair) | repeated |  |






<a name="scalar.tss.v1beta1.SubmitMultisigPubKeysResponse"></a>

### SubmitMultisigPubKeysResponse







<a name="scalar.tss.v1beta1.SubmitMultisigSignaturesRequest"></a>

### SubmitMultisigSignaturesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `sig_id` | [string](#string) |  |  |
| `signatures` | [bytes](#bytes) | repeated |  |






<a name="scalar.tss.v1beta1.SubmitMultisigSignaturesResponse"></a>

### SubmitMultisigSignaturesResponse







<a name="scalar.tss.v1beta1.VotePubKeyRequest"></a>

### VotePubKeyRequest
VotePubKeyRequest represents the message to vote on a public key


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `poll_key` | [scalar.vote.exported.v1beta1.PollKey](#scalar.vote.exported.v1beta1.PollKey) |  |  |
| `result` | [tofnd.MessageOut.KeygenResult](#tofnd.MessageOut.KeygenResult) |  |  |






<a name="scalar.tss.v1beta1.VotePubKeyResponse"></a>

### VotePubKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `log` | [string](#string) |  |  |






<a name="scalar.tss.v1beta1.VoteSigRequest"></a>

### VoteSigRequest
VoteSigRequest represents a message to vote for a signature


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `poll_key` | [scalar.vote.exported.v1beta1.PollKey](#scalar.vote.exported.v1beta1.PollKey) |  |  |
| `result` | [tofnd.MessageOut.SignResult](#tofnd.MessageOut.SignResult) |  |  |






<a name="scalar.tss.v1beta1.VoteSigResponse"></a>

### VoteSigResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `log` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/tss/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/tss/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.tss.v1beta1.MsgService"></a>

### MsgService
Msg defines the tss Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `HeartBeat` | [HeartBeatRequest](#scalar.tss.v1beta1.HeartBeatRequest) | [HeartBeatResponse](#scalar.tss.v1beta1.HeartBeatResponse) |  | POST|/scalar/tss/heartbeat|


<a name="scalar.tss.v1beta1.QueryService"></a>

### QueryService
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [ParamsRequest](#scalar.tss.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.tss.v1beta1.ParamsResponse) |  | GET|/scalar/tss/v1beta1/params|

 <!-- end services -->



<a name="scalar/vote/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/vote/v1beta1/events.proto



<a name="scalar.vote.v1beta1.Voted"></a>

### Voted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `action` | [string](#string) |  |  |
| `poll` | [string](#string) |  |  |
| `voter` | [string](#string) |  |  |
| `state` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/vote/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/vote/v1beta1/params.proto



<a name="scalar.vote.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `default_voting_threshold` | [scalar.utils.v1beta1.Threshold](#scalar.utils.v1beta1.Threshold) |  |  |
| `end_blocker_limit` | [int64](#int64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/vote/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/vote/v1beta1/genesis.proto



<a name="scalar.vote.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.vote.v1beta1.Params) |  |  |
| `poll_metadatas` | [scalar.vote.exported.v1beta1.PollMetadata](#scalar.vote.exported.v1beta1.PollMetadata) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/vote/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/vote/v1beta1/query.proto



<a name="scalar.vote.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params






<a name="scalar.vote.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.vote.v1beta1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/vote/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/vote/v1beta1/types.proto



<a name="scalar.vote.v1beta1.TalliedVote"></a>

### TalliedVote
TalliedVote represents a vote for a poll with the accumulated stake of all
validators voting for the same data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tally` | [bytes](#bytes) |  |  |
| `data` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |
| `is_voter_late` | [TalliedVote.IsVoterLateEntry](#scalar.vote.v1beta1.TalliedVote.IsVoterLateEntry) | repeated |  |






<a name="scalar.vote.v1beta1.TalliedVote.IsVoterLateEntry"></a>

### TalliedVote.IsVoterLateEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [bool](#bool) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/vote/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/vote/v1beta1/tx.proto



<a name="scalar.vote.v1beta1.VoteRequest"></a>

### VoteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |
| `vote` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="scalar.vote.v1beta1.VoteResponse"></a>

### VoteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `log` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/vote/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/vote/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.vote.v1beta1.MsgService"></a>

### MsgService
Msg defines the vote Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Vote` | [VoteRequest](#scalar.vote.v1beta1.VoteRequest) | [VoteResponse](#scalar.vote.v1beta1.VoteResponse) |  | POST|/scalar/vote/vote|


<a name="scalar.vote.v1beta1.QueryService"></a>

### QueryService
QueryService defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [ParamsRequest](#scalar.vote.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.vote.v1beta1.ParamsResponse) |  | GET|/scalar/vote/v1beta1/params|

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

