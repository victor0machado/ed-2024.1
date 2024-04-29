package main

import (
	"fmt"
	"strings"
)

func main() {
	var d int
	var n string

	for {
		fmt.Scanln(&d, &n)
		if d == 0 && n == "0" { break }

		n = strings.Replace(n, string(d + 48), "", -1)
		n = strings.TrimLeft(n, "0")
		if n == "" { n = "0" }
		fmt.Println(n)
	}
}
