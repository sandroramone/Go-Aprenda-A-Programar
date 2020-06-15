// - Nível 10?! Êita! Parabéns!
// - Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
//     - Usando uma função anônima auto-executável
//     - Usando buffer
// - Solução:
//     - 1. https://play.golang.org/p/MNqpJ29FZJ
//     - 2. https://play.golang.org/p/Y0Hx6IZc3U
package main

import "fmt"

func main() {
	// usando buffer
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

	// usando função
	f := make(chan int)
	go func() {
		f <- 42
	}()

	fmt.Println(<-f)
}
