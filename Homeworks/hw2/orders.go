package main

import (
	"errors"
	"strings"
)

var statusNames = makeStatusNames()

func makeStatusNames() map[string]string {
	names := make(map[string]string)

	names["new"] = "Order is created"
	names["processing"] = "Order is processing"
	names["completed"] = "Order is completed"
	names["cancelled"] = "Order is cancelled"

	return names
}

type Item struct {
	Name         string
	PriceKopecks int64
}

type Order struct {
	ID     int
	Items  []Item
	Status string
}

func NewOrder(id int) (Order, error) {
	if id <= 0 {
		return Order{}, errors.New("id must be positive")
	}
	return Order{ID: id, Status: "new"}, nil
}

func AddItem(order *Order, item Item) error {
	if order == nil {
		return errors.New("Order not specified")
	}
	if strings.TrimSpace(item.Name) == "" {
		return errors.New("Empty item name")
	}
	if item.PriceKopecks < 0 {
		return errors.New("Price must be non negative")
	}
	order.Items = append(order.Items, item)
	return nil
}

func SumOrders(order *Order) int64 {
	if order == nil {
		return 0
	}
	var result int64
	for _, item := range order.Items {
		result += item.PriceKopecks
	}
	return result
}

func CloneOrder(order *Order) *Order {
	if order == nil {
		return nil
	}

	cloned := *order
	cloned.Items = nil
	for _, item := range order.Items {
		cloned.Items = append(cloned.Items, item)
	}

	return &cloned
}
