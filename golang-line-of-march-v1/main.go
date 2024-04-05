package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {

	sort.Ints(height)

	return height
}

func main() {
	fmt.Println(Sortheight([]int{172, 170, 150, 165, 144, 155, 159}))
	fmt.Println(Sortheight([]int{155, 156, 160, 161, 162, 165, 170, 172}))
	fmt.Println(Sortheight([]int{180, 177, 175, 173, 170, 166, 165, 160}))
	fmt.Println(Sortheight([]int{}))
}
