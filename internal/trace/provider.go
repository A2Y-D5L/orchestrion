package trace

// Provider allows switching between different tracing backends
type Provider interface {
	NewTracer(config map[string]interface{}) (span.Tracer, error)
}
