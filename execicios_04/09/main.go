package main

import "fmt"

func main() {
	x := map[string][]string{
		"Sandro": []string{"Sandro", "Rodrigo", "Coding"},
		"Maria":  []string{"Maria", "Thais", "Ouvir Música"},
		"Nicole": []string{"Nicole", "Nunes", "Jogar no celular"},
	}

	x["Rabo"] = []string{"Rabo", "Fino", "Miar"}

	for i, v := range x {
		fmt.Println(i)
		for _, v := range v {
			fmt.Printf("\t%v\n", v)
		}
	}
}
