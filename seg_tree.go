package a_range_struct

type SegmentRange struct {
	nums    []int
	segment []int
	total   int
}

func NewSegRange(nums []int) *SegmentRange {
	n := len(nums)
	re := &SegmentRange{
		nums:    nums,
		segment: make([]int, n*2),
	}

	// 构建线段树
	for i, j := n, 0; i < 2*n; i, j = i+1, j+1 {
		re.segment[i] = nums[j]
	}

	for i := n - 1; i > 0; i-- {
		re.segment[i] = re.segment[2*i] + re.segment[2*i+1]
	}

	return re
}

func (sr *SegmentRange) ACC(left, right int) (re int) {
	if left < 0 || right < 0 || left > right || left >= len(sr.nums) {
		return 0
	}
	if right >= len(sr.nums) {
		right = len(sr.nums) - 1
	}

	// 线段树求和
	n := len(sr.nums)
	l, r := left+n, right+n
	for l <= r {
		if l%2 == 1 {
			re += sr.segment[l]
			l++
		}
		if r%2 == 0 {
			re += sr.segment[r]
			r--
		}
		l /= 2
		r /= 2
	}
	sr.total += re
	return
}

func (sr *SegmentRange) Total() (total int) {
	return sr.total
}
