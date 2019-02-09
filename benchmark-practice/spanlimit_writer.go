package benchmarkpractice

import (
	"sync"

	lru "github.com/hashicorp/golang-lru"
	"github.com/jaegertracing/jaeger/model"
	"github.com/jaegertracing/jaeger/storage/spanstore"
)

type condition int

const (
	notOverLimit condition = iota + 1
	atLimit
	overLimit
)

// SpanLimitOptions is the options for constructing a spanlimit writer
type SpanLimitOptions struct {
	maxSpansPerTrace int
	cacheSize        int
}

// SpanLimitWriter is something
type SpanLimitWriter struct {
	spanWriter spanstore.Writer
	//cache structure: trace ID -> number of spans saved (*int)
	traceIDtoSpanCount *lru.Cache
	maxSpansPerTrace   int
	lock               sync.Mutex
}

// NewSpanLimitWriter creates a SpanLimitWriter
func NewSpanLimitWriter(spanWriter spanstore.Writer, options SpanLimitOptions) (*SpanLimitWriter, error) {
	cache, err := lru.New(options.cacheSize)
	if err != nil {
		return nil, err
	}
	return &SpanLimitWriter{
		spanWriter:         spanWriter,
		traceIDtoSpanCount: cache,
		maxSpansPerTrace:   options.maxSpansPerTrace,
	}, nil
}

// WriteSpan calls WriteSpan on wrapped writer.
// When progressing down, it will stop writing to wrapped writer if span limits is hit in local cache.
func (w *SpanLimitWriter) WriteSpan(span *model.Span) error {
	cacheStateForSpan := w.updateCacheState(span)

	if cacheStateForSpan == notOverLimit {
		//return w.spanWriter.WriteSpan(span)
		return nil
	}

	if cacheStateForSpan == atLimit {

	}

	// span over limit is a normal business scenario we try to cover here, no need to return error
	return nil
}

// updateCacheState check if number of spans from same trace exceeds a pre-defined threshold
// and updates local cache traceIDtoSpanCount
func (w *SpanLimitWriter) updateCacheState(span *model.Span) condition {
	// writer wraps parallel processor which spawns several goroutines to process messages
	// so we will need to lock the localCache for get/set operations
	w.lock.Lock()
	defer w.lock.Unlock()
	val, found := w.traceIDtoSpanCount.Get(span.TraceID)
	if !found {
		w.traceIDtoSpanCount.Add(span.TraceID, 1)
		return notOverLimit
	}
	cacheVal := val.(int)

	// already hit limit
	if cacheVal == w.maxSpansPerTrace {
		return overLimit
	}

	cacheVal++
	w.traceIDtoSpanCount.Add(span.TraceID, cacheVal)
	// just hit limit
	if cacheVal == w.maxSpansPerTrace {
		return atLimit
	}
	// not over limit
	return notOverLimit
}
