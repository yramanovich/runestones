package log

import (
	"context"
	"io"
)

// Close wraps the basic Close method and logs error.
func Close(ctx context.Context, closer io.Closer, description string) {
	if err := closer.Close(); err != nil {
		FromContext(ctx).ErrorContext(ctx, description, "err", err)
	}
}
