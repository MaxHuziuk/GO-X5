package main

import "fmt"

func main() {
	order, err := NewOrder(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status:", statusNames[order.Status])
	for i := 1; i <= 10; i++ {
		item := Item{
			Name:         fmt.Sprintf("Item %d", i),
			PriceKopecks: int64(i * 100),
		}

		fmt.Printf(
			"before: len=%d cap=%d\n",
			len(order.Items),
			cap(order.Items),
		)

		if err := AddItem(&order, item); err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf(
			"after:  len=%d cap=%d\n\n",
			len(order.Items),
			cap(order.Items),
		)
	}
}
