package a_range_struct

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPrefixSumRange(t *testing.T) {
	tree := NewPrefixSumRange([]int{1, 2, 3, 4, 5})
	require.Equal(t, []int{0, 1, 3, 6, 10, 15}, tree.prefixSum)
}

func TestPrefixSumRangeACC(t *testing.T) {
	tree := NewPrefixSumRange([]int{1, 2, 3, 4, 5})
	require.Equal(t, 5, tree.ACC(1, 2))
}
