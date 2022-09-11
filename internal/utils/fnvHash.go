package utils

// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#FNV_hash_parameters
const (
	defaultFnvOffsetBasis uint64 = 14695981039346656037
	defaultFnvPrime              = 1099511628211
)

type FnvHash struct {
	fnvOffsetBasis uint64
	fnvPrime       uint64
}

// NewDefault returns a FnvHash with the following default values
// defaultFnvOffsetBasis = 14695981039346656037
// defaultFnvPrime = 1099511628211
func NewDefault() *FnvHash {
	return New(defaultFnvOffsetBasis, defaultFnvPrime)
}

func New(offsetBasis uint64, prime uint64) *FnvHash {
	return &FnvHash{
		fnvOffsetBasis: offsetBasis,
		fnvPrime:       prime,
	}
}

// Hash applies the Fowler–Noll–Vo hash function to the given string
func (f *FnvHash) Hash(key string) uint64 {
	hash := f.fnvOffsetBasis
	sBytes := []byte(key)
	for _, b := range sBytes {
		hash = hash ^ uint64(b)
		hash = hash * f.fnvPrime
	}
	return hash
}
