package keeper

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	chainsExported "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServiceServer = Querier{}

// Querier implements the grpc querier
type Querier struct {
	keeper    *Keeper
	multisigK types.MultisigKeeper
}

// NewGRPCQuerier returns a new Querier
func NewGRPCQuerier(k *Keeper, m types.MultisigKeeper) Querier {
	return Querier{
		keeper:    k,
		multisigK: m,
	}
}

// Get custodians
func (q Querier) Custodians(c context.Context, req *types.CustodiansRequest) (*types.CustodiansResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	//custodians, ok := q.keeper.findCustodians(ctx, req)
	custodians, ok := q.keeper.GetAllCustodians(ctx)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "custodians not found")
	}

	return &types.CustodiansResponse{
		Custodians: custodians,
	}, nil
}

// Get custodian groups
func (q Querier) Groups(c context.Context, req *types.GroupsRequest) (*types.GroupsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	groups, ok := q.keeper.findCustodianGroups(ctx, req)
	// groups, ok := q.keeper.GetAllCustodianGroups(ctx)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "custodian groups not found")
	}
	return &types.GroupsResponse{
		Groups: groups,
	}, nil
}

// Params returns the params of the module
func (q Querier) Params(context.Context, *types.ParamsRequest) (*types.ParamsResponse, error) {
	return nil, nil
}

func (q Querier) RedeemSession(ctx context.Context, req *types.RedeemSessionRequest) (*types.RedeemSessionResponse, error) {
	hash := chainsExported.Hash(common.BytesToHash(req.UID))
	session, ok := q.keeper.GetRedeemSession(sdk.UnwrapSDKContext(ctx), hash)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "redeem session not found")
	}
	return &types.RedeemSessionResponse{
		Session: session,
	}, nil
}

func (q Querier) UTXOSnapshot(ctx context.Context, req *types.UTXOSnapshotRequest) (*types.UTXOSnapshotResponse, error) {
	hash := chainsExported.Hash(common.BytesToHash(req.UID))
	snapshot, ok := q.keeper.GetUtxoSnapshot(sdk.UnwrapSDKContext(ctx), hash)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "utxo snapshot not found, uid: %x", hash)
	}
	return &types.UTXOSnapshotResponse{
		UtxoSnapshot: snapshot,
	}, nil
}

// optimizeSignatureSet returns optimized signature set, sorted in ascending order by corresponding evm address
func optimizeSignatureSet(operators []chainsTypes.Operator, minPassingWeight sdk.Uint) [][]byte {
	sort.SliceStable(operators, func(i, j int) bool {
		return operators[i].Weight.GT(operators[j].Weight)
	})

	cumWeight := sdk.ZeroUint()
	operators = slices.Filter(operators, func(operator chainsTypes.Operator) bool {
		if cumWeight.GTE(minPassingWeight) {
			return false
		}

		cumWeight = cumWeight.Add(operator.Weight)
		return true
	})

	sort.SliceStable(operators, func(i, j int) bool {
		return bytes.Compare(operators[i].Address.Bytes(), operators[j].Address.Bytes()) < 0
	})

	return slices.Map(operators, func(operator chainsTypes.Operator) []byte { return operator.Signature })
}

func getProof(key multisig.Key, signature multisig.MultiSig) ([]common.Address, []sdk.Uint, sdk.Uint, [][]byte) {
	participantsWithSigs := slices.Filter(key.GetParticipants(), func(v sdk.ValAddress) bool {
		_, ok := signature.GetSignature(v)
		return ok
	})

	operators := slices.Map(participantsWithSigs, func(val sdk.ValAddress) chainsTypes.Operator {
		pubKey := funcs.MustOk(key.GetPubKey(val)).ToECDSAPubKey()
		signature := funcs.Must(chainsTypes.ToSignature(funcs.MustOk(signature.GetSignature(val)), common.BytesToHash(signature.GetPayloadHash()), pubKey))

		return chainsTypes.Operator{
			Address:   crypto.PubkeyToAddress(pubKey),
			Signature: signature.ToHomesteadSig(),
			Weight:    key.GetWeight(val),
		}
	})

	addresses, weights, threshold := chainsTypes.GetMultisigAddressesAndWeights(key)
	signatures := optimizeSignatureSet(operators, key.GetMinPassingWeight())

	return addresses, weights, threshold, signatures
}

func CreateExecuteDataAndSigs(key multisig.Key, cmdData []byte, signature multisig.MultiSig) ([]byte, chainsTypes.Proof, error) {

	addresses, weights, threshold, signatures := getProof(key, signature)

	executeData, err := chainsTypes.CreateExecuteDataMultisig(cmdData, addresses, weights, threshold, signatures)
	if err != nil {
		return nil, chainsTypes.Proof{}, fmt.Errorf("could not create transaction data: %s", err)
	}

	proof := chainsTypes.Proof{
		Addresses:  slices.Map(addresses, common.Address.Hex),
		Weights:    slices.Map(weights, sdk.Uint.String),
		Threshold:  threshold.String(),
		Signatures: slices.Map(signatures, hex.EncodeToString),
	}

	return executeData, proof, nil
}

func commandToResp(ctx sdk.Context, cmd types.StandaloneCommand, multisigK types.MultisigKeeper) (types.StandaloneCommandResponse, error) {

	clog.Greenf("commandToResp: cmd: %+v", cmd)

	if cmd.Is(types.StandaloneCommandStatusSigned) && cmd.GetSignature() != nil { // check signature for unmigrated batches
		signature, ok := cmd.GetSignature().(multisig.MultiSig)
		if ok {
			key := funcs.MustOk(multisigK.GetKey(ctx, signature.GetKeyID()))
			cmdData := cmd.GetData()
			executeData, _, err := CreateExecuteDataAndSigs(key, cmdData, signature)
			if err != nil {
				return types.StandaloneCommandResponse{}, sdkerrors.Wrap(err, "could not create transaction data")
			}

			return types.StandaloneCommandResponse{
				ID:          cmd.GetID(),
				Data:        hex.EncodeToString(cmd.GetData()),
				Status:      cmd.GetStatus(),
				KeyID:       cmd.GetKeyID(),
				ExecuteData: hex.EncodeToString(executeData),
			}, nil
		}
	}

	return types.StandaloneCommandResponse{}, fmt.Errorf("signature is not multisig")
}

func (q Querier) StandaloneCommand(c context.Context, req *types.StandaloneCommandRequest) (*types.StandaloneCommandResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if len(req.ID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "command ID cannot be empty")
	}

	command := q.keeper.GetReserveUTXOCommandByID(ctx, req.ID)
	if command.Is(types.StandaloneCommandStatusNonExistent) {
		err := fmt.Errorf("command with ID %x not found", req.ID)
		return nil, status.Error(codes.NotFound, sdkerrors.Wrap(err, "command not found").Error())
	}

	resp, err := commandToResp(ctx, command, q.multisigK)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &resp, nil
}
