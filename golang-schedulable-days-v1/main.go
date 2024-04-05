package main

import (
	"fmt"
	"sort"
)

func SchedulableDays(date1 []int, date2 []int) []int {

	sort.Ints(date1)
	sort.Ints(date2)

	commonDates := []int{}

	i, j := 0, 0
	for i < len(date1) && j < len(date2) {
		if date1[i] == date2[j] {
			commonDates = append(commonDates, date1[i])
			i++
			j++
		} else if date1[i] < date2[j] {
			i++
		} else {
			j++
		}
	}

	return commonDates
}

func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5}))
	fmt.Println(SchedulableDays([]int{11, 12, 13, 14, 15}, []int{5, 10, 12, 13, 20, 21}))
	fmt.Println(SchedulableDays([]int{2, 7, 12, 20, 21, 22}, []int{1, 3, 6, 10}))
	fmt.Println(SchedulableDays([]int{}, []int{}))
}
