package testhandler

import (
	"context"
	"errors"
	"log/slog"
)

var ErrEmpty = errors.New("no records")

var _ slog.Handler = (*TestHandler)(nil)

type TestHandler struct {
	Level   slog.Level
	Records []slog.Record
	attrs   []slog.Attr
	group   string
}

func New(l slog.Level) *TestHandler {
	th := &TestHandler{
		Level: l,
	}
	th.Reset()
	return th
}

func (t *TestHandler) Handle(_ context.Context, r slog.Record) error {
	if r.Level >= t.Level {
		n := slog.NewRecord(r.Time, r.Level, r.Message, r.PC)

		c := r.Clone()
		c.AddAttrs(t.attrs...)
		c.Attrs(func(a slog.Attr) bool {
			key := a.Key
			if t.group != "" {
				key = t.group + "." + key
			}
			n.Add(key, a.Value)
			return true
		})

		t.Records = append(t.Records, n)
	}
	return nil
}

func (t *TestHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= t.Level
}

func (t *TestHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	nt := New(t.Level)
	nt.attrs = attrs
	return nt
}

func (t *TestHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return t
	}
	nt := New(t.Level)
	if t.group == "" {
		nt.group = name
	} else {
		nt.group = t.group + "." + name
	}

	return nt
}

func (t *TestHandler) Reset() {
	t.Records = make([]slog.Record, 0, 1)
}

func (t *TestHandler) First() (slog.Record, error) {
	if len(t.Records) == 0 {
		return slog.Record{}, ErrEmpty
	}
	return t.Records[0], nil
}

func (t *TestHandler) Last() (slog.Record, error) {
	if len(t.Records) == 0 {
		return slog.Record{}, ErrEmpty
	}
	return t.Records[len(t.Records)-1], nil
}
