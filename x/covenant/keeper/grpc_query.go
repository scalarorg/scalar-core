package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServiceServer = Querier{}

// Querier implements the grpc querier
type Querier struct {
	keeper *Keeper
}

// NewGRPCQuerier returns a new Querier
func NewGRPCQuerier(k *Keeper) Querier {
	return Querier{
		keeper: k,
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

// // KeyID returns the key ID assigned to a given chain
// func (q Querier) KeyID(c context.Context, req *types.KeyIDRequest) (*types.KeyIDResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	keyID, ok := q.keeper.GetCurrentKeyID(ctx, nexus.ChainName(req.Chain))
// 	if !ok {
// 		return nil, status.Error(codes.NotFound, sdkerrors.Wrap(types.ErrMultisig, fmt.Sprintf("key id not found for chain [%s]", req.Chain)).Error())
// 	}

// 	return &types.KeyIDResponse{KeyID: keyID}, nil
// }

// // Key returns the key corresponding to a given key ID
// func (q Querier) Key(c context.Context, req *types.KeyRequest) (*types.KeyResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(c)

// 	if _, ok := q.keeper.GetKeygenSession(ctx, req.KeyID); ok {
// 		return nil, status.Error(codes.NotFound, sdkerrors.Wrap(types.ErrMultisig, fmt.Sprintf("keygen in progress for key id [%s]", req.KeyID)).Error())
// 	}

// 	key, ok := q.keeper.GetKey(ctx, req.KeyID)
// 	if !ok {
// 		return nil, status.Error(codes.NotFound, sdkerrors.Wrap(types.ErrMultisig, fmt.Sprintf("key not found for key id [%s]", req.KeyID)).Error())
// 	}

// 	participants := slices.Map(key.GetParticipants(), func(p sdk.ValAddress) types.KeygenParticipant {
// 		return types.KeygenParticipant{
// 			Address: p.String(),
// 			Weight:  key.GetWeight(p),
// 			PubKey:  funcs.MustOk(key.GetPubKey(p)).String(),
// 		}
// 	})
// 	sort.SliceStable(participants, func(i, j int) bool {
// 		return participants[i].Weight.GT(participants[j].Weight)
// 	})

// 	return &types.KeyResponse{
// 		KeyID:              req.KeyID,
// 		State:              key.GetState(),
// 		StartedAt:          key.GetHeight(),
// 		StartedAtTimestamp: key.GetTimestamp(),
// 		ThresholdWeight:    key.GetMinPassingWeight(),
// 		BondedWeight:       key.GetBondedWeight(),
// 		Participants:       participants,
// 	}, nil
// }
