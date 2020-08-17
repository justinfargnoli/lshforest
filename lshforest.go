package lshforest

import "github.com/justinfargnoli/simhash"

// Bit represents a bit in an element's hash bit array
type Bit uint8

// Element is an element in the trie
type Element struct {
	hash  []Bit
	value interface{}
}

// LSHTree is a tree which makes up to forest of LSHForest
type LSHTree interface {
	Insert(*Element)
	Get(*[]Bit) []*Element
}

// LSHForest is an index of high-dimensional data
type LSHForest struct {
	trees []LSHTree
	maxK  uint
}

// NewLSHForestDefault constructs an LSHForest struct with sensible defaults
func NewLSHForestDefault() LSHForest {
	return NewLSHForest(5, 20)
}

// NewLSHForest constructs an LSHForest
// l := the number of trees in the forest of LSHForest
// maxK := the maximum number of hash functions. The larger maxK is, the more
// accurate LSHForest is and the more space LSHForest takes up.
func NewLSHForest(l, maxK uint) LSHForest {
	var trees []LSHTree
	for i := uint(0); i < l; i++ {
		trie := NewTrie()
		trees = append(trees, &trie)
	}
	return LSHForest{trees: trees, maxK: maxK}
}

// Insert puts the vector into the LSHForest
func (f *LSHForest) Insert(vector []float64) {
	hash
}