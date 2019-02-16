package main

import (
	"runtime"
	"runtime/debug"
	"sync"
	"testing"
)

func BenchmarkWithSyncPool(b *testing.B) {
	newFunc := func() interface{} {
		return make([]int, 2)
	}
	pool := sync.Pool{New: newFunc}
	debug.SetGCPercent(100)
	b.ResetTimer()
	b.ReportAllocs()
	for j := 0; j < b.N; j++ {
		for i := 0; i < 1000; i++ {
			v2 := pool.Get().([]int)
			v2[0] = i
			pool.Put(v2)
			runtime.GC()
		}
	}
}

func BenchmarkWithoutSyncPool(b *testing.B) {
	arr := make([]int, 2)
	debug.SetGCPercent(100)
	b.ResetTimer()
	b.ReportAllocs()
	for j := 0; j < b.N; j++ {
		for i := 0; i < 1000; i++ {
			arr[0] = i
			runtime.GC()
		}
	}
}
