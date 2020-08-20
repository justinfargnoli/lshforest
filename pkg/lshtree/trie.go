package lshtree

import "github.com/justinfargnoli/lshforest/pkg/hash"

// Trie is a prefix tree which uses a Element.hash, a []Bit, to determine the
// elements prefix
type Trie struct {
	root *Node
}

// NewTrie constructs an empty Trie
func NewTrie() Trie {
	return Trie{}
}

// Preorder performs a preorder traversal of the tree
func (t Trie) Preorder(function func(*Node)) {
	if t.root != nil {
		t.root.preorder(function)
	}
}

// Postorder performs a postorder traversal of the tree
func (t Trie) Postorder(function func(*Node)) {
	if t.root != nil {
		t.root.postorder(function)
	}
}

// Inorder performs a inorder traversal of the tree
func (t Trie) Inorder(function func(*Node)) {
	if t.root != nil {
		t.root.inorder(function)
	}
}

// Insert adds an element to the tire
func (t *Trie) Insert(element Element) {
	if t.root == nil {
		t.root = &Node{Elements: []Element{element}}
	} else {
		t.root.insert(element, 0)
	}
}

// Descend returns the leaf with the larges prefix matching hash
func (t *Trie) Descend(hash *[]hash.Bit) (*Node, uint) {
	if hash == nil {
		panic("lshforest/lshtree/Trie Trie.Descend()")
	}
	if t.root == nil {
		return nil, 0
	}
	return t.root.descend(hash, 0)
}

// Get returns elements with equal hash values
func (t *Trie) Get(hash *[]hash.Bit) *[]Element {
	if hash == nil {
		panic("lshforest/lshtree/Trie Trie.Get()")
	}
	if t.root == nil {
		return &[]Element{}
	}
	return t.root.get(hash, 0)
}

const (
	left  = 0
	right = 1
)

// Node is a node in the Trie
type Node struct {
	Elements            []Element
	left, right, Parent *Node
}

func (n *Node) isInternal() bool {
	return n.left != nil || n.right != nil
}

func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil && len(n.Elements) == 1
}

func (n *Node) isLeafBucket() bool {
	return n.left == nil && n.right == nil && len(n.Elements) >= 1
}

// Decendants returns the children of the node
func (n *Node) Decendants() []*Node {
	var nodes []*Node
	if n.left != nil {
		n.left.preorder(func(node *Node) {
			nodes = append(nodes, node)
		})
	}
	if n.right != nil {
		n.right.preorder(func(node *Node) {
			nodes = append(nodes, node)
		})
	}
	return nodes
}

func (n *Node) preorder(function func(*Node)) {
	function(n)
	if n.left != nil {
		n.left.preorder(function)
	}
	if n.right != nil {
		n.right.preorder(function)
	}
}

func (n *Node) postorder(function func(*Node)) {
	if n.left != nil {
		n.left.postorder(function)
	}
	if n.right != nil {
		n.right.postorder(function)
	}
	function(n)
}

func (n *Node) inorder(function func(*Node)) {
	if n.left != nil {
		n.left.inorder(function)
	}
	function(n)
	if n.right != nil {
		n.right.inorder(function)
	}
}

func (n *Node) descend(hash *[]hash.Bit, depth uint) (*Node, uint) {
	if n.isInternal() {
		if (*hash)[depth] == left {
			if n.left == nil {
				return n.right.descend(hash, depth+1)
			}
			return n.left.descend(hash, depth+1)
		}
		if n.right == nil {
			return n.left.descend(hash, depth+1)
		}
		return n.right.descend(hash, depth+1)
	}
	return n, depth
}

func (n *Node) get(hash *[]hash.Bit, depth uint) *[]Element {
	if n.isInternal() {
		if (*hash)[depth] == left {
			if n.left == nil {
				return &[]Element{}
			}
			return n.left.get(hash, depth+1)
		}
		if n.right == nil {
			return &[]Element{}
		}
		return n.right.get(hash, depth+1)
	}
	return &n.Elements
}

func (n *Node) insert(element Element, depth uint) {
	if n.isInternal() {
		if (*element.hash)[depth] == left {
			if n.left == nil {
				n.left = &Node{Elements: []Element{element}, Parent: n}
			} else {
				n.left.insert(element, depth+1)
			}
		} else {
			if n.right == nil {
				n.right = &Node{Elements: []Element{element}, Parent: n}
			} else {
				n.right.insert(element, depth+1)
			}
		}
	} else if n.isLeaf() {
		if depth == uint(len(*element.hash)) {
			n.Elements = append(n.Elements, element)
			return
		}
		if (*element.hash)[depth] == (*n.Elements[0].hash)[depth] { // they're going the same way
			if (*element.hash)[depth] == left { // they're going left
				n.left = &Node{Parent: n, Elements: n.Elements}
				n.Elements = []Element{}
				n.left.insert(element, depth+1)
			} else { // they're going right
				n.right = &Node{Parent: n, Elements: n.Elements}
				n.Elements = []Element{}
				n.right.insert(element, depth+1)
			}
		} else { // they're going different ways
			if (*element.hash)[depth] == left { // element goes left & node goes right
				n.left = &Node{Parent: n, Elements: []Element{element}}
				n.right = &Node{Parent: n, Elements: n.Elements}
				n.Elements = []Element{}
			} else { // node goes left & element goes right
				n.left = &Node{Parent: n, Elements: n.Elements}
				n.right = &Node{Parent: n, Elements: []Element{element}}
				n.Elements = []Element{}
			}
		}
	} else {
		n.Elements = append(n.Elements, element)
	}
}
