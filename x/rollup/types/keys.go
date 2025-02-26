package types

const (
	// ModuleName defines the module name
	ModuleName = "rollup"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rollup"

	// Version defines the current version the IBC module supports
	Version = "rollup-1"

	// PortID is the default port id that module binds to
	PortID = "rollup"
)

var (
	ParamsKey = []byte("p_rollup")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("rollup-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
