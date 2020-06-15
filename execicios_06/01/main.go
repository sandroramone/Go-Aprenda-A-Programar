package main

import "fmt"

func returnInt() int {
	return 5
}

func returnIntString() (int, string) {
	return 5, "nome"
}

func main() {
	fmt.Println(returnInt())
	fmt.Println(returnIntString())
}
