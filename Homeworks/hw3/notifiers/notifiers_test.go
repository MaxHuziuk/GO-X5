package notifiers

import (
	"errors"
	"testing"
)

type testNotifier struct {
	message string
	err     error
	called  bool
}

func (n *testNotifier) Send(message string) error {
	n.called = true
	n.message = message
	return n.err
}

func TestEmailNotifierSend(t *testing.T) {
	notifier := EmailNotifier{
		BaseNotifier: BaseNotifier{Channel: "email"},
		Address:      "student@example.com",
	}

	if err := notifier.Send("hello"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := notifier.Send(""); !errors.Is(err, ErrEmptyMessage) {
		t.Errorf("expected %v, got %v", ErrEmptyMessage, err)
	}
}

func TestPushNotifierSend(t *testing.T) {
	notifier := PushNotifier{
		BaseNotifier: BaseNotifier{Channel: "push"},
		DeviceID:     "device-42",
	}

	if err := notifier.Send("hello"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := notifier.Send(""); !errors.Is(err, ErrEmptyMessage) {
		t.Errorf("expected %v, got %v", ErrEmptyMessage, err)
	}
}

func TestDispatcherSend(t *testing.T) {
	email := &testNotifier{}
	push := &testNotifier{}

	dispatcher := Dispatcher{
		Notifiers: []Notifier{email, push},
	}

	err := dispatcher.Send("hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !email.called || email.message != "hello" {
		t.Errorf("email notifier received %q", email.message)
	}

	if !push.called || push.message != "hello" {
		t.Errorf("push notifier received %q", push.message)
	}
}

func TestDispatcherNilNotifier(t *testing.T) {
	dispatcher := Dispatcher{
		Notifiers: []Notifier{nil},
	}

	err := dispatcher.Send("hello")
	if !errors.Is(err, ErrNilNotifier) {
		t.Errorf("expected %v, got %v", ErrNilNotifier, err)
	}
}

func TestDispatcherStopsAfterError(t *testing.T) {
	sendErr := errors.New("send error")

	first := &testNotifier{err: sendErr}
	second := &testNotifier{}

	dispatcher := Dispatcher{
		Notifiers: []Notifier{first, second},
	}

	err := dispatcher.Send("hello")
	if !errors.Is(err, sendErr) {
		t.Errorf("expected %v, got %v", sendErr, err)
	}

	if second.called {
		t.Error("second notifier should not be called")
	}
}
