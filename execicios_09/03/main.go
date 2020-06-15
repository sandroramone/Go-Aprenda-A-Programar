// - Utilizando goroutines, crie um programa incrementador:
//     - Tenha uma variável com o valor da contagem
//     - Crie um monte de goroutines, onde cada uma deve:
//         - Ler o valor do contador
//         - Salvar este valor
//         - Fazer yield da thread com runtime.Gosched()
//         - Incrementar o valor salvo
//         - Copiar o novo valor para a variável inicial
//     - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//     - Demonstre que há uma condição de corrida utilizando a flag -race
// - Solução: 20_exercicios-ninja-9/03/main.go

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

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
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}(x)
	}
}
