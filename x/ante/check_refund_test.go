package ante_test

import (
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/stretchr/testify/assert"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	abci "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/scalarorg/scalar-core/app"
	"github.com/scalarorg/scalar-core/app/params"
	"github.com/scalarorg/scalar-core/testutils/fake"
	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/x/ante"
	"github.com/scalarorg/scalar-core/x/ante/types/mock"
	auxiliarytypes "github.com/scalarorg/scalar-core/x/auxiliary/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/types"
	"github.com/scalarorg/scalar-core/x/nexus/exported"
	nexustypes "github.com/scalarorg/scalar-core/x/nexus/types"
	rewardtypes "github.com/scalarorg/scalar-core/x/reward/types"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/types"
	snapshotkeeper "github.com/scalarorg/scalar-core/x/snapshot/keeper"
	snapshottypes "github.com/scalarorg/scalar-core/x/snapshot/types"
	votetypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func TestCheckRefundFeeDecorator_AnteHandle(t *testing.T) {
	app.SetConfig()
	app.WasmEnabled = "true"
	app.IBCWasmHooksEnabled = "true"
	version.Version = "0.35.0"
	encConfig := app.MakeEncodingConfig()
	sender := rand.AccAddr()

	testCases := []struct {
		label       string
		succeeds    bool
		refundCount int
		msgs        []sdk.Msg
	}{
		{
			label:       "empty message",
			succeeds:    true,
			refundCount: 0,
			msgs:        []sdk.Msg{},
		},
		{
			label:       "single non-refundable message",
			succeeds:    true,
			refundCount: 0,
			msgs:        []sdk.Msg{&exported.WasmMessage{}},
		},
		{
			label:       "multiple non-refundable messages",
			succeeds:    true,
			refundCount: 0,
			msgs:        []sdk.Msg{&exported.WasmMessage{}, &chainsTypes.ConfirmSourceTxsRequest{}, &scalarnet.LinkRequest{}},
		},
		{
			label:       "single refundable message",
			succeeds:    true,
			refundCount: 1,
			msgs:        []sdk.Msg{rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{})},
		},
		{
			label:       "multiple refundable messages",
			succeeds:    true,
			refundCount: 3,
			msgs: []sdk.Msg{
				rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{}),
				rewardtypes.NewRefundMsgRequest(sender, &multisig.SubmitSignatureRequest{}),
				rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{}),
			},
		},
		{
			label:       "multiple mixed messages",
			succeeds:    false,
			refundCount: 0,
			msgs: []sdk.Msg{
				rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{}),
				rewardtypes.NewRefundMsgRequest(sender, &multisig.SubmitSignatureRequest{}),
				&scalarnet.LinkRequest{},
				rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{}),
			},
		},
		{
			label:       "multiple non-refundable batched messages",
			succeeds:    true,
			refundCount: 0,
			msgs: []sdk.Msg{auxiliarytypes.NewBatchRequest(
				sender,
				[]sdk.Msg{&exported.WasmMessage{}, &chainsTypes.ConfirmSourceTxsRequest{}, &scalarnet.LinkRequest{}},
			)},
		},
		{
			label:       "multiple refundable batched messages",
			succeeds:    true,
			refundCount: 4,
			msgs: []sdk.Msg{auxiliarytypes.NewBatchRequest(
				sender,
				[]sdk.Msg{
					rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{}),
					rewardtypes.NewRefundMsgRequest(sender, &multisig.SubmitSignatureRequest{}),
					rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{}),
					rewardtypes.NewRefundMsgRequest(sender, &multisig.SubmitSignatureRequest{}),
				},
			)},
		},
		{
			label:       "multiple mixed batched messages",
			succeeds:    false,
			refundCount: 0,
			msgs: []sdk.Msg{auxiliarytypes.NewBatchRequest(
				sender,
				[]sdk.Msg{
					rewardtypes.NewRefundMsgRequest(sender, &multisig.SubmitSignatureRequest{}),
					&scalarnet.LinkRequest{},
					rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{}),
				},
			)},
		},
		{
			label:       "non-refundable message and refundable batched messages",
			succeeds:    false,
			refundCount: 0,
			msgs: []sdk.Msg{
				auxiliarytypes.NewBatchRequest(
					sender,
					[]sdk.Msg{
						rewardtypes.NewRefundMsgRequest(sender, &multisig.SubmitSignatureRequest{}),
						rewardtypes.NewRefundMsgRequest(sender, &votetypes.VoteRequest{}),
					},
				),
				&scalarnet.LinkRequest{},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.label, func(t *testing.T) {
			ctx := prepareCtx()
			anteHandler, rewardKeeper := prepareAnteHandler(ctx, sender, encConfig)

			// keep track of all the fees to be refunded
			var feeTotal sdk.Coins
			rewardKeeper.SetPendingRefundFunc = func(_ sdk.Context, _ rewardtypes.RefundMsgRequest, refund rewardtypes.Refund) error {
				feeTotal = feeTotal.Add(refund.Fees...)
				return nil
			}

			tx := prepareTx(encConfig, testCase.msgs)
			_, err := anteHandler(ctx, tx, false)

			if testCase.succeeds {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			if testCase.refundCount > 0 {
				assert.Len(t, rewardKeeper.SetPendingRefundCalls(), testCase.refundCount)
				assert.Equal(t, tx.GetFee(), feeTotal)
			} else {
				assert.Len(t, rewardKeeper.SetPendingRefundCalls(), 0)
			}
		})
	}
}

