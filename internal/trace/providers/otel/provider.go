package otel

import (
	"github.com/DataDog/orchestrion/internal/trace"
	"github.com/DataDog/orchestrion/internal/trace/span"
)

type tracer struct {
	// TODO: otel implementation
}

func NewProvider() trace.Provider {
	return &traceProvider{}
}

type traceProvider struct{}

func (p *traceProvider) NewTracer(config map[string]interface{}) (span.Tracer, error) {
	// Initialize otel tracer
	return &tracer{}, nil
}
