// - Atribua uma função a uma variável.
// - Chame essa função.

package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("Função em variavel.")
	}

	x()
}
