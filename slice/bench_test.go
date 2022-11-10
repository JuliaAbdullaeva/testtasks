package slice

import (
	"testing"
)

func BenchmarkIntersectionSlices(b *testing.B) {
	a := []int{1, 5, 3, 4, 2, 5, 56, 7, 9, 8, 10, 2}
	c := []int{2, 14, 4, 6, 56, 2}
	for i := 0; i < b.N; i++ {
		IntersectionSlices(a, c)
	}
}
