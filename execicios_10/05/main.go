// - Utilizando este código: https://play.golang.org/p/YHOMV9NYKK
// - ...demonstre o comma ok idiom.
// - Solução: https://play.golang.org/p/qh2ywLB5OG
package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 10

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
