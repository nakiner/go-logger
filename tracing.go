package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func WithTracingFields(ctx context.Context, lg *zap.SugaredLogger) {
	if spanContext := trace.SpanContextFromContext(ctx); spanContext.IsValid() {
		lg = lg.With("trace_id", spanContext.TraceID().String())
		lg = lg.With("span_id", spanContext.SpanID().String())
	}
}
