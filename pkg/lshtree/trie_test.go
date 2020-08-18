package lshtree

import "github.com/justinfargnoli/lshforest/pkg/hash"

import "testing"

var (
	elements1 = []*Element{
		{hash: []hash.Bit{0}, value: "a"},
		{hash: []hash.Bit{1}, value: "b"},
	}

	elements2 = []*Element{
		{hash: []hash.Bit{0, 0}, value: "a"},
		{hash: []hash.Bit{0, 1}, value: "b"},
		{hash: []hash.Bit{1, 0}, value: "c"},
		{hash: []hash.Bit{1, 1}, value: "d"},
	}

	elements2Bucket = []*Element{
		{hash: []hash.Bit{0, 0}, value: "a"},
		{hash: []hash.Bit{0, 0}, value: "b"},
		{hash: []hash.Bit{1, 1}, value: "c"},
		{hash: []hash.Bit{1, 1}, value: "d"},
	}

	elements3 = []*Element{
		{hash: []hash.Bit{0, 0, 0}, value: "a"},
		{hash: []hash.Bit{0, 0, 1}, value: "b"},
		{hash: []hash.Bit{0, 1, 0}, value: "c"},
		{hash: []hash.Bit{0, 1, 1}, value: "d"},
		{hash: []hash.Bit{1, 0, 0}, value: "e"},
		{hash: []hash.Bit{1, 0, 1}, value: "f"},
		{hash: []hash.Bit{1, 1, 0}, value: "g"},
		{hash: []hash.Bit{1, 1, 1}, value: "h"},
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

func EqArrElement(e1, e2 []*Element) bool {
	if len(e1) != len(e2) {
		return false
	}
	for i := range e1 {
		if e1[i] != e2[i] {
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
	trie.Get(&[]hash.Bit{})
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

func testGet(t *testing.T, e1, e2 []*Element) {
	if !EqArrElement(e1, e2) {
		t.Fatalf("get: (%v) | correct: (%v)\n", e1, e2)
	}
}

func TestGet(t *testing.T) {
	trie := NewTrie()
	insert(&trie, elements1)
	testGet(t, trie.Get(&elements1[0].hash), []*Element{elements1[0]})
	testGet(t, trie.Get(&elements1[1].hash), []*Element{elements1[1]})

	trie = NewTrie()
	insert(&trie, elements2)
	testGet(t, trie.Get(&elements2[0].hash), []*Element{elements2[0]})
	testGet(t, trie.Get(&elements2[1].hash), []*Element{elements2[1]})
	testGet(t, trie.Get(&elements2[2].hash), []*Element{elements2[2]})
	testGet(t, trie.Get(&elements2[3].hash), []*Element{elements2[3]})

	trie = NewTrie()
	insert(&trie, elements2Bucket)
	testGet(t, trie.Get(&elements2Bucket[0].hash),
		[]*Element{elements2Bucket[0], elements2Bucket[1]})
	testGet(t, trie.Get(&elements2Bucket[2].hash),
		[]*Element{elements2Bucket[2], elements2Bucket[3]})
	testGet(t, trie.Get(&[]hash.Bit{0, 1}), []*Element{})
}
