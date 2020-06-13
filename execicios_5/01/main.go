package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func main() {
	pessoas := [2]Pessoa{
		Pessoa{"Sandro", "Rodrigo", []string{"passas ao rum", "ninho"}},
		Pessoa{"Maria", "Thais", []string{"pistache", "milho verde"}},
	}

	for _, v := range pessoas {
		fmt.Printf("%v %v\n", v.Nome, v.Sobrenome)
		for _, v := range v.Sabores {
			fmt.Printf("\t%v\n", v)
		}
	}
}
