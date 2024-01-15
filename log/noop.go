package log

import (
	"context"
	"log/slog"
)

type noop struct{}

func (n noop) Enabled(context.Context, slog.Level) bool  { return false }
func (n noop) Handle(context.Context, slog.Record) error { return nil }
func (n noop) WithAttrs([]slog.Attr) slog.Handler        { return n }
func (n noop) WithGroup(string) slog.Handler             { return n }
