package types

const (
	// ModuleName defines the module name
	ModuleName = "lone"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_lone"
)

var (
	ParamsKey = []byte("p_lone")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
