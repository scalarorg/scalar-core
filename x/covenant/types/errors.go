package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// module errors
var (
	// Code 1 is a reserved code for internal errors and should not be used for anything else
	_               = sdkerrors.Register(ModuleName, 1, "internal error")
	ErrQuery        = sdkerrors.Register(ModuleName, 2, "query error")
	ErrTapScriptSig = sdkerrors.Register(ModuleName, 3, "tap script sig error")
	// ErrRotationInProgress     = sdkerrors.Register(ModuleName, 3, "key rotation in progress")
	// ErrSignCommandsInProgress = sdkerrors.Register(ModuleName, 4, "signing for command batch in progress")
)
