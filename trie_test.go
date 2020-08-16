package lshforest

import "testing"

var (
	elements1 = []*Element{
		{hash: []Bit{0}, value: "a"},
		{hash: []Bit{1}, value: "b"},
	}

	elements2 = []*Element{
		{hash: []Bit{0, 0}, value: "a"},
		{hash: []Bit{0, 1}, value: "b"},
		{hash: []Bit{1, 0}, value: "c"},
		{hash: []Bit{1, 1}, value: "d"},
	}

	elements2Bucket = []*Element{
		{hash: []Bit{0, 0}, value: "a"},
		{hash: []Bit{0, 0}, value: "b"},
		{hash: []Bit{1, 1}, value: "c"},
		{hash: []Bit{1, 1}, value: "d"},
	}

	elements3 = []*Element{
		{hash: []Bit{0, 0, 0}, value: "a"},
		{hash: []Bit{0, 0, 1}, value: "b"},
		{hash: []Bit{0, 1, 0}, value: "c"},
		{hash: []Bit{0, 1, 1}, value: "d"},
		{hash: []Bit{1, 0, 0}, value: "e"},
		{hash: []Bit{1, 0, 1}, value: "f"},
		{hash: []Bit{1, 1, 0}, value: "g"},
		{hash: []Bit{1, 1, 1}, value: "h"},
	}
)

func insert(trie *Trie, elements []*Element) {
	for _, element := range elements {
		trie.Insert(element)
	}
}

func valuesInorder(trie Trie) []string {
	var values []string
	trie.Inorder(func(node *Node) {
		for _, element := range node.elements {
			values = append(values, element.value.(string))
		}
	})
	return values
}

func EqArrString(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestEmptyTrie(t *testing.T) {
	trie := NewTrie()
	trie.Preorder(func(node *Node) {})
	trie.Postorder(func(node *Node) {})
	trie.Inorder(func(node *Node) {})
	trie.Get(&[]Bit{})
	trie.Insert(&Element{})
}

func TestInsert(t *testing.T) {
	trie := NewTrie()
	insert(&trie, elements1)
	if !EqArrString(valuesInorder(trie), []string{"a", "b"}) {
		t.FailNow()
	}

	trie = NewTrie()
	insert(&trie, elements2)
	if !EqArrString(valuesInorder(trie), []string{"a", "b", "c", "d"}) {
		t.FailNow()
	}

	trie = NewTrie()
	insert(&trie, elements2Bucket)
	if !EqArrString(valuesInorder(trie), []string{"a", "b", "c", "d"}) {
		t.FailNow()
	}

	trie = NewTrie()
	insert(&trie, elements3)
	if !EqArrString(valuesInorder(trie), []string{"a", "b", "c", "d", "e", "f", "g", "h"}) {
		t.FailNow()
	}
}

func TestGet(t *testing.T) {

}
