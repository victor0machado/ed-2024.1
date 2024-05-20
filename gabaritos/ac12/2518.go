package main

import (
	"fmt"
	"math"
)

func main() {
	var n, h, l, c float64
	for {
		_, err := fmt.Scanln(&n)
		if err != nil { break }

		fmt.Scanln(&h, &c, &l)
		hip := math.Sqrt(math.Pow(h, 2.0) + math.Pow(c, 2.0))
		area := n * hip * l / 10000.0

		fmt.Printf("%.4f\n", area)
	}
}
