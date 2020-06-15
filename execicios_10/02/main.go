// - Faça esse código funcionar: https://play.golang.org/p/oB-p3KMiH6
// - Solução: https://play.golang.org/p/isnJ8hMMKg
package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
