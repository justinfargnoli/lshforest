package lshtree

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	var lshTree LSHTree

	trie := NewTrie()
	lshTree = &trie

	_ = fmt.Sprint(lshTree)
}
