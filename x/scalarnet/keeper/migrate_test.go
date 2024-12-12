package keeper_test

// import (
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	params "github.com/cosmos/cosmos-sdk/x/params/types"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/tendermint/tendermint/libs/log"
// 	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

// 	"github.com/axelarnetwork/axelar-core/testutils/fake"
// 	"github.com/axelarnetwork/axelar-core/testutils/rand"
// 	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
// 	nexusmock "github.com/scalarorg/scalar-core/x/nexus/exported/mock"
// 	nexustypes "github.com/scalarorg/scalar-core/x/nexus/types"
// 	"github.com/axelarnetwork/utils/funcs"
// 	. "github.com/axelarnetwork/utils/test"
// 	"github.com/scalarorg/scalar-core/app"
// 	"github.com/scalarorg/scalar-core/x/scalarnet/keeper"
// 	"github.com/scalarorg/scalar-core/x/scalarnet/types"
// 	"github.com/scalarorg/scalar-core/x/scalarnet/types/mock"
// 	axelartestutils "github.com/scalarorg/scalar-core/x/scalarnet/types/testutils"
// )

// func TestMigrate6to7(t *testing.T) {
// 	var (
// 		bank                   *mock.BankKeeperMock
// 		account                *mock.AccountKeeperMock
// 		nexusK                 *mock.NexusMock
// 		lockableAsset          *nexusmock.LockableAssetMock
// 		transfers              []types.IBCTransfer
// 		balances               sdk.Coins
// 		nexusModuleAccAddr     sdk.AccAddress
// 		axelarnetModuleAccAddr sdk.AccAddress
// 	)

// 	encCfg := app.MakeEncodingConfig()
// 	subspace := params.NewSubspace(encCfg.Codec, encCfg.Amino, sdk.NewKVStoreKey("axelarnetKey"), sdk.NewKVStoreKey("tAxelarnetKey"), "scalarnet")
// 	k := keeper.NewKeeper(encCfg.Codec, sdk.NewKVStoreKey("scalarnet"), subspace, &mock.ChannelKeeperMock{}, &mock.FeegrantKeeperMock{})
// 	ibcK := keeper.NewIBCKeeper(k, &mock.IBCTransferKeeperMock{})
// 	ctx := sdk.NewContext(fake.NewMultiStore(), tmproto.Header{}, false, log.TestingLogger())

// 	Given("keeper setup before migration", func() {
// 		bank = &mock.BankKeeperMock{
// 			SendCoinsFromModuleToModuleFunc: func(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error {
// 				return nil
// 			},
// 		}
// 		nexusModuleAccAddr = rand.AccAddr()
// 		axelarnetModuleAccAddr = rand.AccAddr()
// 		account = &mock.AccountKeeperMock{
// 			GetModuleAddressFunc: func(module string) sdk.AccAddress {
// 				switch module {
// 				case types.ModuleName:
// 					return axelarnetModuleAccAddr
// 				case nexustypes.ModuleName:
// 					return nexusModuleAccAddr
// 				default:
// 					panic("unexpected module")
// 				}
// 			},
// 		}
// 		lockableAsset = &nexusmock.LockableAssetMock{
// 			LockFromFunc: func(ctx sdk.Context, fromAddr sdk.AccAddress) error {
// 				return nil
// 			},
// 			GetAssetFunc: func() sdk.Coin {
// 				return rand.Coin()
// 			},
// 		}
// 		nexusK = &mock.NexusMock{
// 			NewLockableAssetFunc: func(ctx sdk.Context, ibc nexustypes.IBCKeeper, bank nexustypes.BankKeeper, coin sdk.Coin) (nexus.LockableAsset, error) {
// 				return lockableAsset, nil
// 			},
// 		}
// 	}).
// 		When("there are some failed transfers and Scalarnet module account has balances", func() {
// 			balances = sdk.NewCoins(rand.Coin(), rand.Coin(), rand.Coin())
// 			bank.SpendableCoinsFunc = func(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
// 				return balances
// 			}

// 			for i := 0; i < 50; i++ {
// 				transfer := axelartestutils.RandomIBCTransfer()
// 				if i%2 == 0 {
// 					transfer.Status = types.TransferFailed
// 				}
// 				transfers = append(transfers, transfer)
// 				assert.NoError(t, k.EnqueueIBCTransfer(ctx, transfer))
// 			}
// 		}).
// 		Then("the migration should lock back to escrow account and update sender of failed transfers", func(t *testing.T) {
// 			err := keeper.Migrate6to7(k, bank, account, nexusK, ibcK)(ctx)
// 			assert.NoError(t, err)
// 			assert.Len(t, lockableAsset.LockFromCalls(), 3)
// 			for _, call := range lockableAsset.LockFromCalls() {
// 				assert.Equal(t, nexusModuleAccAddr, call.FromAddr)
// 			}

// 			assert.Len(t, bank.SendCoinsFromModuleToModuleCalls(), 3)
// 			for _, call := range bank.SendCoinsFromModuleToModuleCalls() {
// 				assert.Equal(t, types.ModuleName, call.SenderModule)
// 				assert.Equal(t, nexustypes.ModuleName, call.RecipientModule)
// 			}

// 			for _, transfer := range transfers {
// 				actual := funcs.MustOk(k.GetTransfer(ctx, transfer.ID))
// 				if transfer.Status == types.TransferFailed {
// 					assert.Equal(t, types.AxelarIBCAccount, actual.Sender)
// 				} else {
// 					assert.Equal(t, transfer.Sender, actual.Sender)
// 				}
// 			}
// 		}).
// 		Run(t)
// }
