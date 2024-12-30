package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/client/rpc/cmd/processor"
	"github.com/scalarorg/scalar-core/client/rpc/config"
	"github.com/scalarorg/scalar-core/client/rpc/cosmos"
	"github.com/scalarorg/scalar-core/client/rpc/jobs"
)

const (
	AccountAddressPrefix = "scalar"
	BaseAsset            = "ascal"
)

var (
	AccountPubKeyPrefix    = AccountAddressPrefix + types.PrefixPublic
	ValidatorAddressPrefix = AccountAddressPrefix + types.PrefixValidator + types.PrefixOperator
	ValidatorPubKeyPrefix  = AccountAddressPrefix + types.PrefixValidator + types.PrefixOperator + types.PrefixPublic
	ConsNodeAddressPrefix  = AccountAddressPrefix + types.PrefixValidator + types.PrefixConsensus
	ConsNodePubKeyPrefix   = AccountAddressPrefix + types.PrefixValidator + types.PrefixConsensus + types.PrefixPublic
)

var (
	DestCallApprovedEvent = cosmos.EventQuery{
		TmEvent:   "NewBlock",
		Module:    "scalar.chains",
		Version:   "v1beta1",
		Event:     "DestCallApproved",
		Attribute: "event_id",
		Operator:  "EXISTS",
	}
)

func setCosmosAccountPrefix() {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
}

func initializeApp() error {
	// Configure zerolog with console writer and colors
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.DateTime,
		NoColor:    false,
	}

	// Set global logger
	log.Logger = log.Output(output)

	setCosmosAccountPrefix()
	return config.ReadConfig("config/example.json")
}

func setupNetworkClient() (*cosmos.NetworkClient, types.AccAddress, error) {
	clientCtx, err := cosmos.CreateClientContextWithOptions(
		cosmos.WithRpcClientCtx(config.GlobalConfig.RPCUrl),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create client context: %w", err)
	}

	queryClient := cosmos.NewQueryClient(clientCtx)
	privKey, addr, err := cosmos.CreateAccountFromMnemonic(config.GlobalConfig.Mnemonic, "")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create account: %w", err)
	}

	networkClient, err := cosmos.NewNetworkClientWithOptions(
		queryClient,
		config.GlobalTxConfig,
		cosmos.WithRpcClient(clientCtx.Client),
		cosmos.WithQueryClient(queryClient),
		cosmos.WithAccount(privKey, addr),
		cosmos.WithTxConfig(config.GlobalTxConfig),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create network client: %w", err)
	}

	return networkClient, addr, nil
}

func main() {
	if err := initializeApp(); err != nil {
		panic(err)
	}

	networkClient, addr, err := setupNetworkClient()
	if err != nil {
		panic(err)
	}

	proc := processor.NewProcessor(networkClient)

	fmt.Println("address", addr.String())

	// Start the network client
	if err := networkClient.Start(); err != nil {
		panic(fmt.Errorf("failed to start network client: %w", err))
	}

	subscribedJobs := []*jobs.EventJob{
		jobs.NewEventJob(
			"dest_call_approved_event",
			DestCallApprovedEvent,
			networkClient,
		),
	}

	var wg sync.WaitGroup
	for _, job := range subscribedJobs {
		wg.Add(1)
		go func(j *jobs.EventJob) {
			defer wg.Done()
			jobs.RunJob(j, context.Background(), proc.ProcessDestCallApproved)
		}(job)
	}

	wg.Wait()
}
