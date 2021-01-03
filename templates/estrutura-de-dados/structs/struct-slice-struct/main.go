package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// type items struct {
// 	Wisdom []sage
// 	Transport []car
// }

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	ford := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	corolla := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{buddha, gandhi, mlk}
	cars := []car{ford, corolla}

	// data := items{
	// 	Wisdom: sages,
	// 	Transport: cars,
	// }

	// Other way to create and call a struct, then I only declare a struct here:
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
