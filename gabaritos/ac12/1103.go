package main

import "fmt"

func main() {
	var h1, m1, h2, m2 int
	for {
		fmt.Scanln(&h1, &m1, &h2, &m2)
		if h1 == 0 && m1 == 0 && h2 == 0 && m2 == 0 { break }

		minutos1 := h1 * 60 + m1
		minutos2 := h2 * 60 + m2

		min := minutos2 - minutos1
		if min < 0 {
			min += 24 * 60
		}

		fmt.Println(min)
	}
}
