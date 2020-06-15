// - Crie um programa que demonstra seu OS e ARCH.
// - Rode-o com os seguintes comandos:
//     - go run
//     - go build
//     - go install
// - Solução: 20_exercicios-ninja-9/06/main.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("System: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
}
