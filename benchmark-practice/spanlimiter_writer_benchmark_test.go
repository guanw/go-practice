package benchmarkpractice

import (
	"testing"

	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/spanstore/mocks"
)

func BenchmarkSpanLimiter1(b *testing.B) { BenchmarSpanLimiterWriting1(b) }

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
