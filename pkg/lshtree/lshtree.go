package lshtree

import "github.com/justinfargnoli/lshforest/pkg/hash"

// Element is an element in the trie
type Element struct {
	hash   *[]hash.Bit
	Vector *[]float64
	Value  interface{}
}

// NewElement constructs an element stored in the node of a LSHTree
func NewElement(hash *[]hash.Bit, vector *[]float64, value interface{}) Element {
	return Element{hash: hash, Vector: vector, Value: value}
}

// LSHTree is a trie within the LSHForest
type LSHTree interface{}
