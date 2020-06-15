package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func (p Pessoa) apresentacao() {
	fmt.Printf(
		"Meu nome é %v %v e tenho %v anos!",
		p.Nome,
		p.Sobrenome,
		p.Idade,
	)
}

func main() {
	s := Pessoa{
		Nome:      "Sandro",
		Sobrenome: "Rodrigo",
		Idade:     32,
	}

	s.apresentacao()
}
