package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Vanessa Fernandes" // Para chamar a variável dentro do html + name +
	str := fmt.Sprint(`
<!DOCTYPE html>
	<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
	</head>
	<body>
		<h1>`+
			name +
		`</h1>
	</body>
</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
	fmt.Println(str) //Imprime o template html
}
