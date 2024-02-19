package main

import "fmt"

const MAX_ARRAY = 5

func main() {
	// declaração explícita
	var numeros [5]int
	var nomes [MAX_ARRAY]string

	fmt.Println(numeros)
	fmt.Println(nomes)

	numeros[2] = 5
	fmt.Println(numeros)

	// declaração curta
	outrosNumeros := [4]int{4, 8, 1, -5}
	fmt.Println(outrosNumeros)

	// percorrendo arrays
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	// equivalente ao enumerate
	for i, numero := range numeros {
		fmt.Println(i, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

	for i := range numeros {
		fmt.Println(i)
	}

	// slices
	var nums []int
	nums = numeros[:3]
	fmt.Println(nums)

	maisNumeros := []float64{1.4, 5.8, 4.56}
	fmt.Println(maisNumeros)

	exemplos := make([]int, 0)
	fmt.Println(len(exemplos), cap(exemplos))
	exemplos = append(exemplos, 1)
	fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))
	// exemplos = append(exemplos, 1)
	// fmt.Println(len(exemplos), cap(exemplos))

	fmt.Println(exemplos)
	alteraElemento(exemplos, 10)
	fmt.Println(exemplos)
}

func alteraElemento(slice []int, valor int) {
	slice[0] = valor
}
