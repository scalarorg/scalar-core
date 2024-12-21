package types

const (
	ModuleName = "protocol"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for legacy query routing
	QuerierRoute = ModuleName

	// RestRoute to be used for rest routing
	RestRoute = ModuleName

	DefaultProtocolName = "scalar"
)
