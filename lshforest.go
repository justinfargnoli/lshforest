package lshforest

// Bit represents a bit in an element's hash bit array
type Bit uint8

// Element is an element in the trie
type Element struct {
	hash  []Bit
	value interface{}
}

// LSHForest is an index of high-dimensional data based on cosine similarity 
type LSHForest struct {
	trees []LSHTree
}

// NewLSHForestDefault constructs an LSHForest struct with sensible defaults
func NewLSHForestDefault(dim uint) *LSHForest {
	return NewLSHForest(5, 20, dim)
}

// NewLSHForest constructs an LSHForest. l := the number of trees in the forest 
// of LSHForest. maxK := the maximum number of hash functions. The larger maxK 
// is, the more accurate LSHForest is and the more space LSHForest takes up.
// dim := the dimension of the input vectors
func NewLSHForest(l, maxK, dim uint) *LSHForest {
	var trees []LSHTree
	for i := uint(0); i < l; i++ {
		trie := NewLSHTree()
		trees = append(trees, trie)
	}
	return &LSHForest{trees: trees}
}

// Insert puts the vector into the LSHForest
func (f *LSHForest) Insert(vector []float64) {
	
}