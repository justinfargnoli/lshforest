package lshforest

// Trie is a prefix tree which uses a Element.hash, a []Bit, to determine the
// elements prefix
type Trie struct {
	root *Node
}

// NewTrie constructs an empty trie
func NewTrie() Trie {
	return Trie{}
}

// Preorder performs a preorder traversal of the tree
func (t Trie) Preorder(function func(*Node)) {
	t.root.preorder(function)
}

// Postorder performs a postorder traversal of the tree
func (t Trie) Postorder(function func(*Node)) {
	t.root.postorder(function)
}

// Inorder performs a inorder traversal of the tree
func (t Trie) Inorder(function func(*Node)) {
	t.root.inorder(function)
}

// Insert adds an element to the tire
func (t Trie) Insert(element *Element) {
	if element == nil {
		panic("lshforest/trie Trie.Insert()")
	}
	if t.root == nil {
		t.root = &Node{elements: []*Element{element}}
	} else {
		t.root.insert(element, 0)
	}
}

// Get returns elements with equal hash values
func (t Trie) Get(hash *[]Bit) []*Element {
	if hash == nil {
		panic("lshforest/trie Trie.Get()")
	}
	return t.root.get(hash, 0)
}

const (
	left  = 0
	right = 1
)

// Bit represents a bit in an element's hash bit array
type Bit uint8

// Element is an element in the trie
type Element struct {
	hash  []Bit
	value interface{}
}

// Node is a node in the trie
type Node struct {
	elements            []*Element
	left, right, Parent *Node
}

func (n *Node) isInternal() bool {
	return n.left != nil || n.right != nil
}

func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *Node) isLeafBucket() bool {
	return n.isLeaf() && len(n.elements) >= 1
}

// Decendants returns the children of the node
func (n *Node) Decendants() (*Node, *Node) {
	return n.left, n.right
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

func (n *Node) get(hash *[]Bit, depth uint) []*Element {
	if n.isInternal() {
		if (*hash)[depth] == left {
			if n.left == nil {
				return []*Element{}
			}
			n.left.get(hash, depth+1)
		} else {
			if n.right == nil {
				return []*Element{}
			}
			n.right.get(hash, depth+1)
		}
	}
	return n.elements
}

func (n *Node) insert(element *Element, depth uint) {
	if n.isInternal() {
		if element.hash[depth] == left {
			if n.left == nil {
				n.left = &Node{elements: []*Element{element}, Parent: n}
			} else {
				n.left.insert(element, depth+1)
			}
		} else {
			if n.right == nil {
				n.right = &Node{elements: []*Element{element}, Parent: n}
			} else {
				n.right.insert(element, depth+1)
			}
		}
	} else if n.isLeaf() {
		if depth == uint(len(element.hash)) {
			n.elements = append(n.elements, element)
			return
		}
		internalNode := &Node{Parent: n.Parent}
		if element.hash[depth] == n.elements[0].hash[depth] { // they're going the same way
			if element.hash[depth] == left { // they're going left
				n.Parent = internalNode
				internalNode.left = n
				internalNode.left.insert(element, depth+1)
			} else { // they're going right
				n.Parent = n
				internalNode.left = n
				internalNode.left.insert(element, depth+1)
			}
		} else { // they're going different ways
			if element.hash[depth] == left { // element goes left & node goes right
				n.Parent = internalNode
				internalNode.right = n
				internalNode.left = &Node{
					elements: []*Element{element},
					Parent:   internalNode,
				}
			} else { // node goes left & element goes right
				n.Parent = internalNode
				internalNode.left = n
				internalNode.right = &Node{
					elements: []*Element{element},
					Parent:   internalNode,
				}
			}
		}
	} else {
		n.elements = append(n.elements, element)
	}
}
