package main

import (
	"fmt"
	"hw3/notifiers"
)

func main() {
	notifier_slice := []notifiers.Notifier{
		notifiers.EmailNotifier{Channel: "8000", Address: "student@example.com"},
		notifiers.PushNotifier{Channel: "8000", DeviceID: "device-42"},
	}

	dispatcher := notifiers.Dispatcher{Notifiers: notifier_slice}
	err := dispatcher.Send("GO X5")
	if err != nil {
		fmt.Printf("При отправке возникла ошибка, %v", err)
		return
	}
	fmt.Printf("Сообщения успешно отправлены")
}
