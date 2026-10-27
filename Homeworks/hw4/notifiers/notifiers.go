package notifiers

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

var ErrEmptyMessage error = errors.New("empty message")
var ErrNilNotifier error = errors.New("notifier is nil")

type ErrCustom struct {
	info string
	err  error
}

type EmailNotifier struct{}

type BrokenNotifier struct{}

func (EmailNotifier) Send(message string) error {
	fmt.Println("email отправлен:", message)
	return nil
}

func (BrokenNotifier) Send(string) error {
	panic("драйвер канала повреждён")
}

func safeSend(notifier Notifier, message string) (err error) {
	defer func() {
		if value := recover(); value != nil {
			if panicErr, ok := value.(error); ok {
				err = fmt.Errorf("паника при отправке: %w", panicErr)
				return
			}

			err = fmt.Errorf("паника при отправке: %v", value)
		}
	}()

	return notifier.Send(message)
}

func (e *ErrCustom) Error() string {
	return fmt.Sprintf("info is %v, error is %v", e.info, e.err)
}

func (e *ErrCustom) Unwrap() error {
	return e.err
}

func (e *ErrCustom) Is(target error) bool {
	t, ok := target.(*ErrCustom)
	if !ok {
		return false
	}
	return (e.info == t.info)
}

func NewCustomError(info string, err error) *ErrCustom {
	return &ErrCustom{info: info, err: err}
}

type Dispatcher struct {
	notifiers []Notifier
}

func NewDispatcher(notifiers []Notifier) Dispatcher {
	return Dispatcher{notifiers: notifiers}
}

func (d Dispatcher) Send(message string) error {
	var allErrors []error
	for idx, notifier := range d.notifiers {
		if notifier == nil {
			allErrors = append(allErrors, NewCustomError(fmt.Sprintf("Индекс нотифаера: %d", idx), ErrNilNotifier))
			continue
		}
		err := safeSend(notifier, message)
		if err != nil {
			allErrors = append(allErrors, NewCustomError(fmt.Sprintf("Индекс нотифаера: %d", idx), err))
		}
	}
	return errors.Join(allErrors...)
}
