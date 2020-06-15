package main

import "fmt"

type Veiculo struct {
	portas int
	cor    string
}

type Caminhonete struct {
	Veiculo
	quatroRodas bool
}

type Sedan struct {
	modeloLuxo bool
	veiculo    Veiculo
}

func main() {
	ranger := Caminhonete{
		quatroRodas: true,
		Veiculo: Veiculo{
			portas: 4,
			cor:    "preta",
		},
	}
	honda := Sedan{true, Veiculo{4, "grafite"}}

	fmt.Println(ranger)
	fmt.Println(honda)

	fmt.Println(ranger.cor)
	fmt.Println(honda.veiculo.cor)
}
