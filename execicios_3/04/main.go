package main

import "fmt"

func main() {
	birthdate := 1988

	for {
		if birthdate > 2020 {
			break
		}

		fmt.Println(birthdate)
		birthdate++
	}
}
