package mysync

import "testing"

// goos: darwin
// goarch: arm64
// pkg: mysync
// cpu: Apple M1 Pro
// BenchmarkNewBigStruct-10    	52686558	        22.21 ns/op	      48 B/op	       1 allocs/op
var bsBenchmarkNewBigStruct *BigStruct

func BenchmarkNewBigStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bsBenchmarkNewBigStruct = NewBigStruct("123e4567-e89b-12d3-a456-426614174000", 1, "TestName1")
	}
}

// goos: darwin
// goarch: arm64
// pkg: mysync
// cpu: Apple M1 Pro
// BenchmarkNewBigStructFromPool-10    	140792635	         8.279 ns/op	       0 B/op	       0 allocs/op
var bsBenchmarkNewBigStructFromPool *BigStruct

func BenchmarkNewBigStructFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bsBenchmarkNewBigStructFromPool = NewBigStructFromPool("123e4567-e89b-12d3-a456-426614174000", 1, "TestName1")
		bsBenchmarkNewBigStructFromPool.Release()
	}
}
