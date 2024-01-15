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
