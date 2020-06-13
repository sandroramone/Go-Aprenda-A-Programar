// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.

package main

import "fmt"

func main() {
	x := f()
	x()
}

func f() func() {
	return func() {
		fmt.Println("Ok!")
	}
}
