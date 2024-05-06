package main

import (
	"fmt"
	"sort"
)

func incluso(n int, nums []int) bool {
	for _, elem := range nums {
		if elem == n { return true }
	}

	return false
}

func conta(n int, nums []int) int {
	cont := 0
	for _, elem := range nums {
		if n == elem { cont++ }
	}

	return cont
}

func main() {
	var n, num int

	nums := make([]int, 0)
	unicos := make([]int, 0)
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&num)
		nums = append(nums, num)
	}

	for _, elem := range nums {
		if !incluso(elem, unicos) {
			unicos = append(unicos, elem)
		}
	}

	sort.Ints(unicos)
	for _, elem := range unicos {
		fmt.Println(elem, "aparece", conta(elem, nums), "vez(es)")
	}
}