package log

import (
	"context"
	"log/slog"
)

type contextKey struct{}

// IntoContext injects the given logger into context.
func IntoContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

// FromContext returns logger from the context or noop implementation.
func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(contextKey{}).(*slog.Logger)
	if !ok {
		return slog.New(noop{})
	}
	return logger
}

type correlationIdKey struct{}

// WithCorrelationId injects correlation id into the context.
func WithCorrelationId(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, correlationIdKey{}, id)
}

// CorrelationId retrieves correlation id from the context.
func CorrelationId(ctx context.Context) string {
	id, ok := ctx.Value(correlationIdKey{}).(string)
	if !ok {
		return ""
	}
	return id
}

type contextHandler struct {
	wrap slog.Handler
}

func (h contextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.wrap.Enabled(ctx, level)
}

func (h contextHandler) Handle(ctx context.Context, record slog.Record) error {
	id := CorrelationId(ctx)
	const correlationIdKey = "correlation_id"
	if id != "" {
		record.Add(slog.String(correlationIdKey, id))
	}
	return h.wrap.Handle(ctx, record)
}

func (h contextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.wrap.WithAttrs(attrs)
}

func (h contextHandler) WithGroup(name string) slog.Handler {
	return h.wrap.WithGroup(name)
}
