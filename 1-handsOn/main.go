package main

import (
	"fmt"
	"math"
)

//struct para guardar os dados do quadrado (square)
type square struct {
	square float64
}

//Função que vai usar a struct square para exibir o valor do quadrado
func (s square) area() float64 {
	return s.square * s.square
}

//struct para guardar os dados do circulo (circle)
type circle struct {
	circle float64
}

//Função que vai usar a struct circle para exibir o valor do circulo
func (c circle) area() float64 {
	return math.Pi * c.circle * c.circle
}

//Interface é uma coleção de métodos, um valor pode ser representado em multiplos (types) - polymorphism (é o princípio pelo qual duas ou mais classes derivadas de uma mesma superclasse podem invocar métodos que têm a mesma identificação (assinatura) mas comportamentos distintos, especializados para cada classe derivada, usando para tanto uma referência a um objeto do tipo da superclasse)
type shape interface {
	area() float64
}

//Função que imprime o calculo da area usando a interface shape
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		9,
	}

	c1 := circle{
		4,
	}

	info(s1)
	info(c1)
}