func prepareAnteHandler(ctx sdk.Context, sender sdk.AccAddress, encConfig params.EncodingConfig) (sdk.AnteHandler, *mock.RewardMock) {
	baseApp := app.NewScalarApp(
		log.TestingLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		nil,
		"",
		"",
		0,
		encConfig,
		simapp.EmptyAppOptions{},
		[]wasm.Option{},
	)

	// set up proxy and validator because the refund ante handler checks the sender
	bankKeeper := app.GetKeeper[bankkeeper.BaseKeeper](baseApp.Keepers)
	bankKeeper.SetParams(ctx, banktypes.DefaultParams())
	balance := sdk.NewCoins(sdk.NewInt64Coin("stake", 1e10))
	funcs.MustNoErr(bankKeeper.MintCoins(ctx, nexustypes.ModuleName, balance))
	funcs.MustNoErr(bankKeeper.SendCoinsFromModuleToAccount(ctx, nexustypes.ModuleName, sender, balance))

	stakingKeeper := app.GetKeeper[stakingkeeper.Keeper](baseApp.Keepers)
	stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams())
	validator := stakingtypes.Validator{OperatorAddress: rand.ValAddr().String()}
	stakingKeeper.SetValidator(ctx, validator)

	snapshotKeeper := app.GetKeeper[snapshotkeeper.Keeper](baseApp.Keepers)
	snapshotKeeper.SetParams(ctx, snapshottypes.DefaultParams())
	funcs.MustNoErr(snapshotKeeper.ActivateProxy(ctx, validator.GetOperator(), sender))

	rewardKeeper := &mock.RewardMock{}

	anteHandler := ante.NewCheckRefundFeeDecorator(
		encConfig.InterfaceRegistry,
		app.GetKeeper[authkeeper.AccountKeeper](baseApp.Keepers),
		stakingKeeper,
		snapshotKeeper,
		rewardKeeper,
	)

	// call the batch ante handler first, so we can make sure the refund handler works correctly with batches
	return sdk.ChainAnteDecorators(ante.NewBatchDecorator(encConfig.Codec), anteHandler), rewardKeeper
}

func prepareTx(encConfig params.EncodingConfig, msgs []sdk.Msg) sdk.FeeTx {
	sk, _, _ := testdata.KeyTestPubAddr()

	tx := funcs.Must(helpers.GenTx(
		encConfig.TxConfig,
		msgs,
		sdk.NewCoins(sdk.NewInt64Coin("stake", 1000)),
		1000000000,
		"testchain",
		[]uint64{0},
		[]uint64{0},
		sk,
	))
	return tx.(sdk.FeeTx)
}

func prepareCtx() sdk.Context {
	return sdk.NewContext(fake.NewMultiStore(), abci.Header{}, false, log.TestingLogger()).
		WithConsensusParams(&abcitypes.ConsensusParams{
			Block: &abcitypes.BlockParams{MaxGas: 1000000000},
		})
}
