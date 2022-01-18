package main

/**
 go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/triumphpc/algorithms-templates/go/sprint0/t3
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkMovingAverage-16               44793999                26.76 ns/op           32 B/op          1 allocs/op
Benchmark_movingAverageBad-16           32875808                32.98 ns/op           32 B/op          1 allocs/op
PASS
*/
import (
	"testing"
)

func BenchmarkMovingAverage(t *testing.B) {
	result := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < t.N; i++ {
		movingAverage(result, 4)
	}
}

func Benchmark_movingAverageBad(t *testing.B) {
	result := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < t.N; i++ {
		movingAverageBad(result, 4)
	}
}
