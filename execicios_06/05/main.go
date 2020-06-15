// - Crie um tipo "quadrado"
// - Crie um tipo "círculo"
// - Crie um método "área" para cada tipo que calcule e retorne a área da figura
//     - Área do círculo: 2 * π * raio
//     - Área do quadrado: lado * lado
// - Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
// - Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
// - Crie um valor de tipo "quadrado"
// - Crie um valor de tipo "círculo"
// - Use a função "info" para demonstrar a área do "quadrado"
// - Use a função "info" para demonstrar a área do "círculo"

package main

import (
	"fmt"
	"math"
)

type square struct {
	lado float32
}

type circle struct {
	raio float32
}

func (s square) area() float32 {
	return s.lado * s.lado
}

func (c circle) area() float32 {
	return 2 * math.Pi * c.raio
}

type figura interface {
	area() float32
}

func info(f figura) float32 {
	return f.area()
}

func main() {
	q := square{lado: 15}
	c := circle{raio: 5.25}

	fmt.Println(info(q))
	fmt.Println(info(c))
}
