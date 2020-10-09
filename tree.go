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
