package tasks

func OrderCategory(kopeck int) string {
	if kopeck < 0 {
		return "Invalid input"
	} else if kopeck < 10000 {
		return "small"
	} else if kopeck < 50000 {
		return "regular"
	}
	return "large"
}
