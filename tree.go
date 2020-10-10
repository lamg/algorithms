package algorithms

type Tree struct {
	Value int
	Empty bool
	Left  *Tree
	Right *Tree
}

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

func Leaf(v int) *Tree {
	return &Tree{Value: v, Empty: false}
}

type BNode struct {
	Value int
	Empty bool
	Left  uint
	Right uint
}

type BTree struct {
	Nodes []BNode
}

func NewBTree() *BTree {
	return &BTree{Nodes: []BNode{{Empty: true}}}
}

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
