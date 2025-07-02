package testhandler_test

import (
	"errors"
	"log/slog"
	"reflect"
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

func TestTestHandler_First_Empty(t *testing.T) {
	h := testhandler.New(slog.LevelDebug)

	r, err := h.First()

	emptyRecord := slog.Record{}
	if !reflect.DeepEqual(r, emptyRecord) {
		t.Errorf("expected record to be empty")
	}

	if !errors.Is(err, testhandler.ErrEmpty) {
		t.Errorf("expected err to be ErrEmpty, got %v", err)
	}
}

func TestTestHandler_Last_Empty(t *testing.T) {
	h := testhandler.New(slog.LevelDebug)

	r, err := h.Last()

	emptyRecord := slog.Record{}
	if !reflect.DeepEqual(r, emptyRecord) {
		t.Errorf("expected record to be empty")
	}

	if !errors.Is(err, testhandler.ErrEmpty) {
		t.Errorf("expected err to be ErrEmpty, got %v", err)
	}
}

func TestTestHandler_Levels(t *testing.T) {
	h := testhandler.New(slog.LevelWarn)

	l := slog.New(h)

	l.Info("hello")

	if len(h.Records) != 0 {
		t.Errorf("expected no records, got %d", len(h.Records))
	}

	l.Warn("oh no")

	if len(h.Records) != 1 {
		t.Errorf("expected 1 record, got %d", len(h.Records))
	}

	l.Error("whoops")

	if len(h.Records) != 2 {
		t.Errorf("expected 2 records, got %d", len(h.Records))
	}
}
