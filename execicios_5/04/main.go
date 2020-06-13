package main

import "fmt"

func main() {
	x := struct {
		m map[string]string
		s []string
	}{
		m: map[string]string{
			"teste": "teste",
		},
		s: []string{"teste1", "teste2"},
	}

	fmt.Println(x)
}
