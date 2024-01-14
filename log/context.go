package log

import (
	"context"
	"log/slog"
	"os"
)

type contextKey struct{}

func IntoContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(contextKey{}).(*slog.Logger)
	if !ok {
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	return logger
}
