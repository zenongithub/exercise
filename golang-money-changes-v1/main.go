package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	totalPrice := 0
	for _, product := range products {
		totalPrice += product.Price + product.Tax
	}

	change := amount - totalPrice
	if change < 0 {
		return nil
	}

	denominations := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	changeList := []int{}

	for _, denom := range denominations {
		for change >= denom {
			changeList = append(changeList, denom)
			change -= denom
		}
	}

	return changeList
}

func main() {
	// Test cases
	testCases := []struct {
		Amount   int
		Products []Product
	}{
		{10000, []Product{{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300}}},
		{30000, []Product{{Name: "baju 1", Price: 10000, Tax: 1000}, {Name: "Sepatu", Price: 15550, Tax: 1555}}},
		{5500, []Product{{Name: "Baju", Price: 5000, Tax: 500}}},
	}

	for _, tc := range testCases {
		fmt.Printf("Input: %d, %v\n", tc.Amount, tc.Products)
		fmt.Println("Output:", MoneyChanges(tc.Amount, tc.Products))
	}
}
