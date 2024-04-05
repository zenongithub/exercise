package main

import (
	"fmt"
	"sort"
)

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	}

	count := make(map[int]int)

	for _, dates := range villager {
		for _, date := range dates {
			count[date]++
		}
	}

	var result []int
	for date, freq := range count {
		if freq == len(villager) {
			result = append(result, date)
		}
	}

	sort.Ints(result)

	return result
}

func main() {
	villager1 := [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}
	fmt.Println("Test case 1:")
	fmt.Println("Input:", villager1)
	fmt.Println("Output:", SchedulableDays(villager1))
}
