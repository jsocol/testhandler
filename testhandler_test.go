package testhandler_test

import (
	"log/slog"
	"testing"

	"github.com/jsocol/testhandler"
)

func TestTestHandler(t *testing.T) {
	h := testhandler.New(slog.LevelInfo)
	l := slog.New(h)
	l.Info("hi")

	if len(h.Records) != 1 {
		t.Errorf("expected h.Records to have length 1, got %d", len(h.Records))
	}

	r, err := h.First()
	if err != nil {
		t.Errorf("got error %v", err)
	}

	if r.Message != "hi" {
		t.Errorf(`expected message "hi", got %q`, r.Message)
	}
}

func TestTestHandler_Reset(t *testing.T) {
	h := testhandler.New(slog.LevelInfo)
	l := slog.New(h)
	l.Info("hi")

	if len(h.Records) != 1 {
		t.Errorf("expected h.Records to have length 1, got %d", len(h.Records))
	}

	h.Reset()

	if len(h.Records) != 0 {
		t.Errorf("expected h.Records to have length 0, got %d", len(h.Records))
	}
}
