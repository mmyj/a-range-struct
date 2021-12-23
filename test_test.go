package a_range_struct

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type roughRange struct {
	nums  []int
	total int
}

func newRoughRange(nums []int) *roughRange {
	return &roughRange{nums: nums}
}

func (rr *roughRange) ACC(left, right int) {
	if left < 0 || right < 0 || left > right || left >= len(rr.nums) {
		return
	}
	if right >= len(rr.nums) {
		right = len(rr.nums) - 1
	}
	for i := left; i <= right; i++ {
		rr.total += rr.nums[i]
	}
}

func (rr *roughRange) Total() int {
	return rr.total
}

func (rr *roughRange) resume() {
	rr.total = 0
}

func testUnit(t *testing.T, name string, nums []int, accFunc func(r ranger)) {
	t.Run(name, func(t *testing.T) {
		t.Parallel()
		tr := newRoughRange(nums)
		r := NewRange(nums)
		accFunc(tr)
		accFunc(r)
		require.Equal(t, tr.Total(), r.Total())
	})
}

func TestRoughRange(t *testing.T) {
	t.Parallel()
	rr := newRoughRange([]int{1, 2, 3, 4, 5})
	rr.ACC(0, 2)
	rr.ACC(2, 4)
	require.Equal(t, rr.Total(), 18)
	rr.resume()

	rr.ACC(0, 0)
	rr.ACC(1, 1)
	rr.ACC(2, 2)
	rr.ACC(3, 3)
	rr.ACC(4, 4)
	require.Equal(t, rr.Total(), 15)
	rr.resume()

	rr.ACC(0, 0)
	rr.ACC(0, 0)
	require.Equal(t, rr.Total(), 2)
	rr.resume()

	rr.ACC(0, 6)
	require.Equal(t, rr.Total(), 15)
	rr.resume()

	rr.ACC(6, 6)
	require.Equal(t, rr.Total(), 0)
	rr.resume()
}

type accRange struct {
	l int
	r int
}

func TestRange(t *testing.T) {
	numsList := [][]int{
		{1, 2, 3, 4, 5},
		{},
	}
	accsList := [][]accRange{
		// 全量单个无交集
		{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}},
		// 重复单个
		{{0, 0}, {0, 0}},
		// 全量
		{{0, 4}},
		// 有交集
		{{0, 3}, {2, 4}},
		// 右区间越界
		{{0, 4}, {0, 6}},
		// 左区间越界
		{{0, 4}, {6, 6}},
	}
	for _, nums := range numsList {
		for i, accs := range accsList {
			testUnit(t, fmt.Sprintf("%v_%d", nums, i), nums, func(r ranger) {
				for _, acc := range accs {
					r.ACC(acc.l, acc.r)
				}
			})
		}
	}
}

func makeRandNums(length int) []int {
	re := make([]int, length)
	for i := 0; i < length; i++ {
		re[i] = rand.Intn(10)
	}
	return re
}

func makeRandAccRanges(length int) []accRange {
	re := make([]accRange, length)
	for i := 0; i < length; i++ {
		re[i].l = rand.Intn(10)
		re[i].r = rand.Intn(10)
		if re[i].l > re[i].r {
			re[i].l, re[i].r = re[i].r, re[i].l
		}
	}
	return re
}

func TestRangeRand(t *testing.T) {
	const (
		length   = 100
		testTime = 10
	)
	for i := 0; i < testTime; i++ {
		randNums := makeRandNums(length)
		randAccRanges := makeRandAccRanges(length)
		testUnit(t, fmt.Sprintf("rand_%d", i), randNums, func(r ranger) {
			for _, acc := range randAccRanges {
				r.ACC(acc.l, acc.r)
			}
		})
	}
}

func BenchmarkRangeRandComparison(b *testing.B) {
	lengthList := []int{50, 100, 200, 500, 1000, 10000, 100000}
	for _, length := range lengthList {
		randNums := makeRandNums(length)
		randAccRanges := makeRandAccRanges(length)

		b.Run(strconv.Itoa(length)+"_RoughRange", func(b *testing.B) {
			tr := newRoughRange(randNums)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, acc := range randAccRanges {
					tr.ACC(acc.l, acc.r)
				}
				tr.Total()
			}
		})
		b.Run(strconv.Itoa(length)+"_Range", func(b *testing.B) {
			r := NewRange(randNums)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, acc := range randAccRanges {
					r.ACC(acc.l, acc.r)
				}
				r.Total()
			}
		})
		rangeDescs := make([]rangeDesc, length+1)
		b.Run(strconv.Itoa(length)+"_RangeWithRangeDescs", func(b *testing.B) {
			r := NewRangeWithRangeDescs(randNums, rangeDescs)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, acc := range randAccRanges {
					r.ACC(acc.l, acc.r)
				}
				r.Total()
			}
		})
	}
}

func BenchmarkTestRange(b *testing.B) {
	lengthList := []int{100000}
	for _, length := range lengthList {
		randNums := makeRandNums(length)
		randAccRanges := makeRandAccRanges(length)
		r := newRoughRange(randNums)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, acc := range randAccRanges {
				r.ACC(acc.l, acc.r)
			}
			r.Total()
		}

	}
}

func BenchmarkRange(b *testing.B) {
	lengthList := []int{100000}
	for _, length := range lengthList {
		randNums := makeRandNums(length)
		randAccRanges := makeRandAccRanges(length)

		rangeDescs := make([]rangeDesc, length+1)
		b.ResetTimer()
		r := NewRangeWithRangeDescs(randNums, rangeDescs)
		for i := 0; i < b.N; i++ {
			for _, acc := range randAccRanges {
				r.ACC(acc.l, acc.r)
			}
			r.Total()
		}
	}
}
