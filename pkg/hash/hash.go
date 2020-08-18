package hash

// Bit represents a bit in an element's hash bit array
type Bit uint8

// Hasher is a type which can hash a []float64
type Hasher interface {
	Hash(*[]float64) *[]Bit
}
