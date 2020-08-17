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

func insert(lshTree *LSHTree, elements []*Element) {
	for _, element := range elements {
		lshTree.Insert(element)
	}
}

func valuesInorder(lshTree LSHTree) []string {
	var values []string
	lshTree.Inorder(func(node *Node) {
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

func TestEmptyLSHTree(t *testing.T) {
	lshTree := NewLSHTree()
	lshTree.Preorder(func(node *Node) {})
	lshTree.Postorder(func(node *Node) {})
	lshTree.Inorder(func(node *Node) {})
	lshTree.Get(&[]Bit{})
	lshTree.Insert(&Element{})
}

func TestInsert(t *testing.T) {
	lshTree := NewLSHTree()
	insert(&lshTree, elements1)
	if !EqArrString(valuesInorder(lshTree), []string{"a", "b"}) {
		t.FailNow()
	}

	lshTree = NewLSHTree()
	insert(&lshTree, elements2)
	if !EqArrString(valuesInorder(lshTree), []string{"a", "b", "c", "d"}) {
		t.FailNow()
	}

	lshTree = NewLSHTree()
	insert(&lshTree, elements2Bucket)
	if !EqArrString(valuesInorder(lshTree), []string{"a", "b", "c", "d"}) {
		t.FailNow()
	}

	lshTree = NewLSHTree()
	insert(&lshTree, elements3)
	if !EqArrString(valuesInorder(lshTree), []string{"a", "b", "c", "d", "e", "f", "g", "h"}) {
		t.FailNow()
	}
}

func testGet(t *testing.T, e1, e2 []*Element) {
	if !EqArrElement(e1, e2) {
		t.Fatalf("get: (%v) | correct: (%v)\n", e1, e2)
	}
}

func TestGet(t *testing.T) {
	lshTree := NewLSHTree()
	insert(&lshTree, elements1)
	testGet(t, lshTree.Get(&elements1[0].hash), []*Element{elements1[0]})
	testGet(t, lshTree.Get(&elements1[1].hash), []*Element{elements1[1]})

	lshTree = NewLSHTree()
	insert(&lshTree, elements2)
	testGet(t, lshTree.Get(&elements2[0].hash), []*Element{elements2[0]})
	testGet(t, lshTree.Get(&elements2[1].hash), []*Element{elements2[1]})
	testGet(t, lshTree.Get(&elements2[2].hash), []*Element{elements2[2]})
	testGet(t, lshTree.Get(&elements2[3].hash), []*Element{elements2[3]})

	lshTree = NewLSHTree()
	insert(&lshTree, elements2Bucket)
	testGet(t, lshTree.Get(&elements2Bucket[0].hash),
		[]*Element{elements2Bucket[0], elements2Bucket[1]})
	testGet(t, lshTree.Get(&elements2Bucket[2].hash),
		[]*Element{elements2Bucket[2], elements2Bucket[3]})
	testGet(t, lshTree.Get(&[]Bit{0, 1}), []*Element{})
}
