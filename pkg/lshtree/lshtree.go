package lshtree

import "github.com/justinfargnoli/lshforest/pkg/hash"

// Element is an element in the trie
type Element struct {
	hash  []hash.Bit
	value interface{}
}

// LSHTree is a trie within the LSHForest
type LSHTree interface {}