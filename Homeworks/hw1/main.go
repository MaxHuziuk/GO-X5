package main

import (
	"fmt"
	"hw1/tasks"
)

func main() {
	tasks.FizzBuzz()
	fmt.Println()
	fmt.Println(tasks.OrderCategory(30000))
	fmt.Println()
	fmt.Println(tasks.SumsStatistics([]int{1337, 6767, 1984, 1111}))
}
