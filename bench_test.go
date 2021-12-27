package a_range_struct

import "testing"

func BenchmarkRange(b *testing.B) {
	lengthList := []int{100000}
	for _, length := range lengthList {
		randNums := makeRandNums(length)
		randAccRanges := makeRandAccRanges(length)

		b.ResetTimer()
		r := NewRange(randNums)
		for i := 0; i < b.N; i++ {
			for _, acc := range randAccRanges {
				r.ACC(acc.l, acc.r)
			}
			r.Total()
		}
	}
}

func BenchmarkSegRange(b *testing.B) {
	lengthList := []int{100000}
	for _, length := range lengthList {
		randNums := makeRandNums(length)
		randAccRanges := makeRandAccRanges(length)

		b.ResetTimer()
		r := NewSegRange(randNums)
		for i := 0; i < b.N; i++ {
			for _, acc := range randAccRanges {
				r.ACC(acc.l, acc.r)
			}
			r.Total()
		}
	}
}
