package main

import "fmt"

func ExchangeCoin(amount int) []int {
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	result := make([]int, 0)

	for _, coin := range coins {
		for amount >= coin {
			result = append(result, coin)
			amount -= coin
		}
	}

	return result
}

func main() {
	fmt.Println(ExchangeCoin(1752))
	fmt.Println(ExchangeCoin(5000))
	fmt.Println(ExchangeCoin(1234))
	fmt.Println(ExchangeCoin(0))
}
