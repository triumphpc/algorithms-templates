// Tests and benchmarks.
//
// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.

/*
	https://eli.thegreenplace.net/2022/faster-sorting-with-go-generics/

	~/Devel/go/src/github.com/triumphpc/algorithms-templates/go/sorting/bubble (main*) » go test -bench=.                                                                                                         triumphpc@MacBook-Pro-triumphpc
	goos: darwin
	goarch: amd64
	pkg: github.com/triumphpc/algorithms-templates/go/sorting/bubble
	cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
	BenchmarkSortStringInterface-16    	     142	   8158664 ns/op
	BenchmarkSortStringGeneric-16      	     175	   6475225 ns/op
	BenchmarkSortStringFunc-16         	     158	   7476973 ns/op
	BenchmarkSortStringFuncGen-16      	     158	   7584109 ns/op
	BenchmarkSortStructs-16            	     361	   3279718 ns/op
	BenchmarkSortFuncStructs-16        	     398	   3283640 ns/op
	BenchmarkSortFuncGenStructs-16     	     375	   3179053 ns/op
	PASS
	ok  	github.com/triumphpc/algorithms-templates/go/sorting/bubble	13.651s

	go build -o bubble.out
	go tool pprof -list bubbleSortInterface ./bubble.out cpui.out                                                       triumphpc@MacBook-Pro-triumphpc
	Total: 3.52s
	ROUTINE ======================== main.bubbleSortInterface in /Users/triumphpc/Devel/go/src/github.com/triumphpc/algorithms-templates/go/sorting/bubble/main/bubble.go
		1.18s      3.49s (flat, cum) 99.15% of Total
			.          .     30:
			.          .     31:func bubbleSortInterface(x sort.Interface) {
			.          .     32:	n := x.Len()
			.          .     33:	for {
			.          .     34:		swapped := false
		150ms      150ms     35:		for i := 1; i < n; i++ {
		730ms      2.86s     36:			if x.Less(i, i-1) {
		 90ms      270ms     37:				x.Swap(i, i-1)
			.          .     38:				swapped = true
			.          .     39:			}
			.          .     40:		}
		210ms      210ms     41:		if !swapped {
			.          .     42:			return
			.          .     43:		}
			.          .     44:	}
			.          .     45:}
			.          .     46:

*/

package main

import (
	"fmt"
	"sort"
	"testing"

	"golang.org/x/exp/slices"
)

func TestBubbleSort(t *testing.T) {
	for _, length := range []int{1, 2, 4, 6, 17, 32, 800} {
		testname := fmt.Sprintf("sort-len-%d", length)
		t.Run(testname, func(t *testing.T) {
			// Test that our bubble sort works by comparing it to the built-in sort.
			ss := makeRandomStrings(length)
			ss2 := slices.Clone(ss)
			ss3 := slices.Clone(ss)
			ss4 := slices.Clone(ss)

			sort.Strings(ss)
			bubbleSortInterface(sort.StringSlice(ss2))
			bubbleSortGeneric(ss3)
			bubbleSortFunc(ss4, func(a, b string) bool { return a < b })

			for i := range ss {
				if ss[i] != ss2[i] {
					t.Fatalf("strings mismatch at %d; %s != %s", i, ss[i], ss2[i])
				}
				if ss[i] != ss3[i] {
					t.Fatalf("generic mismatch at %d; %s != %s", i, ss[i], ss3[i])
				}
				if ss[i] != ss4[i] {
					t.Fatalf("generic mismatch at %d; %s != %s", i, ss[i], ss4[i])
				}
			}
		})
	}
}

const N = 1_000

func BenchmarkSortStringInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStrings(N)
		b.StartTimer()
		bubbleSortInterface(sort.StringSlice(ss))
	}
}

func BenchmarkSortStringGeneric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStrings(N)
		b.StartTimer()
		bubbleSortGeneric(ss)
	}
}

func BenchmarkSortStringFunc(b *testing.B) {
	lessFunc := func(a, b string) bool { return a < b }
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStrings(N)
		b.StartTimer()
		bubbleSortFunc(ss, lessFunc)
	}
}

func BenchmarkSortStringFuncGen(b *testing.B) {
	lessFunc := func(a, b string) bool { return a < b }
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStrings(N)
		b.StartTimer()
		bubbleSortFunc(ss, lessFunc)
	}
}

func TestStructSorts(t *testing.T) {
	ss := makeRandomStructs(200)
	ss2 := make([]*myStruct, len(ss))
	for i := range ss {
		ss2[i] = &myStruct{n: ss[i].n}
	}

	bubbleSortInterface(myStructs(ss))
	bubbleSortFunc(ss2, func(a, b *myStruct) bool { return a.n < b.n })

	for i := range ss {
		if *ss[i] != *ss2[i] {
			t.Fatalf("sortfunc mismatch at %d; %v != %v", i, *ss[i], *ss2[i])
		}
	}
}

func BenchmarkSortStructs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStructs(N)
		b.StartTimer()
		bubbleSortInterface(myStructs(ss))
	}
}

func BenchmarkSortFuncStructs(b *testing.B) {
	lessFunc := func(a, b *myStruct) bool { return a.n < b.n }
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStructs(N)
		b.StartTimer()
		bubbleSortFunc(ss, lessFunc)
	}
}

func BenchmarkSortFuncGenStructs(b *testing.B) {
	lessFunc := func(a, b *myStruct) bool { return a.n < b.n }
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		ss := makeRandomStructs(N)
		b.StartTimer()
		bubbleSortFuncGen(ss, lessFunc)
	}
}
