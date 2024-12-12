<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [scalar/btc/v1beta1/types.proto](#scalar/btc/v1beta1/types.proto)
    - [Command](#scalar.btc.v1beta1.Command)
    - [CommandBatchMetadata](#scalar.btc.v1beta1.CommandBatchMetadata)
    - [Proof](#scalar.btc.v1beta1.Proof)
    - [StakingTx](#scalar.btc.v1beta1.StakingTx)
    - [StakingTxMetadata](#scalar.btc.v1beta1.StakingTxMetadata)
  
    - [BatchedCommandsStatus](#scalar.btc.v1beta1.BatchedCommandsStatus)
    - [CommandType](#scalar.btc.v1beta1.CommandType)
    - [NetworkKind](#scalar.btc.v1beta1.NetworkKind)
    - [StakingTxStatus](#scalar.btc.v1beta1.StakingTxStatus)
  
- [scalar/btc/v1beta1/events.proto](#scalar/btc/v1beta1/events.proto)
    - [Event](#scalar.btc.v1beta1.Event)
    - [EventStakingTx](#scalar.btc.v1beta1.EventStakingTx)
    - [VoteEvents](#scalar.btc.v1beta1.VoteEvents)
  
    - [Event.Status](#scalar.btc.v1beta1.Event.Status)
  
- [scalar/btc/v1beta1/params.proto](#scalar/btc/v1beta1/params.proto)
    - [Params](#scalar.btc.v1beta1.Params)
  
- [scalar/btc/v1beta1/genesis.proto](#scalar/btc/v1beta1/genesis.proto)
    - [GenesisState](#scalar.btc.v1beta1.GenesisState)
    - [GenesisState.Chain](#scalar.btc.v1beta1.GenesisState.Chain)
  
- [scalar/btc/v1beta1/poll.proto](#scalar/btc/v1beta1/poll.proto)
    - [PollCompleted](#scalar.btc.v1beta1.PollCompleted)
    - [PollExpired](#scalar.btc.v1beta1.PollExpired)
    - [PollFailed](#scalar.btc.v1beta1.PollFailed)
    - [PollMapping](#scalar.btc.v1beta1.PollMapping)
    - [PollMetadata](#scalar.btc.v1beta1.PollMetadata)
  
- [scalar/btc/v1beta1/tx.proto](#scalar/btc/v1beta1/tx.proto)
    - [ConfirmStakingTxsRequest](#scalar.btc.v1beta1.ConfirmStakingTxsRequest)
    - [ConfirmStakingTxsResponse](#scalar.btc.v1beta1.ConfirmStakingTxsResponse)
    - [EventConfirmStakingTxsStarted](#scalar.btc.v1beta1.EventConfirmStakingTxsStarted)
  
- [scalar/btc/v1beta1/query.proto](#scalar/btc/v1beta1/query.proto)
    - [BatchedCommandsRequest](#scalar.btc.v1beta1.BatchedCommandsRequest)
    - [BatchedCommandsResponse](#scalar.btc.v1beta1.BatchedCommandsResponse)
  
- [scalar/btc/v1beta1/service.proto](#scalar/btc/v1beta1/service.proto)
    - [MsgService](#scalar.btc.v1beta1.MsgService)
    - [QueryService](#scalar.btc.v1beta1.QueryService)
  
- [scalar/btc/v1beta1/vote.proto](#scalar/btc/v1beta1/vote.proto)
    - [BTCEventConfirmed](#scalar.btc.v1beta1.BTCEventConfirmed)
    - [NoEventsConfirmed](#scalar.btc.v1beta1.NoEventsConfirmed)
  
- [scalar/covenant/v1beta1/types.proto](#scalar/covenant/v1beta1/types.proto)
    - [Covenant](#scalar.covenant.v1beta1.Covenant)
    - [CovenantGroup](#scalar.covenant.v1beta1.CovenantGroup)
  
- [scalar/covenant/v1beta1/genesis.proto](#scalar/covenant/v1beta1/genesis.proto)
    - [GenesisState](#scalar.covenant.v1beta1.GenesisState)
  
- [scalar/covenant/v1beta1/query.proto](#scalar/covenant/v1beta1/query.proto)
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
  
- [scalar/multisig/exported/v1beta1/types.proto](#scalar/multisig/exported/v1beta1/types.proto)
    - [KeyState](#scalar.multisig.exported.v1beta1.KeyState)
    - [MultisigState](#scalar.multisig.exported.v1beta1.MultisigState)
  
- [scalar/evm/v1beta1/types.proto](#scalar/evm/v1beta1/types.proto)
    - [Asset](#scalar.evm.v1beta1.Asset)
    - [BurnerInfo](#scalar.evm.v1beta1.BurnerInfo)
    - [Command](#scalar.evm.v1beta1.Command)
    - [CommandBatchMetadata](#scalar.evm.v1beta1.CommandBatchMetadata)
    - [ERC20Deposit](#scalar.evm.v1beta1.ERC20Deposit)
    - [ERC20TokenMetadata](#scalar.evm.v1beta1.ERC20TokenMetadata)
    - [Event](#scalar.evm.v1beta1.Event)
    - [EventContractCall](#scalar.evm.v1beta1.EventContractCall)
    - [EventContractCallWithToken](#scalar.evm.v1beta1.EventContractCallWithToken)
    - [EventMultisigOperatorshipTransferred](#scalar.evm.v1beta1.EventMultisigOperatorshipTransferred)
    - [EventMultisigOwnershipTransferred](#scalar.evm.v1beta1.EventMultisigOwnershipTransferred)
    - [EventTokenDeployed](#scalar.evm.v1beta1.EventTokenDeployed)
    - [EventTokenSent](#scalar.evm.v1beta1.EventTokenSent)
    - [EventTransfer](#scalar.evm.v1beta1.EventTransfer)
    - [Gateway](#scalar.evm.v1beta1.Gateway)
    - [NetworkInfo](#scalar.evm.v1beta1.NetworkInfo)
    - [PollMetadata](#scalar.evm.v1beta1.PollMetadata)
    - [SigMetadata](#scalar.evm.v1beta1.SigMetadata)
    - [TokenDetails](#scalar.evm.v1beta1.TokenDetails)
    - [TransactionMetadata](#scalar.evm.v1beta1.TransactionMetadata)
    - [TransferKey](#scalar.evm.v1beta1.TransferKey)
    - [VoteEvents](#scalar.evm.v1beta1.VoteEvents)
  
    - [BatchedCommandsStatus](#scalar.evm.v1beta1.BatchedCommandsStatus)
    - [CommandType](#scalar.evm.v1beta1.CommandType)
    - [DepositStatus](#scalar.evm.v1beta1.DepositStatus)
    - [Event.Status](#scalar.evm.v1beta1.Event.Status)
    - [SigType](#scalar.evm.v1beta1.SigType)
    - [Status](#scalar.evm.v1beta1.Status)
  
- [scalar/evm/v1beta1/events.proto](#scalar/evm/v1beta1/events.proto)
    - [BurnCommand](#scalar.evm.v1beta1.BurnCommand)
    - [ChainAdded](#scalar.evm.v1beta1.ChainAdded)
    - [CommandBatchAborted](#scalar.evm.v1beta1.CommandBatchAborted)
    - [CommandBatchSigned](#scalar.evm.v1beta1.CommandBatchSigned)
    - [ConfirmDepositStarted](#scalar.evm.v1beta1.ConfirmDepositStarted)
    - [ConfirmGatewayTxStarted](#scalar.evm.v1beta1.ConfirmGatewayTxStarted)
    - [ConfirmGatewayTxsStarted](#scalar.evm.v1beta1.ConfirmGatewayTxsStarted)
    - [ConfirmKeyTransferStarted](#scalar.evm.v1beta1.ConfirmKeyTransferStarted)
    - [ConfirmTokenStarted](#scalar.evm.v1beta1.ConfirmTokenStarted)
    - [ContractCallApproved](#scalar.evm.v1beta1.ContractCallApproved)
    - [ContractCallFailed](#scalar.evm.v1beta1.ContractCallFailed)
    - [ContractCallWithMintApproved](#scalar.evm.v1beta1.ContractCallWithMintApproved)
    - [EVMEventCompleted](#scalar.evm.v1beta1.EVMEventCompleted)
    - [EVMEventConfirmed](#scalar.evm.v1beta1.EVMEventConfirmed)
    - [EVMEventFailed](#scalar.evm.v1beta1.EVMEventFailed)
    - [EVMEventRetryFailed](#scalar.evm.v1beta1.EVMEventRetryFailed)
    - [MintCommand](#scalar.evm.v1beta1.MintCommand)
    - [NoEventsConfirmed](#scalar.evm.v1beta1.NoEventsConfirmed)
    - [PollCompleted](#scalar.evm.v1beta1.PollCompleted)
    - [PollExpired](#scalar.evm.v1beta1.PollExpired)
    - [PollFailed](#scalar.evm.v1beta1.PollFailed)
    - [PollMapping](#scalar.evm.v1beta1.PollMapping)
    - [TokenSent](#scalar.evm.v1beta1.TokenSent)
  
- [scalar/evm/v1beta1/params.proto](#scalar/evm/v1beta1/params.proto)
    - [Params](#scalar.evm.v1beta1.Params)
    - [PendingChain](#scalar.evm.v1beta1.PendingChain)
  
- [scalar/evm/v1beta1/genesis.proto](#scalar/evm/v1beta1/genesis.proto)
    - [GenesisState](#scalar.evm.v1beta1.GenesisState)
    - [GenesisState.Chain](#scalar.evm.v1beta1.GenesisState.Chain)
  
- [scalar/evm/v1beta1/query.proto](#scalar/evm/v1beta1/query.proto)
    - [BatchedCommandsRequest](#scalar.evm.v1beta1.BatchedCommandsRequest)
    - [BatchedCommandsResponse](#scalar.evm.v1beta1.BatchedCommandsResponse)
    - [BurnerInfoRequest](#scalar.evm.v1beta1.BurnerInfoRequest)
    - [BurnerInfoResponse](#scalar.evm.v1beta1.BurnerInfoResponse)
    - [BytecodeRequest](#scalar.evm.v1beta1.BytecodeRequest)
    - [BytecodeResponse](#scalar.evm.v1beta1.BytecodeResponse)
    - [ChainsRequest](#scalar.evm.v1beta1.ChainsRequest)
    - [ChainsResponse](#scalar.evm.v1beta1.ChainsResponse)
    - [CommandRequest](#scalar.evm.v1beta1.CommandRequest)
    - [CommandResponse](#scalar.evm.v1beta1.CommandResponse)
    - [CommandResponse.ParamsEntry](#scalar.evm.v1beta1.CommandResponse.ParamsEntry)
    - [ConfirmationHeightRequest](#scalar.evm.v1beta1.ConfirmationHeightRequest)
    - [ConfirmationHeightResponse](#scalar.evm.v1beta1.ConfirmationHeightResponse)
    - [DepositQueryParams](#scalar.evm.v1beta1.DepositQueryParams)
    - [DepositStateRequest](#scalar.evm.v1beta1.DepositStateRequest)
    - [DepositStateResponse](#scalar.evm.v1beta1.DepositStateResponse)
    - [ERC20TokensRequest](#scalar.evm.v1beta1.ERC20TokensRequest)
    - [ERC20TokensResponse](#scalar.evm.v1beta1.ERC20TokensResponse)
    - [ERC20TokensResponse.Token](#scalar.evm.v1beta1.ERC20TokensResponse.Token)
    - [EventRequest](#scalar.evm.v1beta1.EventRequest)
    - [EventResponse](#scalar.evm.v1beta1.EventResponse)
    - [GatewayAddressRequest](#scalar.evm.v1beta1.GatewayAddressRequest)
    - [GatewayAddressResponse](#scalar.evm.v1beta1.GatewayAddressResponse)
    - [KeyAddressRequest](#scalar.evm.v1beta1.KeyAddressRequest)
    - [KeyAddressResponse](#scalar.evm.v1beta1.KeyAddressResponse)
    - [KeyAddressResponse.WeightedAddress](#scalar.evm.v1beta1.KeyAddressResponse.WeightedAddress)
    - [ParamsRequest](#scalar.evm.v1beta1.ParamsRequest)
    - [ParamsResponse](#scalar.evm.v1beta1.ParamsResponse)
    - [PendingCommandsRequest](#scalar.evm.v1beta1.PendingCommandsRequest)
    - [PendingCommandsResponse](#scalar.evm.v1beta1.PendingCommandsResponse)
    - [Proof](#scalar.evm.v1beta1.Proof)
    - [QueryBurnerAddressResponse](#scalar.evm.v1beta1.QueryBurnerAddressResponse)
    - [QueryCommandResponse](#scalar.evm.v1beta1.QueryCommandResponse)
    - [QueryCommandResponse.ParamsEntry](#scalar.evm.v1beta1.QueryCommandResponse.ParamsEntry)
    - [QueryDepositStateParams](#scalar.evm.v1beta1.QueryDepositStateParams)
    - [QueryTokenAddressResponse](#scalar.evm.v1beta1.QueryTokenAddressResponse)
    - [TokenInfoRequest](#scalar.evm.v1beta1.TokenInfoRequest)
    - [TokenInfoResponse](#scalar.evm.v1beta1.TokenInfoResponse)
  
    - [ChainStatus](#scalar.evm.v1beta1.ChainStatus)
    - [TokenType](#scalar.evm.v1beta1.TokenType)
  
- [scalar/evm/v1beta1/tx.proto](#scalar/evm/v1beta1/tx.proto)
    - [AddChainRequest](#scalar.evm.v1beta1.AddChainRequest)
    - [AddChainResponse](#scalar.evm.v1beta1.AddChainResponse)
    - [ConfirmDepositRequest](#scalar.evm.v1beta1.ConfirmDepositRequest)
    - [ConfirmDepositResponse](#scalar.evm.v1beta1.ConfirmDepositResponse)
    - [ConfirmGatewayTxRequest](#scalar.evm.v1beta1.ConfirmGatewayTxRequest)
    - [ConfirmGatewayTxResponse](#scalar.evm.v1beta1.ConfirmGatewayTxResponse)
    - [ConfirmGatewayTxsRequest](#scalar.evm.v1beta1.ConfirmGatewayTxsRequest)
    - [ConfirmGatewayTxsResponse](#scalar.evm.v1beta1.ConfirmGatewayTxsResponse)
    - [ConfirmTokenRequest](#scalar.evm.v1beta1.ConfirmTokenRequest)
    - [ConfirmTokenResponse](#scalar.evm.v1beta1.ConfirmTokenResponse)
    - [ConfirmTransferKeyRequest](#scalar.evm.v1beta1.ConfirmTransferKeyRequest)
    - [ConfirmTransferKeyResponse](#scalar.evm.v1beta1.ConfirmTransferKeyResponse)
    - [CreateBurnTokensRequest](#scalar.evm.v1beta1.CreateBurnTokensRequest)
    - [CreateBurnTokensResponse](#scalar.evm.v1beta1.CreateBurnTokensResponse)
    - [CreateDeployTokenRequest](#scalar.evm.v1beta1.CreateDeployTokenRequest)
    - [CreateDeployTokenResponse](#scalar.evm.v1beta1.CreateDeployTokenResponse)
    - [CreatePendingTransfersRequest](#scalar.evm.v1beta1.CreatePendingTransfersRequest)
    - [CreatePendingTransfersResponse](#scalar.evm.v1beta1.CreatePendingTransfersResponse)
    - [CreateTransferOperatorshipRequest](#scalar.evm.v1beta1.CreateTransferOperatorshipRequest)
    - [CreateTransferOperatorshipResponse](#scalar.evm.v1beta1.CreateTransferOperatorshipResponse)
    - [CreateTransferOwnershipRequest](#scalar.evm.v1beta1.CreateTransferOwnershipRequest)
    - [CreateTransferOwnershipResponse](#scalar.evm.v1beta1.CreateTransferOwnershipResponse)
    - [LinkRequest](#scalar.evm.v1beta1.LinkRequest)
    - [LinkResponse](#scalar.evm.v1beta1.LinkResponse)
    - [RetryFailedEventRequest](#scalar.evm.v1beta1.RetryFailedEventRequest)
    - [RetryFailedEventResponse](#scalar.evm.v1beta1.RetryFailedEventResponse)
    - [SetGatewayRequest](#scalar.evm.v1beta1.SetGatewayRequest)
    - [SetGatewayResponse](#scalar.evm.v1beta1.SetGatewayResponse)
    - [SignCommandsRequest](#scalar.evm.v1beta1.SignCommandsRequest)
    - [SignCommandsResponse](#scalar.evm.v1beta1.SignCommandsResponse)
  
- [scalar/evm/v1beta1/service.proto](#scalar/evm/v1beta1/service.proto)
    - [MsgService](#scalar.evm.v1beta1.MsgService)
    - [QueryService](#scalar.evm.v1beta1.QueryService)
  
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
  
- [scalar/multisig/v1beta1/params.proto](#scalar/multisig/v1beta1/params.proto)
    - [Params](#scalar.multisig.v1beta1.Params)
  
- [scalar/multisig/v1beta1/types.proto](#scalar/multisig/v1beta1/types.proto)
    - [Key](#scalar.multisig.v1beta1.Key)
    - [Key.PubKeysEntry](#scalar.multisig.v1beta1.Key.PubKeysEntry)
    - [KeyEpoch](#scalar.multisig.v1beta1.KeyEpoch)
    - [KeygenSession](#scalar.multisig.v1beta1.KeygenSession)
    - [KeygenSession.IsPubKeyReceivedEntry](#scalar.multisig.v1beta1.KeygenSession.IsPubKeyReceivedEntry)
    - [MultiSig](#scalar.multisig.v1beta1.MultiSig)
    - [MultiSig.SigsEntry](#scalar.multisig.v1beta1.MultiSig.SigsEntry)
    - [SigningSession](#scalar.multisig.v1beta1.SigningSession)
  
- [scalar/multisig/v1beta1/genesis.proto](#scalar/multisig/v1beta1/genesis.proto)
    - [GenesisState](#scalar.multisig.v1beta1.GenesisState)
  
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
  
- [scalar/protocol/v1beta1/types.proto](#scalar/protocol/v1beta1/types.proto)
    - [Protocol](#scalar.protocol.v1beta1.Protocol)
  
- [scalar/protocol/v1beta1/genesis.proto](#scalar/protocol/v1beta1/genesis.proto)
    - [GenesisState](#scalar.protocol.v1beta1.GenesisState)
  
- [scalar/protocol/v1beta1/query.proto](#scalar/protocol/v1beta1/query.proto)
    - [ProtocolRequest](#scalar.protocol.v1beta1.ProtocolRequest)
    - [ProtocolResponse](#scalar.protocol.v1beta1.ProtocolResponse)
  
    - [ProtocolStatus](#scalar.protocol.v1beta1.ProtocolStatus)
  
- [scalar/protocol/v1beta1/service.proto](#scalar/protocol/v1beta1/service.proto)
- [scalar/protocol/v1beta1/tx.proto](#scalar/protocol/v1beta1/tx.proto)
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
  
- [Scalar Value Types](#scalar-value-types)



<a name="scalar/btc/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/types.proto



<a name="scalar.btc.v1beta1.Command"></a>

### Command



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [bytes](#bytes) |  |  |
| `params` | [bytes](#bytes) |  |  |
| `key_id` | [string](#string) |  |  |
| `type` | [CommandType](#scalar.btc.v1beta1.CommandType) |  |  |






<a name="scalar.btc.v1beta1.CommandBatchMetadata"></a>

### CommandBatchMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [bytes](#bytes) |  |  |
| `command_ids` | [bytes](#bytes) | repeated |  |
| `data` | [bytes](#bytes) |  |  |
| `sig_hash` | [bytes](#bytes) |  |  |
| `status` | [BatchedCommandsStatus](#scalar.btc.v1beta1.BatchedCommandsStatus) |  |  |
| `key_id` | [string](#string) |  |  |
| `prev_batched_commands_id` | [bytes](#bytes) |  |  |
| `signature` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="scalar.btc.v1beta1.Proof"></a>

### Proof



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [string](#string) | repeated |  |
| `weights` | [string](#string) | repeated |  |
| `threshold` | [string](#string) |  |  |
| `signatures` | [string](#string) | repeated |  |






<a name="scalar.btc.v1beta1.StakingTx"></a>

### StakingTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `amount` | [bytes](#bytes) |  |  |
| `asset` | [string](#string) |  | TODO: change to asset type: sats, runes, btc, etc |
| `destination_chain` | [string](#string) |  |  |
| `destination_recipient_address` | [bytes](#bytes) |  |  |
| `log_index` | [uint64](#uint64) |  |  |






<a name="scalar.btc.v1beta1.StakingTxMetadata"></a>

### StakingTxMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tag` | [bytes](#bytes) |  |  |
| `version` | [bytes](#bytes) |  |  |
| `network_id` | [NetworkKind](#scalar.btc.v1beta1.NetworkKind) |  |  |
| `flags` | [uint32](#uint32) |  |  |
| `service_tag` | [bytes](#bytes) |  |  |
| `have_only_covenants` | [bool](#bool) |  |  |
| `covenant_quorum` | [uint32](#uint32) |  |  |
| `destination_chain_type` | [uint32](#uint32) |  |  |
| `destination_chain_id` | [uint64](#uint64) |  |  |
| `destination_contract_address` | [bytes](#bytes) |  |  |
| `destination_recipient_address` | [bytes](#bytes) |  |  |





 <!-- end messages -->


<a name="scalar.btc.v1beta1.BatchedCommandsStatus"></a>

### BatchedCommandsStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| BATCHED_COMMANDS_STATUS_UNSPECIFIED | 0 |  |
| BATCHED_COMMANDS_STATUS_SIGNING | 1 |  |
| BATCHED_COMMANDS_STATUS_ABORTED | 2 |  |
| BATCHED_COMMANDS_STATUS_SIGNED | 3 |  |



<a name="scalar.btc.v1beta1.CommandType"></a>

### CommandType


| Name | Number | Description |
| ---- | ------ | ----------- |
| COMMAND_TYPE_APPROVE_CONTRACT_CALL | 0 |  |



<a name="scalar.btc.v1beta1.NetworkKind"></a>

### NetworkKind


| Name | Number | Description |
| ---- | ------ | ----------- |
| NETWORK_KIND_MAINNET | 0 |  |
| NETWORK_KIND_TESTNET | 1 |  |



<a name="scalar.btc.v1beta1.StakingTxStatus"></a>

### StakingTxStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| STAKING_TX_STATUS_UNSPECIFIED | 0 |  |
| STAKING_TX_STATUS_PENDING | 1 |  |
| STAKING_TX_STATUS_CONFIRMED | 2 |  |
| STAKING_TX_STATUS_COMPLETED | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/btc/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/events.proto



<a name="scalar.btc.v1beta1.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `status` | [Event.Status](#scalar.btc.v1beta1.Event.Status) |  |  |
| `index` | [uint64](#uint64) |  |  |
| `staking_tx` | [EventStakingTx](#scalar.btc.v1beta1.EventStakingTx) |  |  |






<a name="scalar.btc.v1beta1.EventStakingTx"></a>

### EventStakingTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prev_out_point` | [string](#string) |  |  |
| `amount` | [uint64](#uint64) |  |  |
| `asset` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `metadata` | [StakingTxMetadata](#scalar.btc.v1beta1.StakingTxMetadata) |  |  |






<a name="scalar.btc.v1beta1.VoteEvents"></a>

### VoteEvents



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `events` | [Event](#scalar.btc.v1beta1.Event) | repeated |  |





 <!-- end messages -->


<a name="scalar.btc.v1beta1.Event.Status"></a>

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



<a name="scalar/btc/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/params.proto



<a name="scalar.btc.v1beta1.Params"></a>

### Params



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain_name` | [string](#string) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `network_kind` | [NetworkKind](#scalar.btc.v1beta1.NetworkKind) |  |  |
| `revote_locking_period` | [int64](#int64) |  |  |
| `chain_id` | [uint64](#uint64) |  |  |
| `voting_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
| `min_voter_count` | [int64](#int64) |  |  |
| `voting_grace_period` | [int64](#int64) |  |  |
| `end_blocker_limit` | [int64](#int64) |  |  |
| `transfer_limit` | [uint64](#uint64) |  |  |
| `vault_tag` | [bytes](#bytes) |  |  |
| `vault_version` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/btc/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/genesis.proto



<a name="scalar.btc.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chains` | [GenesisState.Chain](#scalar.btc.v1beta1.GenesisState.Chain) | repeated |  |






<a name="scalar.btc.v1beta1.GenesisState.Chain"></a>

### GenesisState.Chain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.btc.v1beta1.Params) |  |  |
| `command_queue` | [axelar.utils.v1beta1.QueueState](#axelar.utils.v1beta1.QueueState) |  |  |
| `confirmed_staking_txs` | [StakingTx](#scalar.btc.v1beta1.StakingTx) | repeated |  |
| `command_batches` | [CommandBatchMetadata](#scalar.btc.v1beta1.CommandBatchMetadata) | repeated |  |
| `events` | [Event](#scalar.btc.v1beta1.Event) | repeated |  |
| `confirmed_event_queue` | [axelar.utils.v1beta1.QueueState](#axelar.utils.v1beta1.QueueState) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/btc/v1beta1/poll.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/poll.proto



<a name="scalar.btc.v1beta1.PollCompleted"></a>

### PollCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.btc.v1beta1.PollExpired"></a>

### PollExpired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.btc.v1beta1.PollFailed"></a>

### PollFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.btc.v1beta1.PollMapping"></a>

### PollMapping



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.btc.v1beta1.PollMetadata"></a>

### PollMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/btc/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/tx.proto



<a name="scalar.btc.v1beta1.ConfirmStakingTxsRequest"></a>

### ConfirmStakingTxsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_ids` | [bytes](#bytes) | repeated |  |






<a name="scalar.btc.v1beta1.ConfirmStakingTxsResponse"></a>

### ConfirmStakingTxsResponse







<a name="scalar.btc.v1beta1.EventConfirmStakingTxsStarted"></a>

### EventConfirmStakingTxsStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poll_mappings` | [PollMapping](#scalar.btc.v1beta1.PollMapping) | repeated |  |
| `chain` | [string](#string) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [bytes](#bytes) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/btc/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/query.proto



<a name="scalar.btc.v1beta1.BatchedCommandsRequest"></a>

### BatchedCommandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `id` | [string](#string) |  | id defines an optional id for the commandsbatch. If not specified the latest will be returned |






<a name="scalar.btc.v1beta1.BatchedCommandsResponse"></a>

### BatchedCommandsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |
| `status` | [BatchedCommandsStatus](#scalar.btc.v1beta1.BatchedCommandsStatus) |  |  |
| `key_id` | [string](#string) |  |  |
| `execute_data` | [string](#string) |  |  |
| `prev_batched_commands_id` | [string](#string) |  |  |
| `command_ids` | [string](#string) | repeated |  |
| `proof` | [Proof](#scalar.btc.v1beta1.Proof) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/btc/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.btc.v1beta1.MsgService"></a>

### MsgService
Msg defines the evm Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ConfirmStakingTxs` | [ConfirmStakingTxsRequest](#scalar.btc.v1beta1.ConfirmStakingTxsRequest) | [ConfirmStakingTxsResponse](#scalar.btc.v1beta1.ConfirmStakingTxsResponse) |  | POST|/scalar/btc/confirm_staking_txs|


<a name="scalar.btc.v1beta1.QueryService"></a>

### QueryService
QueryService defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `BatchedCommands` | [BatchedCommandsRequest](#scalar.btc.v1beta1.BatchedCommandsRequest) | [BatchedCommandsResponse](#scalar.btc.v1beta1.BatchedCommandsResponse) | BatchedCommands queries the batched commands for a specified chain and BatchedCommandsID if no BatchedCommandsID is specified, then it returns the latest batched commands | GET|/scalar/btc/v1beta1/batched_commands/{chain}/{id}|

 <!-- end services -->



<a name="scalar/btc/v1beta1/vote.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/btc/v1beta1/vote.proto



<a name="scalar.btc.v1beta1.BTCEventConfirmed"></a>

### BTCEventConfirmed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.btc.v1beta1.NoEventsConfirmed"></a>

### NoEventsConfirmed
Vote handler


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/types.proto



<a name="scalar.covenant.v1beta1.Covenant"></a>

### Covenant



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `btcpubkey` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |






<a name="scalar.covenant.v1beta1.CovenantGroup"></a>

### CovenantGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_hash` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `covenants` | [Covenant](#scalar.covenant.v1beta1.Covenant) | repeated |  |





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
| `covenants` | [Covenant](#scalar.covenant.v1beta1.Covenant) | repeated |  |
| `groups` | [CovenantGroup](#scalar.covenant.v1beta1.CovenantGroup) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/covenant/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/covenant/v1beta1/query.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

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
| `name` | [string](#string) |  |  |
| `supports_foreign_assets` | [bool](#bool) |  |  |
| `key_type` | [axelar.tss.exported.v1beta1.KeyType](#axelar.tss.exported.v1beta1.KeyType) |  |  |
| `module` | [string](#string) |  |  |






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



<a name="scalar/evm/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/evm/v1beta1/types.proto



<a name="scalar.evm.v1beta1.Asset"></a>

### Asset



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.BurnerInfo"></a>

### BurnerInfo
BurnerInfo describes information required to burn token at an burner address
that is deposited by an user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `burner_address` | [bytes](#bytes) |  |  |
| `token_address` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |
| `salt` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.Command"></a>

### Command



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [bytes](#bytes) |  |  |
| `command` | [string](#string) |  | **Deprecated.**  |
| `params` | [bytes](#bytes) |  |  |
| `key_id` | [string](#string) |  |  |
| `max_gas_cost` | [uint32](#uint32) |  |  |
| `type` | [CommandType](#scalar.evm.v1beta1.CommandType) |  |  |






<a name="scalar.evm.v1beta1.CommandBatchMetadata"></a>

### CommandBatchMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [bytes](#bytes) |  |  |
| `command_ids` | [bytes](#bytes) | repeated |  |
| `data` | [bytes](#bytes) |  |  |
| `sig_hash` | [bytes](#bytes) |  |  |
| `status` | [BatchedCommandsStatus](#scalar.evm.v1beta1.BatchedCommandsStatus) |  |  |
| `key_id` | [string](#string) |  |  |
| `prev_batched_commands_id` | [bytes](#bytes) |  |  |
| `signature` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="scalar.evm.v1beta1.ERC20Deposit"></a>

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






<a name="scalar.evm.v1beta1.ERC20TokenMetadata"></a>

### ERC20TokenMetadata
ERC20TokenMetadata describes information about an ERC20 token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset` | [string](#string) |  |  |
| `chain_id` | [bytes](#bytes) |  |  |
| `details` | [TokenDetails](#scalar.evm.v1beta1.TokenDetails) |  |  |
| `token_address` | [string](#string) |  |  |
| `tx_hash` | [string](#string) |  |  |
| `status` | [Status](#scalar.evm.v1beta1.Status) |  |  |
| `is_external` | [bool](#bool) |  |  |
| `burner_code` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.Event"></a>

### Event



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `index` | [uint64](#uint64) |  |  |
| `status` | [Event.Status](#scalar.evm.v1beta1.Event.Status) |  |  |
| `token_sent` | [EventTokenSent](#scalar.evm.v1beta1.EventTokenSent) |  |  |
| `contract_call` | [EventContractCall](#scalar.evm.v1beta1.EventContractCall) |  |  |
| `contract_call_with_token` | [EventContractCallWithToken](#scalar.evm.v1beta1.EventContractCallWithToken) |  |  |
| `transfer` | [EventTransfer](#scalar.evm.v1beta1.EventTransfer) |  |  |
| `token_deployed` | [EventTokenDeployed](#scalar.evm.v1beta1.EventTokenDeployed) |  |  |
| `multisig_ownership_transferred` | [EventMultisigOwnershipTransferred](#scalar.evm.v1beta1.EventMultisigOwnershipTransferred) |  | **Deprecated.**  |
| `multisig_operatorship_transferred` | [EventMultisigOperatorshipTransferred](#scalar.evm.v1beta1.EventMultisigOperatorshipTransferred) |  |  |






<a name="scalar.evm.v1beta1.EventContractCall"></a>

### EventContractCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.EventContractCallWithToken"></a>

### EventContractCallWithToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `contract_address` | [string](#string) |  |  |
| `payload_hash` | [bytes](#bytes) |  |  |
| `symbol` | [string](#string) |  |  |
| `amount` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.EventMultisigOperatorshipTransferred"></a>

### EventMultisigOperatorshipTransferred



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `new_operators` | [bytes](#bytes) | repeated |  |
| `new_threshold` | [bytes](#bytes) |  |  |
| `new_weights` | [bytes](#bytes) | repeated |  |






<a name="scalar.evm.v1beta1.EventMultisigOwnershipTransferred"></a>

### EventMultisigOwnershipTransferred



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pre_owners` | [bytes](#bytes) | repeated |  |
| `prev_threshold` | [bytes](#bytes) |  |  |
| `new_owners` | [bytes](#bytes) | repeated |  |
| `new_threshold` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.EventTokenDeployed"></a>

### EventTokenDeployed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `symbol` | [string](#string) |  |  |
| `token_address` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.EventTokenSent"></a>

### EventTokenSent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `destination_address` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `amount` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.EventTransfer"></a>

### EventTransfer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `to` | [bytes](#bytes) |  |  |
| `amount` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.Gateway"></a>

### Gateway



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.NetworkInfo"></a>

### NetworkInfo
NetworkInfo describes information about a network


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `id` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.PollMetadata"></a>

### PollMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.SigMetadata"></a>

### SigMetadata
SigMetadata stores necessary information for external apps to map signature
results to evm relay transaction types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [SigType](#scalar.evm.v1beta1.SigType) |  |  |
| `chain` | [string](#string) |  |  |
| `command_batch_id` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.TokenDetails"></a>

### TokenDetails



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_name` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `decimals` | [uint32](#uint32) |  |  |
| `capacity` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.TransactionMetadata"></a>

### TransactionMetadata



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `raw_tx` | [bytes](#bytes) |  |  |
| `pub_key` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.TransferKey"></a>

### TransferKey
TransferKey contains information for a transfer operatorship


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `next_key_id` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.VoteEvents"></a>

### VoteEvents



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `events` | [Event](#scalar.evm.v1beta1.Event) | repeated |  |





 <!-- end messages -->


<a name="scalar.evm.v1beta1.BatchedCommandsStatus"></a>

### BatchedCommandsStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| BATCHED_COMMANDS_STATUS_UNSPECIFIED | 0 |  |
| BATCHED_COMMANDS_STATUS_SIGNING | 1 |  |
| BATCHED_COMMANDS_STATUS_ABORTED | 2 |  |
| BATCHED_COMMANDS_STATUS_SIGNED | 3 |  |



<a name="scalar.evm.v1beta1.CommandType"></a>

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



<a name="scalar.evm.v1beta1.DepositStatus"></a>

### DepositStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| DEPOSIT_STATUS_UNSPECIFIED | 0 |  |
| DEPOSIT_STATUS_PENDING | 1 |  |
| DEPOSIT_STATUS_CONFIRMED | 2 |  |
| DEPOSIT_STATUS_BURNED | 3 |  |



<a name="scalar.evm.v1beta1.Event.Status"></a>

### Event.Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 |  |
| STATUS_CONFIRMED | 1 |  |
| STATUS_COMPLETED | 2 |  |
| STATUS_FAILED | 3 |  |



<a name="scalar.evm.v1beta1.SigType"></a>

### SigType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SIG_TYPE_UNSPECIFIED | 0 |  |
| SIG_TYPE_TX | 1 |  |
| SIG_TYPE_COMMAND | 2 |  |



<a name="scalar.evm.v1beta1.Status"></a>

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



<a name="scalar/evm/v1beta1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/evm/v1beta1/events.proto



<a name="scalar.evm.v1beta1.BurnCommand"></a>

### BurnCommand



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `command_id` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `deposit_address` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.ChainAdded"></a>

### ChainAdded



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.CommandBatchAborted"></a>

### CommandBatchAborted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `command_batch_id` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.CommandBatchSigned"></a>

### CommandBatchSigned



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `command_batch_id` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.ConfirmDepositStarted"></a>

### ConfirmDepositStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `deposit_address` | [bytes](#bytes) |  |  |
| `token_address` | [bytes](#bytes) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [axelar.vote.exported.v1beta1.PollParticipants](#axelar.vote.exported.v1beta1.PollParticipants) |  |  |
| `asset` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.ConfirmGatewayTxStarted"></a>

### ConfirmGatewayTxStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `gateway_address` | [bytes](#bytes) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [axelar.vote.exported.v1beta1.PollParticipants](#axelar.vote.exported.v1beta1.PollParticipants) |  |  |






<a name="scalar.evm.v1beta1.ConfirmGatewayTxsStarted"></a>

### ConfirmGatewayTxsStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poll_mappings` | [PollMapping](#scalar.evm.v1beta1.PollMapping) | repeated |  |
| `chain` | [string](#string) |  |  |
| `gateway_address` | [bytes](#bytes) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [bytes](#bytes) | repeated |  |






<a name="scalar.evm.v1beta1.ConfirmKeyTransferStarted"></a>

### ConfirmKeyTransferStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `gateway_address` | [bytes](#bytes) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [axelar.vote.exported.v1beta1.PollParticipants](#axelar.vote.exported.v1beta1.PollParticipants) |  |  |






<a name="scalar.evm.v1beta1.ConfirmTokenStarted"></a>

### ConfirmTokenStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `gateway_address` | [bytes](#bytes) |  |  |
| `token_address` | [bytes](#bytes) |  |  |
| `token_details` | [TokenDetails](#scalar.evm.v1beta1.TokenDetails) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [axelar.vote.exported.v1beta1.PollParticipants](#axelar.vote.exported.v1beta1.PollParticipants) |  |  |






<a name="scalar.evm.v1beta1.ContractCallApproved"></a>

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






<a name="scalar.evm.v1beta1.ContractCallFailed"></a>

### ContractCallFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `msg_id` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.ContractCallWithMintApproved"></a>

### ContractCallWithMintApproved



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






<a name="scalar.evm.v1beta1.EVMEventCompleted"></a>

### EVMEventCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.EVMEventConfirmed"></a>

### EVMEventConfirmed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.EVMEventFailed"></a>

### EVMEventFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.EVMEventRetryFailed"></a>

### EVMEventRetryFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.MintCommand"></a>

### MintCommand



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `transfer_id` | [uint64](#uint64) |  |  |
| `command_id` | [bytes](#bytes) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `destination_address` | [string](#string) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="scalar.evm.v1beta1.NoEventsConfirmed"></a>

### NoEventsConfirmed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.evm.v1beta1.PollCompleted"></a>

### PollCompleted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.evm.v1beta1.PollExpired"></a>

### PollExpired



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.evm.v1beta1.PollFailed"></a>

### PollFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.evm.v1beta1.PollMapping"></a>

### PollMapping



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `poll_id` | [uint64](#uint64) |  |  |






<a name="scalar.evm.v1beta1.TokenSent"></a>

### TokenSent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |
| `transfer_id` | [uint64](#uint64) |  |  |
| `sender` | [string](#string) |  |  |
| `destination_chain` | [string](#string) |  |  |
| `destination_address` | [string](#string) |  |  |
| `asset` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/evm/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/evm/v1beta1/params.proto



<a name="scalar.evm.v1beta1.Params"></a>

### Params
Params is the parameter set for this module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `network` | [string](#string) |  |  |
| `token_code` | [bytes](#bytes) |  |  |
| `burnable` | [bytes](#bytes) |  |  |
| `revote_locking_period` | [int64](#int64) |  |  |
| `networks` | [NetworkInfo](#scalar.evm.v1beta1.NetworkInfo) | repeated |  |
| `voting_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
| `min_voter_count` | [int64](#int64) |  |  |
| `commands_gas_limit` | [uint32](#uint32) |  |  |
| `voting_grace_period` | [int64](#int64) |  |  |
| `end_blocker_limit` | [int64](#int64) |  |  |
| `transfer_limit` | [uint64](#uint64) |  |  |






<a name="scalar.evm.v1beta1.PendingChain"></a>

### PendingChain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.evm.v1beta1.Params) |  |  |
| `chain` | [scalar.nexus.exported.v1beta1.Chain](#scalar.nexus.exported.v1beta1.Chain) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/evm/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/evm/v1beta1/genesis.proto



<a name="scalar.evm.v1beta1.GenesisState"></a>

### GenesisState
GenesisState represents the genesis state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chains` | [GenesisState.Chain](#scalar.evm.v1beta1.GenesisState.Chain) | repeated |  |






<a name="scalar.evm.v1beta1.GenesisState.Chain"></a>

### GenesisState.Chain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.evm.v1beta1.Params) |  |  |
| `burner_infos` | [BurnerInfo](#scalar.evm.v1beta1.BurnerInfo) | repeated |  |
| `command_queue` | [axelar.utils.v1beta1.QueueState](#axelar.utils.v1beta1.QueueState) |  |  |
| `confirmed_deposits` | [ERC20Deposit](#scalar.evm.v1beta1.ERC20Deposit) | repeated |  |
| `burned_deposits` | [ERC20Deposit](#scalar.evm.v1beta1.ERC20Deposit) | repeated |  |
| `command_batches` | [CommandBatchMetadata](#scalar.evm.v1beta1.CommandBatchMetadata) | repeated |  |
| `gateway` | [Gateway](#scalar.evm.v1beta1.Gateway) |  |  |
| `tokens` | [ERC20TokenMetadata](#scalar.evm.v1beta1.ERC20TokenMetadata) | repeated |  |
| `events` | [Event](#scalar.evm.v1beta1.Event) | repeated |  |
| `confirmed_event_queue` | [axelar.utils.v1beta1.QueueState](#axelar.utils.v1beta1.QueueState) |  |  |
| `legacy_confirmed_deposits` | [ERC20Deposit](#scalar.evm.v1beta1.ERC20Deposit) | repeated |  |
| `legacy_burned_deposits` | [ERC20Deposit](#scalar.evm.v1beta1.ERC20Deposit) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/evm/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/evm/v1beta1/query.proto



<a name="scalar.evm.v1beta1.BatchedCommandsRequest"></a>

### BatchedCommandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `id` | [string](#string) |  | id defines an optional id for the commandsbatch. If not specified the latest will be returned |






<a name="scalar.evm.v1beta1.BatchedCommandsResponse"></a>

### BatchedCommandsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `data` | [string](#string) |  |  |
| `status` | [BatchedCommandsStatus](#scalar.evm.v1beta1.BatchedCommandsStatus) |  |  |
| `key_id` | [string](#string) |  |  |
| `execute_data` | [string](#string) |  |  |
| `prev_batched_commands_id` | [string](#string) |  |  |
| `command_ids` | [string](#string) | repeated |  |
| `proof` | [Proof](#scalar.evm.v1beta1.Proof) |  |  |






<a name="scalar.evm.v1beta1.BurnerInfoRequest"></a>

### BurnerInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.BurnerInfoResponse"></a>

### BurnerInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `burner_info` | [BurnerInfo](#scalar.evm.v1beta1.BurnerInfo) |  |  |






<a name="scalar.evm.v1beta1.BytecodeRequest"></a>

### BytecodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `contract` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.BytecodeResponse"></a>

### BytecodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bytecode` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.ChainsRequest"></a>

### ChainsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [ChainStatus](#scalar.evm.v1beta1.ChainStatus) |  |  |






<a name="scalar.evm.v1beta1.ChainsResponse"></a>

### ChainsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chains` | [string](#string) | repeated |  |






<a name="scalar.evm.v1beta1.CommandRequest"></a>

### CommandRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `id` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.CommandResponse"></a>

### CommandResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `params` | [CommandResponse.ParamsEntry](#scalar.evm.v1beta1.CommandResponse.ParamsEntry) | repeated |  |
| `key_id` | [string](#string) |  |  |
| `max_gas_cost` | [uint32](#uint32) |  |  |






<a name="scalar.evm.v1beta1.CommandResponse.ParamsEntry"></a>

### CommandResponse.ParamsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.ConfirmationHeightRequest"></a>

### ConfirmationHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.ConfirmationHeightResponse"></a>

### ConfirmationHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [uint64](#uint64) |  |  |






<a name="scalar.evm.v1beta1.DepositQueryParams"></a>

### DepositQueryParams
DepositQueryParams describe the parameters used to query for an EVM
deposit address


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.DepositStateRequest"></a>

### DepositStateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `params` | [QueryDepositStateParams](#scalar.evm.v1beta1.QueryDepositStateParams) |  |  |






<a name="scalar.evm.v1beta1.DepositStateResponse"></a>

### DepositStateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [DepositStatus](#scalar.evm.v1beta1.DepositStatus) |  |  |






<a name="scalar.evm.v1beta1.ERC20TokensRequest"></a>

### ERC20TokensRequest
ERC20TokensRequest describes the chain for which the type of ERC20 tokens are
requested.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `type` | [TokenType](#scalar.evm.v1beta1.TokenType) |  |  |






<a name="scalar.evm.v1beta1.ERC20TokensResponse"></a>

### ERC20TokensResponse
ERC20TokensResponse describes the asset and symbol for all
ERC20 tokens requested for a chain


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tokens` | [ERC20TokensResponse.Token](#scalar.evm.v1beta1.ERC20TokensResponse.Token) | repeated |  |






<a name="scalar.evm.v1beta1.ERC20TokensResponse.Token"></a>

### ERC20TokensResponse.Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.EventRequest"></a>

### EventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.EventResponse"></a>

### EventResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `event` | [Event](#scalar.evm.v1beta1.Event) |  |  |






<a name="scalar.evm.v1beta1.GatewayAddressRequest"></a>

### GatewayAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.GatewayAddressResponse"></a>

### GatewayAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.KeyAddressRequest"></a>

### KeyAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.KeyAddressResponse"></a>

### KeyAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key_id` | [string](#string) |  |  |
| `addresses` | [KeyAddressResponse.WeightedAddress](#scalar.evm.v1beta1.KeyAddressResponse.WeightedAddress) | repeated |  |
| `threshold` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.KeyAddressResponse.WeightedAddress"></a>

### KeyAddressResponse.WeightedAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `weight` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.ParamsRequest"></a>

### ParamsRequest
ParamsRequest represents a message that queries the params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.ParamsResponse"></a>

### ParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#scalar.evm.v1beta1.Params) |  |  |






<a name="scalar.evm.v1beta1.PendingCommandsRequest"></a>

### PendingCommandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.PendingCommandsResponse"></a>

### PendingCommandsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commands` | [QueryCommandResponse](#scalar.evm.v1beta1.QueryCommandResponse) | repeated |  |






<a name="scalar.evm.v1beta1.Proof"></a>

### Proof



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [string](#string) | repeated |  |
| `weights` | [string](#string) | repeated |  |
| `threshold` | [string](#string) |  |  |
| `signatures` | [string](#string) | repeated |  |






<a name="scalar.evm.v1beta1.QueryBurnerAddressResponse"></a>

### QueryBurnerAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.QueryCommandResponse"></a>

### QueryCommandResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |
| `params` | [QueryCommandResponse.ParamsEntry](#scalar.evm.v1beta1.QueryCommandResponse.ParamsEntry) | repeated |  |
| `key_id` | [string](#string) |  |  |
| `max_gas_cost` | [uint32](#uint32) |  |  |






<a name="scalar.evm.v1beta1.QueryCommandResponse.ParamsEntry"></a>

### QueryCommandResponse.ParamsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.QueryDepositStateParams"></a>

### QueryDepositStateParams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_id` | [bytes](#bytes) |  |  |
| `burner_address` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.QueryTokenAddressResponse"></a>

### QueryTokenAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `confirmed` | [bool](#bool) |  |  |






<a name="scalar.evm.v1beta1.TokenInfoRequest"></a>

### TokenInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [string](#string) |  |  |
| `asset` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `address` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.TokenInfoResponse"></a>

### TokenInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset` | [string](#string) |  |  |
| `details` | [TokenDetails](#scalar.evm.v1beta1.TokenDetails) |  |  |
| `address` | [string](#string) |  |  |
| `confirmed` | [bool](#bool) |  |  |
| `is_external` | [bool](#bool) |  |  |
| `burner_code_hash` | [string](#string) |  |  |





 <!-- end messages -->


<a name="scalar.evm.v1beta1.ChainStatus"></a>

### ChainStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHAIN_STATUS_UNSPECIFIED | 0 |  |
| CHAIN_STATUS_ACTIVATED | 1 |  |
| CHAIN_STATUS_DEACTIVATED | 2 |  |



<a name="scalar.evm.v1beta1.TokenType"></a>

### TokenType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TOKEN_TYPE_UNSPECIFIED | 0 |  |
| TOKEN_TYPE_INTERNAL | 1 |  |
| TOKEN_TYPE_EXTERNAL | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/evm/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/evm/v1beta1/tx.proto



<a name="scalar.evm.v1beta1.AddChainRequest"></a>

### AddChainRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `name` | [string](#string) |  |  |
| `key_type` | [axelar.tss.exported.v1beta1.KeyType](#axelar.tss.exported.v1beta1.KeyType) |  | **Deprecated.**  |
| `params` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.AddChainResponse"></a>

### AddChainResponse







<a name="scalar.evm.v1beta1.ConfirmDepositRequest"></a>

### ConfirmDepositRequest
MsgConfirmDeposit represents an erc20 deposit confirmation message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `amount` | [bytes](#bytes) |  | **Deprecated.**  |
| `burner_address` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.ConfirmDepositResponse"></a>

### ConfirmDepositResponse







<a name="scalar.evm.v1beta1.ConfirmGatewayTxRequest"></a>

### ConfirmGatewayTxRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.ConfirmGatewayTxResponse"></a>

### ConfirmGatewayTxResponse







<a name="scalar.evm.v1beta1.ConfirmGatewayTxsRequest"></a>

### ConfirmGatewayTxsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_ids` | [bytes](#bytes) | repeated |  |






<a name="scalar.evm.v1beta1.ConfirmGatewayTxsResponse"></a>

### ConfirmGatewayTxsResponse







<a name="scalar.evm.v1beta1.ConfirmTokenRequest"></a>

### ConfirmTokenRequest
MsgConfirmToken represents a token deploy confirmation message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |
| `asset` | [Asset](#scalar.evm.v1beta1.Asset) |  |  |






<a name="scalar.evm.v1beta1.ConfirmTokenResponse"></a>

### ConfirmTokenResponse







<a name="scalar.evm.v1beta1.ConfirmTransferKeyRequest"></a>

### ConfirmTransferKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `tx_id` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.ConfirmTransferKeyResponse"></a>

### ConfirmTransferKeyResponse







<a name="scalar.evm.v1beta1.CreateBurnTokensRequest"></a>

### CreateBurnTokensRequest
CreateBurnTokensRequest represents the message to create commands to burn
tokens with AxelarGateway


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.CreateBurnTokensResponse"></a>

### CreateBurnTokensResponse







<a name="scalar.evm.v1beta1.CreateDeployTokenRequest"></a>

### CreateDeployTokenRequest
CreateDeployTokenRequest represents the message to create a deploy token
command for AxelarGateway


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `asset` | [Asset](#scalar.evm.v1beta1.Asset) |  |  |
| `token_details` | [TokenDetails](#scalar.evm.v1beta1.TokenDetails) |  |  |
| `address` | [bytes](#bytes) |  |  |
| `daily_mint_limit` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.CreateDeployTokenResponse"></a>

### CreateDeployTokenResponse







<a name="scalar.evm.v1beta1.CreatePendingTransfersRequest"></a>

### CreatePendingTransfersRequest
CreatePendingTransfersRequest represents a message to trigger the creation of
commands handling all pending transfers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.CreatePendingTransfersResponse"></a>

### CreatePendingTransfersResponse







<a name="scalar.evm.v1beta1.CreateTransferOperatorshipRequest"></a>

### CreateTransferOperatorshipRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.CreateTransferOperatorshipResponse"></a>

### CreateTransferOperatorshipResponse







<a name="scalar.evm.v1beta1.CreateTransferOwnershipRequest"></a>

### CreateTransferOwnershipRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `key_id` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.CreateTransferOwnershipResponse"></a>

### CreateTransferOwnershipResponse







<a name="scalar.evm.v1beta1.LinkRequest"></a>

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






<a name="scalar.evm.v1beta1.LinkResponse"></a>

### LinkResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit_addr` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.RetryFailedEventRequest"></a>

### RetryFailedEventRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `event_id` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.RetryFailedEventResponse"></a>

### RetryFailedEventResponse







<a name="scalar.evm.v1beta1.SetGatewayRequest"></a>

### SetGatewayRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |
| `address` | [bytes](#bytes) |  |  |






<a name="scalar.evm.v1beta1.SetGatewayResponse"></a>

### SetGatewayResponse







<a name="scalar.evm.v1beta1.SignCommandsRequest"></a>

### SignCommandsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  |  |
| `chain` | [string](#string) |  |  |






<a name="scalar.evm.v1beta1.SignCommandsResponse"></a>

### SignCommandsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batched_commands_id` | [bytes](#bytes) |  |  |
| `command_count` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/evm/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/evm/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="scalar.evm.v1beta1.MsgService"></a>

### MsgService
Msg defines the evm Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetGateway` | [SetGatewayRequest](#scalar.evm.v1beta1.SetGatewayRequest) | [SetGatewayResponse](#scalar.evm.v1beta1.SetGatewayResponse) |  | POST|/scalar/evm/set_gateway|
| `ConfirmGatewayTx` | [ConfirmGatewayTxRequest](#scalar.evm.v1beta1.ConfirmGatewayTxRequest) | [ConfirmGatewayTxResponse](#scalar.evm.v1beta1.ConfirmGatewayTxResponse) | Deprecated: use ConfirmGatewayTxs instead | POST|/scalar/evm/confirm_gateway_tx|
| `ConfirmGatewayTxs` | [ConfirmGatewayTxsRequest](#scalar.evm.v1beta1.ConfirmGatewayTxsRequest) | [ConfirmGatewayTxsResponse](#scalar.evm.v1beta1.ConfirmGatewayTxsResponse) |  | POST|/scalar/evm/confirm_gateway_txs|
| `Link` | [LinkRequest](#scalar.evm.v1beta1.LinkRequest) | [LinkResponse](#scalar.evm.v1beta1.LinkResponse) |  | POST|/scalar/evm/link|
| `ConfirmToken` | [ConfirmTokenRequest](#scalar.evm.v1beta1.ConfirmTokenRequest) | [ConfirmTokenResponse](#scalar.evm.v1beta1.ConfirmTokenResponse) |  | POST|/scalar/evm/confirm_token|
| `ConfirmDeposit` | [ConfirmDepositRequest](#scalar.evm.v1beta1.ConfirmDepositRequest) | [ConfirmDepositResponse](#scalar.evm.v1beta1.ConfirmDepositResponse) |  | POST|/scalar/evm/confirm_deposit|
| `ConfirmTransferKey` | [ConfirmTransferKeyRequest](#scalar.evm.v1beta1.ConfirmTransferKeyRequest) | [ConfirmTransferKeyResponse](#scalar.evm.v1beta1.ConfirmTransferKeyResponse) |  | POST|/scalar/evm/confirm_transfer_key|
| `CreateDeployToken` | [CreateDeployTokenRequest](#scalar.evm.v1beta1.CreateDeployTokenRequest) | [CreateDeployTokenResponse](#scalar.evm.v1beta1.CreateDeployTokenResponse) |  | POST|/scalar/evm/create_deploy_token|
| `CreateBurnTokens` | [CreateBurnTokensRequest](#scalar.evm.v1beta1.CreateBurnTokensRequest) | [CreateBurnTokensResponse](#scalar.evm.v1beta1.CreateBurnTokensResponse) |  | POST|/scalar/evm/create_burn_tokens|
| `CreatePendingTransfers` | [CreatePendingTransfersRequest](#scalar.evm.v1beta1.CreatePendingTransfersRequest) | [CreatePendingTransfersResponse](#scalar.evm.v1beta1.CreatePendingTransfersResponse) |  | POST|/scalar/evm/create_pending_transfers|
| `CreateTransferOperatorship` | [CreateTransferOperatorshipRequest](#scalar.evm.v1beta1.CreateTransferOperatorshipRequest) | [CreateTransferOperatorshipResponse](#scalar.evm.v1beta1.CreateTransferOperatorshipResponse) |  | POST|/scalar/evm/create_transfer_operatorship|
| `SignCommands` | [SignCommandsRequest](#scalar.evm.v1beta1.SignCommandsRequest) | [SignCommandsResponse](#scalar.evm.v1beta1.SignCommandsResponse) |  | POST|/scalar/evm/sign_commands|
| `AddChain` | [AddChainRequest](#scalar.evm.v1beta1.AddChainRequest) | [AddChainResponse](#scalar.evm.v1beta1.AddChainResponse) |  | POST|/scalar/evm/add_chain|
| `RetryFailedEvent` | [RetryFailedEventRequest](#scalar.evm.v1beta1.RetryFailedEventRequest) | [RetryFailedEventResponse](#scalar.evm.v1beta1.RetryFailedEventResponse) |  | POST|/scalar/evm/retry-failed-event|


<a name="scalar.evm.v1beta1.QueryService"></a>

### QueryService
QueryService defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `BatchedCommands` | [BatchedCommandsRequest](#scalar.evm.v1beta1.BatchedCommandsRequest) | [BatchedCommandsResponse](#scalar.evm.v1beta1.BatchedCommandsResponse) | BatchedCommands queries the batched commands for a specified chain and BatchedCommandsID if no BatchedCommandsID is specified, then it returns the latest batched commands | GET|/scalar/evm/v1beta1/batched_commands/{chain}/{id}|
| `BurnerInfo` | [BurnerInfoRequest](#scalar.evm.v1beta1.BurnerInfoRequest) | [BurnerInfoResponse](#scalar.evm.v1beta1.BurnerInfoResponse) | BurnerInfo queries the burner info for the specified address | GET|/scalar/evm/v1beta1/burner_info|
| `ConfirmationHeight` | [ConfirmationHeightRequest](#scalar.evm.v1beta1.ConfirmationHeightRequest) | [ConfirmationHeightResponse](#scalar.evm.v1beta1.ConfirmationHeightResponse) | ConfirmationHeight queries the confirmation height for the specified chain | GET|/scalar/evm/v1beta1/confirmation_height/{chain}|
| `DepositState` | [DepositStateRequest](#scalar.evm.v1beta1.DepositStateRequest) | [DepositStateResponse](#scalar.evm.v1beta1.DepositStateResponse) | DepositState queries the state of the specified deposit | GET|/scalar/evm/v1beta1/deposit_state|
| `PendingCommands` | [PendingCommandsRequest](#scalar.evm.v1beta1.PendingCommandsRequest) | [PendingCommandsResponse](#scalar.evm.v1beta1.PendingCommandsResponse) | PendingCommands queries the pending commands for the specified chain | GET|/scalar/evm/v1beta1/pending_commands/{chain}|
| `Chains` | [ChainsRequest](#scalar.evm.v1beta1.ChainsRequest) | [ChainsResponse](#scalar.evm.v1beta1.ChainsResponse) | Chains queries the available evm chains | GET|/scalar/evm/v1beta1/chains|
| `Command` | [CommandRequest](#scalar.evm.v1beta1.CommandRequest) | [CommandResponse](#scalar.evm.v1beta1.CommandResponse) | Command queries the command of a chain provided the command id | GET|/scalar/evm/v1beta1/command_request|
| `KeyAddress` | [KeyAddressRequest](#scalar.evm.v1beta1.KeyAddressRequest) | [KeyAddressResponse](#scalar.evm.v1beta1.KeyAddressResponse) | KeyAddress queries the address of key of a chain | GET|/scalar/evm/v1beta1/key_address/{chain}|
| `GatewayAddress` | [GatewayAddressRequest](#scalar.evm.v1beta1.GatewayAddressRequest) | [GatewayAddressResponse](#scalar.evm.v1beta1.GatewayAddressResponse) | GatewayAddress queries the address of axelar gateway at the specified chain | GET|/scalar/evm/v1beta1/gateway_address/{chain}|
| `Bytecode` | [BytecodeRequest](#scalar.evm.v1beta1.BytecodeRequest) | [BytecodeResponse](#scalar.evm.v1beta1.BytecodeResponse) | Bytecode queries the bytecode of a specified gateway at the specified chain | GET|/scalar/evm/v1beta1/bytecode/{chain}/{contract}|
| `Event` | [EventRequest](#scalar.evm.v1beta1.EventRequest) | [EventResponse](#scalar.evm.v1beta1.EventResponse) | Event queries an event at the specified chain | GET|/scalar/evm/v1beta1/event/{chain}/{event_id}|
| `ERC20Tokens` | [ERC20TokensRequest](#scalar.evm.v1beta1.ERC20TokensRequest) | [ERC20TokensResponse](#scalar.evm.v1beta1.ERC20TokensResponse) | ERC20Tokens queries the ERC20 tokens registered for a chain | GET|/scalar/evm/v1beta1/erc20_tokens/{chain}|
| `TokenInfo` | [TokenInfoRequest](#scalar.evm.v1beta1.TokenInfoRequest) | [TokenInfoResponse](#scalar.evm.v1beta1.TokenInfoResponse) | TokenInfo queries the token info for a registered ERC20 Token | GET|/scalar/evm/v1beta1/token_info/{chain}|
| `Params` | [ParamsRequest](#scalar.evm.v1beta1.ParamsRequest) | [ParamsResponse](#scalar.evm.v1beta1.ParamsResponse) |  | GET|/scalar/evm/v1beta1/params/{chain}|

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



<a name="scalar/multisig/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/multisig/v1beta1/params.proto



<a name="scalar.multisig.v1beta1.Params"></a>

### Params
Params represent the genesis parameters for the module


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keygen_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
| `signing_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
| `keygen_timeout` | [int64](#int64) |  |  |
| `keygen_grace_period` | [int64](#int64) |  |  |
| `signing_timeout` | [int64](#int64) |  |  |
| `signing_grace_period` | [int64](#int64) |  |  |
| `active_epoch_count` | [uint64](#uint64) |  |  |





 <!-- end messages -->

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
| `snapshot` | [axelar.snapshot.exported.v1beta1.Snapshot](#axelar.snapshot.exported.v1beta1.Snapshot) |  |  |
| `pub_keys` | [Key.PubKeysEntry](#scalar.multisig.v1beta1.Key.PubKeysEntry) | repeated |  |
| `signing_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
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
| `keygen_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
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






<a name="scalar.nexus.v1beta1.MessageFailed"></a>

### MessageFailed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |






<a name="scalar.nexus.v1beta1.MessageProcessing"></a>

### MessageProcessing



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |






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
| `chain_activation_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
| `chain_maintainer_missing_vote_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
| `chain_maintainer_incorrect_vote_threshold` | [axelar.utils.v1beta1.Threshold](#axelar.utils.v1beta1.Threshold) |  |  |
| `chain_maintainer_check_window` | [int32](#int32) |  |  |
| `gateway` | [bytes](#bytes) |  |  |
| `end_blocker_limit` | [uint64](#uint64) |  |  |





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
| `missing_votes` | [axelar.utils.v1beta1.Bitmap](#axelar.utils.v1beta1.Bitmap) |  |  |
| `incorrect_votes` | [axelar.utils.v1beta1.Bitmap](#axelar.utils.v1beta1.Bitmap) |  |  |
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



<a name="scalar/protocol/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/types.proto



<a name="scalar.protocol.v1beta1.Protocol"></a>

### Protocol



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `covenant_group` | [string](#string) |  |  |
| `tokens` | [scalar.evm.v1beta1.ERC20TokenMetadata](#scalar.evm.v1beta1.ERC20TokenMetadata) | repeated |  |





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



<a name="scalar/protocol/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/query.proto



<a name="scalar.protocol.v1beta1.ProtocolRequest"></a>

### ProtocolRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [ProtocolStatus](#scalar.protocol.v1beta1.ProtocolStatus) |  |  |






<a name="scalar.protocol.v1beta1.ProtocolResponse"></a>

### ProtocolResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocols` | [Protocol](#scalar.protocol.v1beta1.Protocol) | repeated |  |





 <!-- end messages -->


<a name="scalar.protocol.v1beta1.ProtocolStatus"></a>

### ProtocolStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| PROTOCOL_STATUS_UNSPECIFIED | 0 |  |
| PROTOCOL_STATUS_ACTIVATED | 1 |  |
| PROTOCOL_STATUS_DEACTIVATED | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/protocol/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/service.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="scalar/protocol/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/tx.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

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
| `transfer_queue` | [axelar.utils.v1beta1.QueueState](#axelar.utils.v1beta1.QueueState) |  |  |
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

