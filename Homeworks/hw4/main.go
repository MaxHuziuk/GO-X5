package main

import (
	"errors"
	"fmt"
	"hw4/notifiers"
)

func main() {
	defer func() {
		if value := recover(); value != nil {
			fmt.Println("main перехватил панику:", value)
		}
	}()

	errExample := notifiers.NewCustomError("very important info", notifiers.ErrEmptyMessage)

	fmt.Println("ошибка содержит ErrBadInput:", errors.Is(errExample, notifiers.ErrEmptyMessage))
	fmt.Println("ошибка содержит notifiers.ErrCustom с info \"very important info\":", errors.Is(errExample, notifiers.NewCustomError("very important info", notifiers.ErrEmptyMessage)))

	var err_custom *notifiers.ErrCustom
	if errors.As(errExample, &err_custom) {
		fmt.Println("ErrCustom найдена:", err_custom)
	}

	unwrappedErr := errors.Unwrap(errExample)
	fmt.Println("вложенная ошибка - ErrEmptyMessage:", errors.Is(unwrappedErr, notifiers.ErrEmptyMessage))

	notifierList := []notifiers.Notifier{
		notifiers.EmailNotifier{},
		notifiers.BrokenNotifier{},
		nil,
	}

	dispatcher := notifiers.NewDispatcher(notifierList)
	err := dispatcher.Send("GO X5")
	if err != nil {
		fmt.Println("При отправке произошли ошибки:")
		fmt.Println(err)
	}
	fmt.Println("прямой вызов без safeSend")
	notifiers.BrokenNotifier{}.Send("Заказ готов")
}
