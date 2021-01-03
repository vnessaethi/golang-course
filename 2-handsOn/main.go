package main

import "fmt"

type person struct {
	person string
}

type secretAgent struct {
	agent string
}

func (p person) pSpeak() {
	fmt.Println("Pessoa: ", p.person)
}

func (s secretAgent) saSpeak() {
	fmt.Println("Secret Agent: ", s.agent)
}

func main() {
	p := person{
		"Vanessa",
	}

	s := secretAgent{
		"hidden",
	}

	p.pSpeak()
	s.saSpeak()
}
