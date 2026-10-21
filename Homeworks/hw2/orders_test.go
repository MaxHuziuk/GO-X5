package main

import "testing"

func TestNewOrder(t *testing.T) {
	order, err := NewOrder(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if order.ID != 1 {
		t.Fatalf("expected ID 1, got %d", order.ID)
	}

	if SumOrders(&order) != 0 {
		t.Fatalf("expected sum 0, got %d", SumOrders(&order))
	}
}

func TestNewOrderInvalidID(t *testing.T) {
	_, err := NewOrder(0)

	if err == nil {
		t.Fatal("expected error for non-positive ID")
	}
}

func TestAddItem(t *testing.T) {
	order, err := NewOrder(1)
	if err != nil {
		t.Fatal(err)
	}

	err = AddItem(&order, Item{
		Name:         "Keyboard",
		PriceKopecks: 5000,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = AddItem(&order, Item{
		Name:         "Mouse",
		PriceKopecks: 2500,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(order.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(order.Items))
	}

	if SumOrders(&order) != 7500 {
		t.Fatalf("expected sum 7500, got %d", SumOrders(&order))
	}
}

func TestAddItemNilOrder(t *testing.T) {
	err := AddItem(nil, Item{
		Name:         "Keyboard",
		PriceKopecks: 5000,
	})

	if err == nil {
		t.Fatal("expected error for nil order")
	}
}

func TestAddItemEmptyName(t *testing.T) {
	order, err := NewOrder(1)
	if err != nil {
		t.Fatal(err)
	}

	err = AddItem(&order, Item{
		Name:         "",
		PriceKopecks: 5000,
	})

	if err == nil {
		t.Fatal("expected error for empty item name")
	}

	if len(order.Items) != 0 {
		t.Fatal("invalid item was added to order")
	}
}

func TestAddItemNegativePrice(t *testing.T) {
	order, err := NewOrder(1)
	if err != nil {
		t.Fatal(err)
	}

	err = AddItem(&order, Item{
		Name:         "Keyboard",
		PriceKopecks: -1,
	})

	if err == nil {
		t.Fatal("expected error for negative price")
	}

	if len(order.Items) != 0 {
		t.Fatal("invalid item was added to order")
	}
}

func TestCloneOrder(t *testing.T) {
	original, err := NewOrder(1)
	if err != nil {
		t.Fatal(err)
	}

	err = AddItem(&original, Item{
		Name:         "Keyboard",
		PriceKopecks: 5000,
	})
	if err != nil {
		t.Fatal(err)
	}

	cloned := CloneOrder(&original)

	cloned.ID = 2
	cloned.Status = "completed"
	cloned.Items[0].Name = "Mouse"
	cloned.Items = append(cloned.Items, Item{
		Name:         "Monitor",
		PriceKopecks: 20000,
	})

	if original.ID != 1 {
		t.Fatalf("original ID changed to %d", original.ID)
	}

	if original.Status != "new" {
		t.Fatalf("original status changed to %q", original.Status)
	}

	if original.Items[0].Name != "Keyboard" {
		t.Fatalf(
			"original item changed to %q",
			original.Items[0].Name,
		)
	}

	if len(original.Items) != 1 {
		t.Fatalf(
			"expected 1 original item, got %d",
			len(original.Items),
		)
	}

	if len(cloned.Items) != 2 {
		t.Fatalf(
			"expected 2 cloned items, got %d",
			len(cloned.Items),
		)
	}
}

func TestCloneNilOrder(t *testing.T) {
	if CloneOrder(nil) != nil {
		t.Fatal("expected nil clone for nil order")
	}
}

func TestSumNilOrder(t *testing.T) {
	result := SumOrders(nil)

	if result != 0 {
		t.Fatalf(
			"expected sum 0 for nil order, got %d",
			result,
		)
	}
}

func TestStatusNames(t *testing.T) {
	statuses := makeStatusNames()
	statuses["returned"] = "Order is returned"

	if statuses["new"] != "Order is created" {
		t.Fatalf("unexpected new status name: %q", statuses["new"])
	}

	if statuses["returned"] != "Order is returned" {
		t.Fatal("failed to write into status map")
	}
}
