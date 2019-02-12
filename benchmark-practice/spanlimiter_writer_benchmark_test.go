package benchmarkpractice

import (
	"sync"
	"testing"

	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/spanstore/mocks"
)

//func BenchmarkSpanLimiter1(b *testing.B) { BenchmarSpanLimiterWriting1(b) }

func BenchmarSpanLimiterWriting1(b *testing.B) {
	w1 := &mocks.Writer{}
	spanLimitParam := SpanLimitOptions{
		maxSpansPerTrace: 1000000,
		cacheSize:        1000000,
	}

	slw, _ := NewSpanLimitWriter(w1, spanLimitParam)

	trace := model.TraceID{
		Low:  uint64(0),
		High: uint64(1),
	}
	span := &model.Span{
		TraceID: trace,
	}
	w1.On("WriteSpan", span).Return(nil)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		slw.WriteSpan(span)
	}
}

type Writer struct {
}

func (_m *Writer) WriteSpan(span *model.Span) error {
	return nil
}

func BenchmarSpanLimiterPrefillCacheWriting2(size int, numGoroutines int, spansEach int, b *testing.B) {
	w1 := &Writer{}
	var wg sync.WaitGroup
	wg.Add(numGoroutines * b.N)
	spanLimitParam := SpanLimitOptions{
		maxSpansPerTrace: 100000,
		cacheSize:        size,
	}

	slw, _ := NewSpanLimitWriter(w1, spanLimitParam)

	var spans []*model.Span
	spans = make([]*model.Span, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		trace := model.TraceID{
			Low:  uint64(i),
			High: uint64(i + 1),
		}
		spans[i] = &model.Span{
			TraceID: trace,
		}
	}

	//prefill
	for j := 0; j < size; j++ {
		spans[0].TraceID.Low = uint64(j)
		spans[0].TraceID.High = uint64(j + 1)
		slw.WriteSpan(spans[0])
	}

	b.ResetTimer()
	b.ReportAllocs()
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		for i := 0; i < numGoroutines; i++ {
			go func(sp *model.Span) {
				defer wg.Done()
				goroutine(slw, sp, spansEach)
			}(spans[i])
		}
	}

	wg.Wait()
	b.StopTimer()
}

func goroutine(slw *SpanLimitWriter, sp *model.Span, spansEach int) {
	for j := 75000; j < spansEach + 75000; j++ {
		sp.TraceID = model.TraceID{
			Low:  uint64(j),
			High: uint64(j + 1),
		}
		slw.WriteSpan(sp)
	}
}

//func Benchmark_baseline(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 56000, b) }

//func Benchmark10000r(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 10000, 56000, b) }
//func Benchmark100r(b *testing.B)   { BenchmarSpanLimiterPrefillCacheWriting2(150000, 100, 56000, b) }
//
//func Benchmark560000t(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 560000, b) }
//func Benchmark5600t(b *testing.B)   { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 5600, b) }
//
//func Benchmark1500000s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(1500000, 1000, 56000, b) }
//func Benchmark15000s(b *testing.B)   { BenchmarSpanLimiterPrefillCacheWriting2(15000, 1000, 56000, b) }
//func Benchmark100s(b *testing.B)     { BenchmarSpanLimiterPrefillCacheWriting2(100, 1000, 56000, b) }
//func Benchmark1s(b *testing.B)       { BenchmarSpanLimiterPrefillCacheWriting2(1, 1000, 56000, b) }

//func Benchmark1(b *testing.B) {BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b)}

func Benchmark1s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark10s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark11s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark12s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark13s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark14s(b *testing.B)  { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark15s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark16s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark17s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark18s(b *testing.B) { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1, b) }
//func Benchmark10(b *testing.B)     { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 10, b) }
//func Benchmark100(b *testing.B)       { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 100, b) }
//func Benchmark1000(b *testing.B)       { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 1000, b) }
//func Benchmark10000(b *testing.B)       { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 10000, b) }
//func Benchmark100000(b *testing.B)       { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 100000, b) }
//func Benchmark60000(b *testing.B)       { BenchmarSpanLimiterPrefillCacheWriting2(150000, 1000, 60000, b) }