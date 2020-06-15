// - Utilize mutex para consertar a condição de corrida do exercício anterior.
// - Solução: 20_exercicios-ninja-9/04/main.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

var contador int

const quantidade = 90000

func main() {
	novagoroutine(quantidade)
	wg.Wait()
	fmt.Println("Total goroutines: ", quantidade)
	fmt.Println("Contador: ", contador)
}

func novagoroutine(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			mutex.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mutex.Unlock()
			wg.Done()
		}(x)
	}
}
