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

	// também pode ser criado com map[string]Pessoa{}
	pessoas2 := make(map[string]Pessoa)

	for _, v := range pessoas {
		pessoas2[v.Sobrenome] = v
	}

	for i, v := range pessoas2 {
		fmt.Println(i)
		fmt.Printf("\t%v %v\n", v.Nome, v.Sobrenome)
		for _, v := range v.Sabores {
			fmt.Printf("\t%v\n", v)
		}
	}
}
