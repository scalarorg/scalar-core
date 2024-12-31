package keeper

import (
	"context"

	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (s msgServer) SubmitTapScriptSig(context.Context, *types.SubmitTapScriptSigRequest) (*types.SubmitTapScriptSigResponse, error) {
	return &types.SubmitTapScriptSigResponse{}, nil
}
