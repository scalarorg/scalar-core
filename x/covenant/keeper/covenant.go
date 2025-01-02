package keeper

import types "github.com/scalarorg/scalar-core/x/covenant/types"

func (k *Keeper) SetCovenantRouter(router types.CovenantRouter) {
	if k.covRouter != nil {
		panic("router already set")
	}

	k.covRouter = router

	// In order to avoid invalid or non-deterministic behavior, we seal the router immediately
	// to prevent additional handlers from being registered after the keeper is initialized.
	k.covRouter.Seal()
}

// GetCovenantRouter returns the covenant router. If no router was set, it returns a (sealed) router with no handlers
func (k Keeper) GetCovenantRouter() types.CovenantRouter {
	if k.covRouter == nil {
		k.SetCovenantRouter(types.NewCovenantRouter())
	}

	return k.covRouter
}
