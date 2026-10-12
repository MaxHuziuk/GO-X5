package tasks

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	want_min, want_max, want_sum, want_mean := 0, 0, 0, float64(0)
	test_slice := []int{}
	got_min, got_max, got_sum, got_mean := SumsStatistics(test_slice)

	if want_max != got_max ||
		want_min != got_min ||
		want_sum != got_sum ||
		want_mean != got_mean {
		t.Fatalf("SumsStatistics(%v) = (%d, %d, %d, %g) want (%d, %d, %d, %g)",
			test_slice,
			got_max,
			got_min,
			got_sum,
			got_mean,
			want_max,
			want_min,
			want_sum,
			want_mean)

	}

}

func TestExample(t *testing.T) {
	want_min, want_max, want_sum, want_mean := 4500, 31000, 55000, float64(13750)
	test_slice := []int{12000, 7500, 31000, 4500}
	got_min, got_max, got_sum, got_mean := SumsStatistics(test_slice)

	if want_max != got_max ||
		want_min != got_min ||
		want_sum != got_sum ||
		want_mean != got_mean {
		t.Fatalf("SumsStatistics(%v) = (%d, %d, %d, %g) want (%d, %d, %d, %g)",
			test_slice,
			got_min,
			got_max,
			got_sum,
			got_mean,
			want_min,
			want_max,
			want_sum,
			want_mean)

	}
}
