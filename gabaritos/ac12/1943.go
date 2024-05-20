package main

import "fmt"

func main() {
	var k int
	posicoes := [7]int{1, 3, 5, 10, 25, 50, 100}

	fmt.Scanln(&k)
	for _, posicao := range posicoes {
		if k <= posicao {
			fmt.Println("Top", posicao)
			break
		}
	}

}
