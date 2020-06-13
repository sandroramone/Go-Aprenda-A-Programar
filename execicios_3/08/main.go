package main

import "fmt"

func main() {
	numero := 10
	switch {
	case numero > 10:
		fmt.Println("Maior que 10")
	case numero < 10:
		fmt.Println("Menor que 10")
	case numero == 10:
		fmt.Println("Nenhum nem outro")
	}
}
