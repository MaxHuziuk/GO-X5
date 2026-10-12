package tasks

func SumsStatistics(input []int) (int, int, int, float64) {

	if len(input) == 0 {
		return 0, 0, 0, 0
	}

	var maximum int = input[0]
	var minimum int = input[0]
	var total_sum int = 0
	var mean float64 = 0

	for _, sum := range input {
		if minimum > sum {
			minimum = sum
		}
		if maximum < sum {
			maximum = sum
		}
		total_sum += sum
	}

	mean = float64(total_sum) / float64(len(input))
	return minimum, maximum, total_sum, mean

}
