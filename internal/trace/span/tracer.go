package trace

import "context"

// Tag represents a key-value pair for span metadata
type Tag struct {
	Key   string
	Value interface{}
}

// Option defines span creation options
type Option func(*SpanConfig)

type SpanConfig struct {
	Tags []Tag
}

// Span represents a single unit of work
type Span interface {
	Finish()
	SetTag(key string, value interface{})
	Context() context.Context
}

// Tracer defines the interface for creating and managing spans
type Tracer interface {
	Trace(name string, opts ...Option) (Span, context.Context)
	Extract(ctx context.Context) (Span, bool)
	Inject(ctx context.Context, span Span) context.Context
}
