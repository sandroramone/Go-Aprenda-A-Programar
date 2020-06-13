package main

import "fmt"

type Pessoa struct {
	Nome  string
	idade int
}

func main() {
	x := Pessoa{Nome: "sandro", idade: 32}

	fmt.Println(x)
	muda(&x)
	fmt.Println(x)
	muda2(x)
	fmt.Println(x)
}

func muda(p *Pessoa) {
	p.Nome = "Maria"
}

func muda2(p Pessoa) {
	p.Nome = "Sandro"
}
