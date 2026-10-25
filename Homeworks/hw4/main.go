package main

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

var ErrBadInput error = errors.New("bad input")

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
			err = fmt.Errorf("паника при отправке: %v", value)
		}
	}()

	return notifier.Send(message)
}

func (e ErrCustom) Error() string {
	return fmt.Sprintf("info is %v, error is %v", e.info, e.err)
}

func (e ErrCustom) Unwrap() error {
	return e.err
}

func (e ErrCustom) Is(target error) bool {
	t, ok := target.(*ErrCustom)
	if !ok {
		return false
	}
	return (e.info == t.info)
}

func do() error {
	return ErrCustom{info: "my info", err: ErrBadInput}
}

func main() {
	defer func() {
		if value := recover(); value != nil {
			fmt.Println("main перехватил панику:", value)
		}
	}()

	err := do()

	fmt.Println("ошибка содержит ErrBadInput:", errors.Is(err, ErrBadInput))

	fmt.Println("ошибка содержит ErrCustom с info \"my info\":", errors.Is(err, &ErrCustom{info: "my info"}))

	multErr := errors.Join(
		err,
		errors.New("dsfssf"),
		errors.New("dsqwewqe"),
	)
	fmt.Println(multErr)

	notifiers := []Notifier{
		EmailNotifier{},
		BrokenNotifier{},
	}

	for _, notifier := range notifiers {
		err := safeSend(notifier, "Заказ готов")
		fmt.Println("safeSend:", err)
	}

	fmt.Println("прямой вызов без safeSend")
	BrokenNotifier{}.Send("Заказ готов")
}
