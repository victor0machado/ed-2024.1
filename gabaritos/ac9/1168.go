package main

import "fmt"

func main() {
	var n, leds int
	var numero string

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&numero)

		leds = 0
		for _, dig := range numero {
			switch dig {
			case '0':
				leds += 6
			case '1':
				leds += 2
			case '2':
				leds += 5
			case '3':
				leds += 5
			case '4':
				leds += 4
			case '5':
				leds += 5
			case '6':
				leds += 6
			case '7':
				leds += 3
			case '8':
				leds += 7
			case '9':
				leds += 6
			}
		}
		fmt.Println(leds, "leds")
	}
}
