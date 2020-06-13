// - Callback: passe uma função como argumento a outra função.
package main

import "fmt"

func main() {
	x := func() { fmt.Println("Executou o callback!") }
	execfunction(x)
}

func execfunction(callback func()) {
	callback()
}
