package algorithms

// Tree t
type Tree struct {
	Value int
	Empty bool
	Left  *Tree
	Right *Tree
}

// Insert inserts
func Insert(a *Tree, v int) {
	if a.Empty {
		a.Value = v
		a.Empty = false
	} else {
		if a.Right == nil {
			a.Right = &Tree{Empty: true}
		}
		if a.Left == nil {
			a.Left = &Tree{Empty: true}
		}
		var branch *Tree
		// greater than root goes right, left otherwise
		if a.Value < v {
			branch = a.Right
		} else {
			branch = a.Left
		}
		Insert(branch, v)
	}
}

// Leaf leafs
func Leaf(v int) *Tree {
	return &Tree{Value: v, Empty: false}
}

// BNode bnodes
type BNode struct {
	Value int
	Empty bool
	Left  uint
	Right uint
}

// BTree btrees
type BTree struct {
	Nodes []BNode
}

// NewBTree bebe
func NewBTree() *BTree {
	return &BTree{Nodes: []BNode{{Empty: true}}}
}

// Insert a
func (b *BTree) Insert(v int) {
	current := &b.Nodes[0]
	for !current.Empty {
		if v < current.Value {
			current = &b.Nodes[current.Left]
		} else {
			current = &b.Nodes[current.Right]
		}
	}
	current.Value, current.Empty = v, false
	current.Left = uint(len(b.Nodes))
	current.Right = uint(len(b.Nodes)) + 1
	b.Nodes = append(b.Nodes, BNode{Empty: true}, BNode{Empty: true})
}

// Remove r
func (b *BTree) Remove(v int) (ok bool) {
	current := &b.Nodes[0]
	for !current.Empty && current.Value != v {
		if v < current.Value {
			current = &b.Nodes[current.Left]
		} else if v > current.Value {
			current = &b.Nodes[current.Right]
		}
	}
	ok = current.Value == v
	if ok {
		left, previous := current, current
		for !left.Empty {
			previous = left
			left = &b.Nodes[left.Left]
		}
		current.Value = previous.Value
		previous.Empty = true
	}
	return
}

// Postorder postorders
type Postorder struct {
	Tree     *BTree
	Path     []int
	PathProc []int
	Current  int
}

// Next n
func (p *Postorder) Next() (i int, def bool) {
	return
}
