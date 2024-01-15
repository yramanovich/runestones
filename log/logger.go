package log

import (
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

// New returns new configured slog.Logger instance.
func New(level string) *slog.Logger {
	var lvl slog.Level
	_ = lvl.UnmarshalText([]byte(level)) //nolint:errcheck

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.String(slog.TimeKey, a.Value.Time().Format(time.RFC822))
			}
			// Remove the directory from the source's filename.
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source) //nolint:errcheck
				source.File = filepath.Base(source.File)
			}
			return a
		},
	})

	return slog.New(contextHandler{wrap: h})
}
