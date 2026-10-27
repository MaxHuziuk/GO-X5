package notifiers

import (
	"errors"
	"fmt"
)

type Notifier interface {
	Send(message string) error
}

type BaseNotifier struct {
	Channel string
}

var ErrEmptyMessage error = errors.New("empty message")
var ErrNilNotifier error = errors.New("notifier is nil")

type EmailNotifier struct {
	BaseNotifier
	Address string
}

func (n EmailNotifier) Send(message string) error {
	if len(message) == 0 {
		return ErrEmptyMessage
	}
	fmt.Printf("delivered email message: %s to address %s, in channel %s\n", message, n.Address, n.Channel)
	return nil
}

type PushNotifier struct {
	BaseNotifier
	DeviceID string
}

func (n PushNotifier) Send(message string) error {
	if len(message) == 0 {
		return ErrEmptyMessage
	}
	fmt.Printf("delivered push message: %s to device %s in channel %s\n", message, n.DeviceID, n.Channel)
	return nil
}

type Dispatcher struct {
	Notifiers []Notifier
}

func (d Dispatcher) Send(message string) error {
	for _, notifier := range d.Notifiers {
		if notifier == nil {
			return ErrNilNotifier
		}

		err := notifier.Send(message)
		if err != nil {
			return err
		}
	}
	return nil
}
