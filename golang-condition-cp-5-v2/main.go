package main

import "fmt"

func TicketPlayground(height, age int) int {

	if age < 5 {
		return -1
	}

	var HargaTiket int

	switch {
	case age > 12:
		HargaTiket = 100000
	case height > 160 || age == 12:
		HargaTiket = 60000
	case height > 150 || age == 11 || age == 10:
		HargaTiket = 40000
	case height > 135 || age == 9 || age == 8:
		HargaTiket = 25000
	case height > 120 || age == 7 || age == 6 || age == 5:
		HargaTiket = 15000
	}
	return HargaTiket // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
