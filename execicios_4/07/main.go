package main

import "fmt"

func main() {
	x := [][]string{}

	x = append(x, []string{"Sandro", "Rodrigo", "Coding"})
	x = append(x, []string{"Maria", "Thais", "Ouvir Música"})
	x = append(x, []string{"Nicole", "Nunes", "Jogar no celular"})

	for _, v := range x {
		fmt.Println(v[0])
		for _, v := range v {
			fmt.Printf("\t%v\n", v)
		}
	}
}
