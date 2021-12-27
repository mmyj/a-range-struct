package a_range_struct

type PrefixSumRange struct {
	nums      []int
	prefixSum []int
	total     int
}

func NewPrefixSumRange(nums []int) *PrefixSumRange {
	n := len(nums)
	re := &PrefixSumRange{
		nums:      nums,
		prefixSum: make([]int, n+1),
	}

	re.prefixSum[0] = 0
	for i := 0; i < n; i++ {
		re.prefixSum[i+1] = re.prefixSum[i] + nums[i]
	}

	return re
}

func (psr *PrefixSumRange) ACC(left, right int) int {
	if left < 0 || right < 0 || left > right || left >= len(psr.nums) {
		return 0
	}
	if right >= len(psr.nums) {
		right = len(psr.nums) - 1
	}
	re := psr.prefixSum[right+1] - psr.prefixSum[left]
	psr.total += re
	return re
}

func (psr *PrefixSumRange) Total() int {
	return psr.total
}
