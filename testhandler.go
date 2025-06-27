package testhandler

import (
	"context"
	"log/slog"
)

var _ slog.Handler = (*TestHandler)(nil)

type TestHandler struct {
	Level   slog.Level
	Records []slog.Record
}

func (t *TestHandler) Handle(_ context.Context, r slog.Record) error {
	t.Records = append(t.Records, r)
	return nil
}

func (t *TestHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= t.Level
}

func (t *TestHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return nil
}

func (t *TestHandler) WithGroup(name string) slog.Handler {
	return nil
}
