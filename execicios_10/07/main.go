// - Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// - Tire estes números do canal e demonstre-os.
// - Solução: /22_exercicios-ninja-10/07/
package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup
var vezes int = 10
var total = vezes * 10

func main() {
	var counter int32 = 0
	c := make(chan int32)

	for i := 0; i < 10; i++ {
		envianumero(c, &counter)
	}

	for k := 0; k < total; k++ {
		fmt.Println(<-c)
	}

	fmt.Println("Counter value: ", counter)
	wg.Wait()
}

func envianumero(c chan int32, n *int32) {
	wg.Add(1)
	go func() {
		for i := 0; i < vezes; i++ {
			mutex.Lock()
			*n++
			c <- *n
			mutex.Unlock()
		}
		wg.Done()
	}()
}
