package main

import "fmt"

func main() {
	x := make([]string, 26)
	x = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará",
		"Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
		"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná",
		"Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(x), cap(x))
	for y := 0; y < len(x); y++ {
		fmt.Println(x[y])
	}
}
