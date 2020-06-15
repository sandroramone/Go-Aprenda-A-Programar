package main

import (
	"fmt"
)

const (
	x = iota + 2021
	y
	w
	z
)

func main() {
	fmt.Printf("%v\t%v\t%v\t%v\n", x, y, w, z)
}
