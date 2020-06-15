// - Esse exercício vai reforçar seus conhecimentos sobre conjuntos de métodos.
//     - Crie um tipo para um struct chamado "pessoa"
//     - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
//     - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
//     - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
//     - Demonstre no seu código:
//         - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
//         - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
// - Se precisar de dicas, veja: https://gobyexample.com/interfaces
// - Solução: 20_exercicios-ninja-9/02/main.go

package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func (p *Pessoa) falar() {
	fmt.Printf("Eu %s sou um humano!\n", p.nome)
}

type humanos interface {
	falar()
}

func falarAlgo(h humanos) {
	h.falar()
}

func main() {
	p := Pessoa{
		nome:  "Sandro",
		idade: 32,
	}

	falarAlgo(&p)
}
