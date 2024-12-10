<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [scalar/btc/v1beta1/types.proto](#scalar/btc/v1beta1/types.proto)
    - [Command](#scalar.btc.v1beta1.Command)
    - [CommandBatchMetadata](#scalar.btc.v1beta1.CommandBatchMetadata)
    - [PollMapping](#scalar.btc.v1beta1.PollMapping)
    - [PollMetadata](#scalar.btc.v1beta1.PollMetadata)
    - [Proof](#scalar.btc.v1beta1.Proof)
    - [StakingTx](#scalar.btc.v1beta1.StakingTx)
    - [StakingTxMetadata](#scalar.btc.v1beta1.StakingTxMetadata)
  
    - [BatchedCommandsStatus](#scalar.btc.v1beta1.BatchedCommandsStatus)
    - [CommandType](#scalar.btc.v1beta1.CommandType)
    - [NetworkKind](#scalar.btc.v1beta1.NetworkKind)
    - [StakingTxStatus](#scalar.btc.v1beta1.StakingTxStatus)
  
- [scalar/btc/v1beta1/events.proto](#scalar/btc/v1beta1/events.proto)
    - [ConfirmStakingTxsStarted](#scalar.btc.v1beta1.ConfirmStakingTxsStarted)
    - [Event](#scalar.btc.v1beta1.Event)
    - [EventStakingTx](#scalar.btc.v1beta1.EventStakingTx)
    - [VoteEvents](#scalar.btc.v1beta1.VoteEvents)
  
    - [Event.Status](#scalar.btc.v1beta1.Event.Status)
  
- [scalar/btc/v1beta1/params.proto](#scalar/btc/v1beta1/params.proto)
    - [Params](#scalar.btc.v1beta1.Params)
  
- [scalar/btc/v1beta1/genesis.proto](#scalar/btc/v1beta1/genesis.proto)
    - [GenesisState](#scalar.btc.v1beta1.GenesisState)
    - [GenesisState.Chain](#scalar.btc.v1beta1.GenesisState.Chain)
  
- [scalar/btc/v1beta1/tx.proto](#scalar/btc/v1beta1/tx.proto)
    - [ConfirmStakingTxsRequest](#scalar.btc.v1beta1.ConfirmStakingTxsRequest)
    - [ConfirmStakingTxsResponse](#scalar.btc.v1beta1.ConfirmStakingTxsResponse)
  
- [scalar/btc/v1beta1/query.proto](#scalar/btc/v1beta1/query.proto)
    - [BatchedCommandsRequest](#scalar.btc.v1beta1.BatchedCommandsRequest)
    - [BatchedCommandsResponse](#scalar.btc.v1beta1.BatchedCommandsResponse)
  
- [scalar/btc/v1beta1/service.proto](#scalar/btc/v1beta1/service.proto)
    - [MsgService](#scalar.btc.v1beta1.MsgService)
    - [QueryService](#scalar.btc.v1beta1.QueryService)
  
- [scalar/covenant/v1beta1/types.proto](#scalar/covenant/v1beta1/types.proto)
    - [Covenant](#scalar.covenant.v1beta1.Covenant)
    - [CovenantGroup](#scalar.covenant.v1beta1.CovenantGroup)
  
- [scalar/covenant/v1beta1/genesis.proto](#scalar/covenant/v1beta1/genesis.proto)
    - [GenesisState](#scalar.covenant.v1beta1.GenesisState)
  
- [scalar/covenant/v1beta1/query.proto](#scalar/covenant/v1beta1/query.proto)
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



<a name="scalar.btc.v1beta1.ConfirmStakingTxsStarted"></a>

### ConfirmStakingTxsStarted



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `poll_mappings` | [PollMapping](#scalar.btc.v1beta1.PollMapping) | repeated |  |
| `chain` | [string](#string) |  |  |
| `confirmation_height` | [uint64](#uint64) |  |  |
| `participants` | [bytes](#bytes) | repeated |  |






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



<a name="scalar/protocol/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## scalar/protocol/v1beta1/types.proto



<a name="scalar.protocol.v1beta1.Protocol"></a>

### Protocol



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `covenant_group` | [string](#string) |  |  |
| `tokens` | [axelar.evm.v1beta1.ERC20TokenMetadata](#axelar.evm.v1beta1.ERC20TokenMetadata) | repeated |  |





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
| `chain` | [axelar.nexus.exported.v1beta1.Chain](#axelar.nexus.exported.v1beta1.Chain) |  | **Deprecated.** chain was deprecated in v0.27 |
| `addr_prefix` | [string](#string) |  |  |
| `native_assets` | [axelar.nexus.exported.v1beta1.Asset](#axelar.nexus.exported.v1beta1.Asset) | repeated | **Deprecated.** native_assets was deprecated in v0.27 |
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
| `asset` | [axelar.nexus.exported.v1beta1.Asset](#axelar.nexus.exported.v1beta1.Asset) |  |  |
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

