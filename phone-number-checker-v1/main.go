package main

import (
	"fmt"
	"strings"
)

func PhoneNumberChecker(number string, result *string) {

	if !(strings.HasPrefix(number, "62") && len(number) >= 11) && !(strings.HasPrefix(number, "08") && len(number) >= 10) {
		*result = "invalid"
		return
	}

	if strings.HasPrefix(number, "62") {
		number = "0" + strings.TrimPrefix(number, "62")
	}

	prefix := number[:4]

	if prefix >= "0811" && prefix <= "0815" {
		*result = "Telkomsel"
	} else if prefix >= "0816" && prefix <= "0819" {
		*result = "Indosat"
	} else if prefix >= "0821" && prefix <= "0823" {
		*result = "XL"
	} else if prefix >= "0827" && prefix <= "0829" {
		*result = "Tri"
	} else if prefix >= "0852" && prefix <= "0853" {
		*result = "AS"
	} else if prefix >= "0881" && prefix <= "0888" {
		*result = "Smartfren"
	} else {
		*result = "invalid"
	}

}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "6285236188898"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
