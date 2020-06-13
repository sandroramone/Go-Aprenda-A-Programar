// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função,
// onde esta outra função faz uso de uma variável alem de seu scope.

package main

import "fmt"

func main() {
	x := intGenerator()
	y := intGenerator()

	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
}

func intGenerator() func() int {
	numero := -1
	return func() int {
		numero++
		return numero
	}
}
