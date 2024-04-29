package main

import "fmt"

func main() {
	var n, desloc int
	var textoCod, textoDecod string

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&textoCod)
		fmt.Scanln(&desloc)

		textoDecod = ""
		for _, carac := range textoCod {
			ascii := int(carac)
			ascii = (ascii - desloc - 65) % 26
			if ascii < 0 {
				ascii += 26
			}
			textoDecod += string(ascii + 65)
		}

		fmt.Println(textoDecod)
	}
}