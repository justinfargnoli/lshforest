package lshforest

import (
	"github.com/justinfargnoli/lshforest/pkg/hash"
	"github.com/justinfargnoli/lshforest/pkg/lshtree"
)

// LSHForest is an index of high-dimensional data based on cosine similarity
type LSHForest struct {
	trees   []lshtree.Trie
	hashers []hash.Hasher
}

// NewDefault constructs an LSHForest struct for cosine similarity with
// sensible defaults
func NewDefault(dim, metric uint) *LSHForest {
	return New(5, 20, dim, metric)
}

var (
	// Cosine indicates to use cosine similarity and simhash
	Cosine = uint(0)
)

// New constructs an LSHForest struct for cosine similarity. l := the
// number of trees in the forest of LSHForest. maxK := the maximum number of
// hash functions. The larger maxK is, the more accurate LSHForest is and the
// more space LSHForest takes up. dim := the dimension of the input vectors
func New(l, maxK, dim, metric uint) *LSHForest {
	var trees []lshtree.Trie
	var hashers []hash.Hasher
	for i := uint(0); i < l; i++ {
		trees = append(trees, lshtree.NewTrie())
		if metric == Cosine {
			hashers = append(hashers, hash.NewOnline(maxK, dim))
		} else {
			panic("lshforest invalid hasher")
		}
	}
	return &LSHForest{trees: trees, hashers: hashers}
}

// Insert puts the vector into the LSHForest
func (f *LSHForest) Insert(vector *[]float64, value interface{}) {
	for i, tree := range f.trees {
		tree.Insert(lshtree.NewElement(f.hashers[i].Hash(vector), vector, value))
	}
}

// func (f *LSHForest) syncAscend()
