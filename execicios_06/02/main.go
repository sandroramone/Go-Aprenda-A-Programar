package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(soma())
	fmt.Println(soma(numeros...))
	fmt.Println(somaSlice(numeros))
}

func soma(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func somaSlice(n []int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}
