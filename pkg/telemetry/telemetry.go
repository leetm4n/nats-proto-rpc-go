package telemetry

import (
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Options struct {
	Tracer     trace.Tracer
	Propagator propagation.TextMapPropagator
}
