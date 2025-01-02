package types

import (
	fmt "fmt"

	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
)

// SigRouter implements a sig router based on module name
type CovenantRouter interface {
	AddHandler(module string, handler exported.CovenantHandler) CovenantRouter
	HasHandler(module string) bool
	GetHandler(module string) exported.CovenantHandler
	Seal()
}

var _ CovenantRouter = (*router)(nil)

type router struct {
	handlers map[string]exported.CovenantHandler
	sealed   bool
}

// NewSigRouter is the contructor for sig router
func NewCovenantRouter() CovenantRouter {
	return &router{
		handlers: make(map[string]exported.CovenantHandler),
	}
}

// AddHandler registers a new handler for the given module; panics if the
// router is sealed, if the module is invalid, or if the module has been
// registered already.
func (r *router) AddHandler(module string, handler exported.CovenantHandler) CovenantRouter {
	if handler == nil {
		panic("nil handler received")
	}

	if r.sealed {
		panic("router already sealed")
	}

	funcs.MustNoErr(utils.ValidateString(module))

	if r.HasHandler(module) {
		panic(fmt.Sprintf("handler for module %s already registered", module))
	}

	r.handlers[module] = handler

	return r
}

// HasHandler returns true if the router has a handler registered for the
// given module
func (r router) HasHandler(module string) bool {
	_, ok := r.handlers[module]

	return ok
}

// GetHandler returns the handler for the given module.
func (r router) GetHandler(module string) exported.CovenantHandler {
	if !r.HasHandler(module) {
		panic(fmt.Sprintf("no handler for module %s registered", module))
	}

	return r.handlers[module]
}

// Seal prevents additional handlers from being added to the router
func (r *router) Seal() {
	r.sealed = true
}
