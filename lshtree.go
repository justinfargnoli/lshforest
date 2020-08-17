package lshforest

// LSHTree is a prefix tree which uses a Element.hash, a []Bit, to determine the
// elements prefix
type LSHTree struct {
	root *Node
}

// NewLSHTree constructs an empty LSHTree
func NewLSHTree() LSHTree {
	return LSHTree{}
}

// Preorder performs a preorder traversal of the tree
func (t LSHTree) Preorder(function func(*Node)) {
	if t.root != nil {
		t.root.preorder(function)
	}
}

// Postorder performs a postorder traversal of the tree
func (t LSHTree) Postorder(function func(*Node)) {
	if t.root != nil {
		t.root.postorder(function)
	}
}

// Inorder performs a inorder traversal of the tree
func (t LSHTree) Inorder(function func(*Node)) {
	if t.root != nil {
		t.root.inorder(function)
	}
}

// Insert adds an element to the tire
func (t *LSHTree) Insert(element *Element) {
	if element == nil {
		panic("lshforest/LSHTree LSHTree.Insert()")
	}
	if t.root == nil {
		t.root = &Node{elements: []*Element{element}}
	} else {
		t.root.insert(element, 0)
	}
}

// Get returns elements with equal hash values
func (t *LSHTree) Get(hash *[]Bit) []*Element {
	if hash == nil {
		panic("lshforest/LSHTree LSHTree.Get()")
	}
	if t.root == nil {
		return []*Element{}
	}
	return t.root.get(hash, 0)
}

const (
	left  = 0
	right = 1
)

// Node is a node in the LSHTree
type Node struct {
	elements            []*Element
	left, right, Parent *Node
}

func (n *Node) isInternal() bool {
	return n.left != nil || n.right != nil
}

func (n *Node) isLeaf() bool {
	return n.left == nil && n.right == nil && len(n.elements) == 1
}

func (n *Node) isLeafBucket() bool {
	return n.left == nil && n.right == nil && len(n.elements) >= 1
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
			return n.left.get(hash, depth+1)
		}
		if n.right == nil {
			return []*Element{}
		}
		return n.right.get(hash, depth+1)
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
		if element.hash[depth] == n.elements[0].hash[depth] { // they're going the same way
			if element.hash[depth] == left { // they're going left
				n.left = &Node{Parent: n, elements: n.elements}
				n.elements = []*Element{}
				n.left.insert(element, depth+1)
			} else { // they're going right
				n.right = &Node{Parent: n, elements: n.elements}
				n.elements = []*Element{}
				n.right.insert(element, depth+1)
			}
		} else { // they're going different ways
			if element.hash[depth] == left { // element goes left & node goes right
				n.left = &Node{Parent: n, elements: []*Element{element}}
				n.right = &Node{Parent: n, elements: n.elements}
				n.elements = []*Element{}
			} else { // node goes left & element goes right
				n.left = &Node{Parent: n, elements: n.elements}
				n.right = &Node{Parent: n, elements: []*Element{element}}
				n.elements = []*Element{}
			}
		}
	} else {
		n.elements = append(n.elements, element)
	}
}
