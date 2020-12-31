package main

import (
	"fmt"
)

type hotdog int

type person struct {
	fname string //visivel apenas na struct
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good Morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	x := 7
	xi := []int{2, 4, 7, 9, 42}
	m := map[string]int{
		"Vanessa": 36,
		"Job":     12,
	}

	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Println(xi)
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()

	//Polimorfismo
	saySomething(p1)
	saySomething(sa1)
}
