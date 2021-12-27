package a_range_struct

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSegRange(t *testing.T) {
	tree := NewSegRange([]int{1, 2, 3, 4, 5})
	require.Equal(t, []int{0, 15, 10, 5, 9, 1, 2, 3, 4, 5}, tree.segment)
}
