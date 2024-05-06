package main

import "fmt"

func fatorial(n int) int {
	fat := 1
	for i := 1; i <= n; i++ {
		fat *= i
	}

	return fat
}

func main() {
	var n1, n2 int
	for {
		_, err := fmt.Scanln(&n1, &n2)
		if err != nil { break }

		fmt.Println(fatorial(n1) + fatorial(n2))
	}
}
