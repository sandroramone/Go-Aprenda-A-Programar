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
	wg.Add(1)
	go func() {
		fmt.Println("Goroutine: 1")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Goroutine: 2")
		wg.Done()
	}()

	wg.Wait()
}
