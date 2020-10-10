package algorithms

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	ts := []struct {
		xs       []int
		expected *Tree
	}{
		{
			xs: []int{8, 10, 9, 11, 5, 6, 3},
			expected: &Tree{
				Value: 8,
				Empty: false,
				Left: &Tree{
					Value: 5,
					Empty: false,
					Left:  Leaf(3),
					Right: Leaf(6),
				},
				Right: &Tree{
					Value: 10,
					Empty: false,
					Left:  Leaf(9),
					Right: Leaf(11),
				},
			},
		},
	}
	for i, j := range ts {
		a := &Tree{Empty: true}
		for _, k := range j.xs {
			Insert(a, k)
		}
		require.Equal(t, j.expected, a, "At %d", i)
	}
}

func TestBInsert(t *testing.T) {
	ts := []struct {
		xs       []int
		expected *BTree
	}{
		{
			xs: []int{8, 10, 9, 11, 5, 6, 3},
			expected: &BTree{
				Nodes: []BNode{
					{Value: 8, Left: 1, Right: 2},
					{Value: 5, Left: 9, Right: 10},
					{Value: 10, Left: 3, Right: 4},
					{Value: 9, Left: 5, Right: 6},
					{Value: 11, Left: 7, Right: 8},
					{Empty: true},
					{Empty: true},
					{Empty: true},
					{Empty: true},
					{Value: 3, Left: 13, Right: 14},
					{Value: 6, Left: 11, Right: 12},
					{Empty: true},
					{Empty: true},
					{Empty: true},
					{Empty: true},
				},
			},
		},
	}
	for i, j := range ts {
		b := NewBTree()
		for _, k := range j.xs {
			b.Insert(k)
		}
		t.Logf("%v", b.Nodes)
		require.Equal(t, j.expected, b, "At %d", i)
	}
}
