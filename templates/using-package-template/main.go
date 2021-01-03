package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println("Erro ao fazer o parser do arquivo!", err.Error())
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("erro ao exibir o template!", err.Error())
	}
}
