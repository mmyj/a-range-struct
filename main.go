package a_range_struct

type ranger interface {
	// ACC 标记区间[left, right]
	ACC(left, right int)
	// Total 累加每次ACC标记区间内元素的和
	Total() int
}

type Range struct {
	nums       []int
	rangeDescs []rangeDesc
}

// rangeDesc 标记下标处左右区间的数量
type rangeDesc struct {
	l int
	r int
}

func NewRange(nums []int) *Range {
	return &Range{
		nums:       nums,
		rangeDescs: make([]rangeDesc, len(nums)+1),
	}
}

func NewRangeWithRangeDescs(nums []int, rangeDescs []rangeDesc) *Range {
	for i := 0; i < len(rangeDescs); i++ {
		rangeDescs[i].l = 0
		rangeDescs[i].r = 0
	}
	return &Range{
		nums:       nums,
		rangeDescs: rangeDescs,
	}
}

func (r *Range) ACC(left, right int) {
	if left < 0 || right < 0 || left > right || left >= len(r.nums) {
		return
	}
	if right >= len(r.nums) {
		right = len(r.nums) - 1
	}
	r.rangeDescs[left+1].l++
	r.rangeDescs[right+1].r++
}

func (r *Range) Total() (total int) {
	var rCount int
	for i, v := range r.nums {
		rCount += r.rangeDescs[i+1].l
		rCount -= r.rangeDescs[i].r
		total += v * rCount
	}
	return total
}

func (r *Range) resume() {
	for i := 0; i < len(r.rangeDescs); i++ {
		r.rangeDescs[i].l = 0
		r.rangeDescs[i].r = 0
	}
}
