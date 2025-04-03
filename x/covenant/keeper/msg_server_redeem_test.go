package keeper_test

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/scalarorg/scalar-core/testutils/fake"
// 	"github.com/scalarorg/scalar-core/testutils/rand"
// 	"github.com/scalarorg/scalar-core/x/covenant/keeper"
// 	"github.com/scalarorg/scalar-core/x/covenant/types"
// 	"github.com/scalarorg/scalar-core/x/covenant/types/mock"
// 	"github.com/tendermint/tendermint/libs/log"
// 	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
// )

// func setup() (sdk.Context, types.MsgServiceServer) {
// 	ctx := sdk.NewContext(fake.NewMultiStore(), tmproto.Header{Height: rand.PosI64()}, false, log.TestingLogger())

// 	return ctx,
// 		keeper.NewMsgServerImpl(&keeper.MsgServerConstructArgs{
// 			Keeper:      &mock.KeeperMock{},
// 			Snapshotter: &mock.SnapshotterMock{},
// 			Staker:      &mock.StakingKeeperMock{},
// 			Slashing:    &mock.SlashingKeeperMock{},
// 			Multisig:    &mock.MultisigKeeperMock{},
// 			Nexus:       &mock.NexusMock{},
// 			Protocol:    &mock.ProtocolKeeperMock{},
// 			Chains:      &mock.BaseKeeperMock{},
// 			Voter:       &mock.VoterMock{},
// 		})
// }

// func TestHandleMsgConfirmGatewayTxs(t *testing.T) {
// 	validators := slices.Expand(func(int) snapshot.Participant { return snapshot.NewParticipant(rand2.ValAddr(), sdk.OneUint()) }, 10)
// 	txIDs := slices.Expand2(evmTestUtils.RandomHash, int(rand.I64Between(5, 50)))
// 	req := types.NewConfirmGatewayTxsRequest(rand.AccAddr(), nexus.ChainName(rand.Str(5)), txIDs)

// 	var (
// 		ctx         sdk.Context
// 		bk          *mock.BaseKeeperMock
// 		ck          *mock.ChainKeeperMock
// 		s           *mock.SlashingKeeperMock
// 		n           *mock.NexusMock
// 		snapshotter *mock.SnapshotterMock
// 		v           *mock.VoterMock
// 		msgServer   types.MsgServiceServer
// 		pollID      vote.PollID
// 	)

// 	givenMsgServer := Given("an EVM msg server", func() {
// 		ctx = rand2.Context(fake.NewMultiStore())

// 		bk = &mock.BaseKeeperMock{
// 			LoggerFunc:   func(ctx sdk.Context) log.Logger { return ctx.Logger() },
// 			ForChainFunc: func(sdk.Context, nexus.ChainName) (types.ChainKeeper, error) { return nil, fmt.Errorf("unknown chain") },
// 		}
// 		snapshotter = &mock.SnapshotterMock{
// 			CreateSnapshotFunc: func(sdk.Context, []sdk.ValAddress, func(snapshot.ValidatorI) bool, func(consensusPower sdk.Uint) sdk.Uint, utils.Threshold) (snapshot.Snapshot, error) {
// 				return snapshot.NewSnapshot(ctx.BlockTime(), ctx.BlockHeight(), validators, sdk.NewUint(10)), nil
// 			},
// 		}
// 		ck = &mock.ChainKeeperMock{
// 			GetRequiredConfirmationHeightFunc: func(sdk.Context) uint64 { return 10 },
// 			GetParamsFunc:                     func(sdk.Context) types.Params { return types.DefaultParams()[0] },
// 		}
// 		n = &mock.NexusMock{
// 			GetChainMaintainersFunc: func(sdk.Context, nexus.Chain) []sdk.ValAddress {
// 				return slices.Expand2(rand2.ValAddr, 10)
// 			},
// 		}
// 		s = &mock.SlashingKeeperMock{}
// 		v = &mock.VoterMock{}
// 		pollID = vote.PollID(0)

