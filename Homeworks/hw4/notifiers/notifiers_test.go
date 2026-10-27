package notifiers

import (
	"errors"
	"strings"
	"testing"
)

func TestBrokenNotifierPanics(t *testing.T) {
	defer func() {
		value := recover()

		if value == nil {
			t.Fatal("expected panic")
		}

		if value != "драйвер канала повреждён" {
			t.Errorf("unexpected panic: %v", value)
		}
	}()

	BrokenNotifier{}.Send("hello")
}

func TestSafeSendRecoversPanic(t *testing.T) {
	err := safeSend(BrokenNotifier{}, "hello")

	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), "драйвер канала повреждён") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCustomErrorUnwrap(t *testing.T) {
	innerErr := errors.New("inner error")
	customErr := NewCustomError("some info", innerErr)

	unwrappedErr := errors.Unwrap(customErr)

	if !errors.Is(unwrappedErr, innerErr) {
		t.Errorf("expected %v, got %v", innerErr, unwrappedErr)
	}
}

func TestCustomErrorIs(t *testing.T) {
	first := NewCustomError("same info", errors.New("first"))
	second := NewCustomError("same info", errors.New("second"))
	third := NewCustomError("different info", errors.New("third"))

	if !errors.Is(first, second) {
		t.Error("errors with the same info should match")
	}

	if errors.Is(first, third) {
		t.Error("errors with different info should not match")
	}
}

func TestCustomErrorContainsInnerError(t *testing.T) {
	innerErr := errors.New("inner error")
	customErr := NewCustomError("some info", innerErr)

	if !errors.Is(customErr, innerErr) {
		t.Errorf("expected error to contain %v", innerErr)
	}
}

func TestErrorsAs(t *testing.T) {
	customErr := NewCustomError(
		"very important info",
		ErrEmptyMessage,
	)

	err := errors.Join(
		errors.New("another error"),
		customErr,
	)

	var target *ErrCustom

	if !errors.As(err, &target) {
		t.Fatal("expected ErrCustom")
	}

	if target.info != "very important info" {
		t.Errorf("unexpected info: %s", target.info)
	}

	if target != customErr {
		t.Error("errors.As returned unexpected error")
	}
}

func TestDispatcherSend(t *testing.T) {
	dispatcher := NewDispatcher([]Notifier{
		EmailNotifier{},
	})

	err := dispatcher.Send("GO X5")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDispatcherCollectsErrors(t *testing.T) {
	dispatcher := NewDispatcher([]Notifier{
		EmailNotifier{},
		BrokenNotifier{},
		nil,
	})

	err := dispatcher.Send("GO X5")

	if err == nil {
		t.Fatal("expected errors")
	}

	if !errors.Is(err, ErrNilNotifier) {
		t.Error("expected ErrNilNotifier")
	}

	if !strings.Contains(err.Error(), "драйвер канала повреждён") {
		t.Error("expected panic error")
	}

	var customErr *ErrCustom

	if !errors.As(err, &customErr) {
		t.Fatal("expected ErrCustom")
	}

	if customErr.info != "Индекс нотифаера: 1" {
		t.Errorf("unexpected info: %s", customErr.info)
	}
}

func TestDispatcherWithoutNotifiers(t *testing.T) {
	dispatcher := NewDispatcher(nil)

	err := dispatcher.Send("GO X5")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
