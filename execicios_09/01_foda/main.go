// - Alem da goroutine principal, crie duas outras goroutines.
// - Cada goroutine adicional devem fazer um print separado.
// - Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
// - Solução:
//     - 20_exercicios-ninja-9/01_moleza/main.go
//     - 20_exercicios-ninja-9/01_foda/main.go

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novagoroutine(2)
	wg.Wait()
}

func novagoroutine(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Goroutine: ", i+1)
			wg.Done()
		}(x)
	}
}
