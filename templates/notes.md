## Entendendo templates

Template permite que possamos criar um documento e fazer com que dados dinamicos sejam unificados a este documento. Web templates nos permite personalizar os resultados para um usuário específico ou aplicação.

Um exemplo de template para envio de mensagem:
> Caro Sr. fulano,
>
> Está cansado de receber tipos de investimentos que não te ajudam a enriquecer?
>
> Nós temos uma série de videos que podem ajudá-lo nesse processo...

O template nesse caso ficaria:
> Caro {{Name}},
>
> Está cansado de receber tipos de investimentos que não te ajudam a enriquecer?
>
> Nós temos uma série de videos que podem ajudá-lo nesse processo...

## Maneiras de se criar templates

O jeito mais fácil, seria incluir o template em uma variável e fazer a substituição de determinados campos:

```golang
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
```

Não é o jeito recomendado de fazer, mas é possível.

O outro formato é usando o pacote "text/template" ou "html/template", criando um arquivo que precisa ser renderizado conforme os inputs recebidos:
```bash
cat <<< '<!DOCTYPE html>
	<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
	</head>
	<body>
		<h1>Hello</h1>
	</body>
</html>' > ./index.gohtml 
```

```golang
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println("Erro ao fazer o parser do arquivo!", err.Error())
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("erro ao exibir o template!", err.Error())
	}
```

Esse é o modo recomendado de fazer. O nome e extensão do template pode ser qualquer coisa, visto que se trata apenas de um arquivo. O pacote template cria um ponteiro para um template, é um container que guarda todos os templates.