// 		msgServer = keeper.NewMsgServerImpl(bk, n, v, snapshotter, &mock.StakingKeeperMock{}, s, &mock.MultisigKeeperMock{})
// 	})

// 	whenChainIsValid := When("chain is set and activated", func() {
// 		n.GetChainFunc = func(sdk.Context, nexus.ChainName) (nexus.Chain, bool) { return nexus.Chain{}, true }
// 		n.IsChainActivatedFunc = func(sdk.Context, nexus.Chain) bool { return true }
// 		bk.ForChainFunc = func(_ sdk.Context, chain nexus.ChainName) (types.ChainKeeper, error) { return ck, nil }
// 		ck.GetGatewayAddressFunc = func(sdk.Context) (types.Address, bool) { return evmTestUtils.RandomAddress(), true }
// 	})

// 	whenSnapshotIsCreated := When("snapshot is created", func() {
// 		snapshotter.GetProxyFunc = func(sdk.Context, sdk.ValAddress) (sdk.AccAddress, bool) {
// 			return rand2.AccAddr(), true
// 		}
// 		s.IsTombstonedFunc = func(ctx sdk.Context, consAddr sdk.ConsAddress) bool { return false }
// 	})

// 	whenPollsAreInitialized := When("polls are initialized", func() {
// 		v.InitializePollFunc = func(sdk.Context, vote.PollBuilder) (vote.PollID, error) {
// 			pollID += 1
// 			return pollID, nil
// 		}
// 	})

// 	t.Run("confirm gateway txs", func(t *testing.T) {
// 		givenMsgServer.Branch(
// 			whenChainIsValid.
// 				When("failed to create snapshot", func() {
// 					snapshotter.CreateSnapshotFunc = func(sdk.Context, []sdk.ValAddress, func(snapshot.ValidatorI) bool, func(consensusPower sdk.Uint) sdk.Uint, utils.Threshold) (snapshot.Snapshot, error) {
// 						return snapshot.Snapshot{}, fmt.Errorf("failed to create snapshot")
// 					}
// 				}).
// 				Then("should return error", func(t *testing.T) {
// 					_, err := msgServer.ConfirmGatewayTxs(sdk.WrapSDKContext(ctx), req)
// 					assert.ErrorContains(t, err, "failed to create snapshot")
// 				}),
// 			whenChainIsValid.
// 				When2(whenSnapshotIsCreated).
// 				When("failed to initialize polls", func() {
// 					v.InitializePollFunc = func(sdk.Context, vote.PollBuilder) (vote.PollID, error) {
// 						return 0, fmt.Errorf("failed to initialize polls")
// 					}
// 				}).
// 				Then("should return error", func(t *testing.T) {
// 					_, err := msgServer.ConfirmGatewayTxs(sdk.WrapSDKContext(ctx), req)
// 					assert.ErrorContains(t, err, "failed to initialize polls")
// 				}),
// 			whenChainIsValid.
// 				When2(whenSnapshotIsCreated).
// 				When2(whenPollsAreInitialized).
// 				Then("should emit ConfirmGatewayTxsEvent", func(t *testing.T) {
// 					_, err := msgServer.ConfirmGatewayTxs(sdk.WrapSDKContext(ctx), req)
// 					assert.Equal(t, 1, len(ctx.EventManager().Events()))
// 					assert.NoError(t, err)
// 				}),
// 		).Run(t)
// 	})
// }

// func createSignedDeployTx() *evmTypes.Transaction {
// 	generator := rand.PInt64Gen()

// 	nonce := uint64(generator.Next())
// 	gasPrice := big.NewInt(generator.Next())
// 	gasLimit := uint64(generator.Next())
// 	value := big.NewInt(0)
// 	byteCode := rand.Bytes(int(rand.I64Between(1, 10000)))

// 	return sign(evmTypes.NewContractCreation(nonce, value, gasLimit, gasPrice, byteCode))
// }

