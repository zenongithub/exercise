package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {

	HargaTiket := map[string]int{
		"VIP":     30,
		"Regular": 20,
		"Student": 10,
	}

	total := float32(VIP*HargaTiket["VIP"] + regular*HargaTiket["Regular"] + student*HargaTiket["Student"])

	if total >= 100 {

		var discount float32

		if day%2 == 0 {
			if VIP+regular+student < 5 {
				discount = 0.10
			} else {
				discount = 0.20
			}
		} else {
			if VIP+regular+student < 5 {
				discount = 0.15
			} else {
				discount = 0.25
			}
		}

		total -= total * discount
	}

	return total

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
