package tasks

import "testing"

func TestOrderCategory(t *testing.T) {
	kopeck_inputs := []int{-10000, -100, -10, 0, 100, 1000, 10000, 25000, 49999, 50000, 100000}
	want_results := []string{"Invalid input", "Invalid input", "Invalid input",
		"small", "small", "small", "regular", "regular", "regular", "large", "large"}

	for i := range len(kopeck_inputs) {
		kopeck := kopeck_inputs[i]
		want := want_results[i]
		got := OrderCategory(kopeck)
		if want != got {
			t.Fatalf("OrderCategory(%d) = '%s' want '%s'", kopeck, got, want)
		}
	}

}