// func sign(tx *evmTypes.Transaction) *evmTypes.Transaction {
// 	privateKey, err := evmCrypto.GenerateKey()
// 	if err != nil {
// 		panic(err)
// 	}
// 	signedTx, err := evmTypes.SignTx(tx, evmTypes.NewEIP155Signer(networkConf.ChainID), privateKey)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return signedTx
// }

// func createSignedTx() *evmTypes.Transaction {
// 	generator := rand.PInt64Gen()
// 	contractAddr := common.BytesToAddress(rand.Bytes(common.AddressLength))
// 	nonce := uint64(generator.Next())
// 	gasPrice := big.NewInt(rand.PInt64Gen().Next())
// 	gasLimit := uint64(generator.Next())
// 	value := big.NewInt(0)

// 	data := rand.Bytes(int(rand.I64Between(0, 1000)))
// 	return sign(evmTypes.NewTransaction(nonce, contractAddr, value, gasLimit, gasPrice, data))
// }

// func newKeeper(ctx sdk.Context, chain nexus.ChainName, confHeight int64) types.BaseKeeper {
// 	encCfg := app.MakeEncodingConfig()
// 	paramsK := paramsKeeper.NewKeeper(encCfg.Codec, encCfg.Amino, sdk.NewKVStoreKey("subspace"), sdk.NewKVStoreKey("tsubspace"))
// 	k := keeper.NewKeeper(encCfg.Codec, sdk.NewKVStoreKey("testKey"), paramsK)
// 	k.InitChains(ctx)
// 	funcs.MustNoErr(k.CreateChain(ctx, types.Params{
// 		Chain:               exported.Ethereum.Name,
// 		Network:             network,
// 		ConfirmationHeight:  uint64(confHeight),
// 		TokenCode:           tokenBC,
// 		Burnable:            burnerBC,
// 		RevoteLockingPeriod: 50,
// 		VotingThreshold:     utils.Threshold{Numerator: 15, Denominator: 100},
// 		MinVoterCount:       15,
// 		CommandsGasLimit:    5000000,
// 		Networks: []types.NetworkInfo{{
// 			Name: network,
// 			Id:   sdk.NewIntFromUint64(uint64(rand.I64Between(1, 10))),
// 		}},
// 		EndBlockerLimit: 50,
// 		TransferLimit:   50,
// 	}))
// 	funcs.Must(k.ForChain(ctx, chain)).SetGateway(ctx, types.Address(common.HexToAddress(gateway)))

// 	return k
// }

// func createMsgSignDeploy(details types.TokenDetails) *types.CreateDeployTokenRequest {
// 	account := rand.AccAddr()

// 	asset := types.NewAsset(scalarnet.Scalarnet.Name.String(), scalarnet.NativeAsset)
// 	return types.NewCreateDeployTokenRequest(account, exported.Ethereum.Name.String(), asset, details, types.ZeroAddress, sdk.NewUint(uint64(rand.PosI64())).String())
// }

// func createDetails(name, symbol string) types.TokenDetails {
// 	decimals := rand.Bytes(1)[0]
// 	capacity := sdk.NewIntFromUint64(uint64(rand.PosI64()))

// 	return types.NewTokenDetails(name, symbol, decimals, capacity)
// }

// func createMockERC20Token(asset string, details types.TokenDetails) types.ERC20Token {
// 	meta := types.ERC20TokenMetadata{
// 		Asset:        asset,
// 		Details:      details,
// 		Status:       types.Initialized,
// 		TokenAddress: types.Address(common.BytesToAddress(rand.Bytes(common.AddressLength))),
// 		ChainID:      sdk.NewIntFromUint64(uint64(rand.I64Between(1, 10))),
// 	}
// 	return types.CreateERC20Token(
// 		func(meta types.ERC20TokenMetadata) {},
// 		meta,
// 	)
// }

// func randomNormalizedStr(size int) string {
// 	return strings.ReplaceAll(utils.NormalizeString(rand.Str(size)), utils.DefaultDelimiter, "-")
// }
