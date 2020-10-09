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
		require.Equal(t, a, j.expected, "At %d", i)
	}
}
