package hash

import "math/rand"

// Online builds simhash data sketches online
type Online struct {
	hyperplanes *[]Hyperplane
}

// NewOnline constructs a simhash.Online builder given the number of hyperplanes
// to construct and the dimension of the input vectors
func NewOnline(hyperplaneCount, dim uint) Online {
	return Online{hyperplanes: NewHyperplanes(hyperplaneCount, dim)}
}

// Hash constructs a simhash data sketch of the given vector
func (o Online) Hash(vector *[]float64) *[]Bit {
	return NewSimhash(o.hyperplanes, vector)
}

// Offline sketches each vector in an offline fashion
func Offline(vectors *[][]float64, hyperplaneCount uint) *[][]Bit {
	simhashs := make([][]Bit, len(*vectors))
	hyperplanes := NewHyperplanes(hyperplaneCount, uint(len((*vectors)[0])))
	for i, vector := range *vectors {
		simhashs[i] = *NewSimhash(hyperplanes, &vector)
	}
	return &simhashs
}

// NewSimhash constructs a simhash data sketch of the given vector
func NewSimhash(hyperplanes *[]Hyperplane, vector *[]float64) *[]Bit {
	simhash := make([]Bit, len(*hyperplanes))
	for i, hyperplane := range *hyperplanes {
		var dotProduct float64 // the dot product of hyperplanes[i] and vector
		for j, v := range *vector {
			dotProduct += hyperplane[j] * v
		}
		if dotProduct >= 0 {
			simhash[i] = Bit(1)
		} else {
			simhash[i] = Bit(0)
		}
	}
	return &simhash
}

// Hyperplane is a dim dimensional hyperplane
type Hyperplane []float64

// NewHyperplanes constructs a hyperplane given number of hyperplanes to
// construct and the dimension of each hyperplane
func NewHyperplanes(count, dim uint) *[]Hyperplane {
	hyperplanes := make([]Hyperplane, count)
	for i := uint(0); i < count; i++ {
		hyperplane := make(Hyperplane, dim)
		for j := uint(0); j < dim; j++ {
			hyperplane[j] = rand.NormFloat64()
		}
		hyperplanes[i] = hyperplane
	}
	return &hyperplanes
}
