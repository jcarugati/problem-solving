package golang

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisitAllPlacesMinimumTime(t *testing.T) {
	res := VisitAllPlacesMinimumTime([]int{1, 2, 3, 4, 3, 4, 3, 4, 3, 4})
	assert.Equal(t, res, 4)
}

func benchmarkVisitAllPlacesMinimumTime(input []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		VisitAllPlacesMinimumTime(input)
	}
}

func BenchmarkVisitAllPlacesMinimumTime1(b *testing.B) {
	benchmarkVisitAllPlacesMinimumTime([]int{1, 4, 52, 3, 5, 2, 35, 12, 1, 5, 12, 31, 4, 5, 1, 2, 31, 23, 12, 5, 6, 3, 45, 4, 2, 4, 14, 123}, b)
}

func BenchmarkVisitAllPlacesMinimumTime2(b *testing.B) {
	var arr []int
	max := 1000
	min := 1
	for i := 0; i < 100; {
		arr = append(arr, rand.Intn((max-min)+min))
		i++
	}
	benchmarkVisitAllPlacesMinimumTime(arr, b)
}

func BenchmarkVisitAllPlacesMinimumTime3(b *testing.B) {
	var arr []int
	max := 1000
	min := 1
	for i := 0; i < 1000; {
		arr = append(arr, rand.Intn((max-min)+min))
		i++
	}
	benchmarkVisitAllPlacesMinimumTime(arr, b)
}

