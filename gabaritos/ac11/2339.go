package main

import "fmt"

func main() {
	var comp, qtd, qtd_pessoa int
	fmt.Scanln(&comp, &qtd, &qtd_pessoa)

	if qtd_pessoa * comp > qtd {
		fmt.Println("N")
	} else {
		fmt.Println("S")
	}
}
