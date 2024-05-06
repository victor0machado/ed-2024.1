package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 int

	for {
		fmt.Scanln(&x1, &y1, &x2, &y2)
		if x1 == 0 && x2 == 0 && y1 == 0 && y2 == 0 { break }

		if x1 == x2 && y1 == y2 {
			fmt.Println(0)
		} else if x1 == x2 || y1 == y2 {
			fmt.Println(1)
		} else if math.Abs(float64(x1 - x2)) == math.Abs(float64(y1 - y2)) {
			fmt.Println(1)
		} else {
			fmt.Println(2)
		}
	}
}
