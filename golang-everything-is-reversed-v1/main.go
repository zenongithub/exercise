package main

import "fmt"

func ReverseData(arr [5]int) [5]int {
	reversed := [5]int{}

	for i := 0; i < len(arr); i++ {
		reversed[i] = reverseNumber(arr[len(arr)-1-i])
	}

	return reversed
}

func reverseNumber(num int) int {
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return reversed
}

func main() {
	fmt.Println(ReverseData([5]int{123, 456, 11, 1, 2}))          
	fmt.Println(ReverseData([5]int{456789, 44332, 2221, 12, 10})) 
	fmt.Println(ReverseData([5]int{10, 10, 10, 10, 10}))          
	fmt.Println(ReverseData([5]int{23456, 789, 123, 456, 500}))   
	fmt.Println(ReverseData([5]int{0, 0, 0, 0, 0}))               
}
