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

Outro formato usando ParseGlob, onde você consegue fazer o parser de mais arquivos:
```golang
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "file1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "file2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "file3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
```

- Nesse cenário, criamos uma variável que será como um container e que conterá todos os templates que estivermos analisando (parsing). Basicamente lê os dados de um tipo de arquivo especificado e faz a construção de um template na memória para facilitar a execução de algum tipo de transformação de dados.

- A função init permite executar alguma função antes do main.go

- template.Must é um auxiliar que envolve uma chamada de uma função que retorna um template e erro se for diferente de nulo (nil)

- template.ParseGlob cria um novo template e analisa as definições dos arquivos identificados pelo padrão, os arquivos correspondem ao filepath.Match correspondendo pelo menos a um arquivo.

Passando valores para um template:
```bash
cat <<< '<!DOCTYPE html>

<head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
</head>

<body>
    <h1>The meaning of life is: {{.}}</h1> <!-- {{.}} é o argumento que está sendo passado na execução do template, pode ser qualquer tipo, string, float, etc-->
</body>

</html>' > ./index.gohtml
```

```golang
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "Van")
	if err != nil {
		log.Fatalln(err)
	}
}
```

Usando variáveis:
```html
<!DOCTYPE html>

<head>
    <meta charset="UTF-8">
    <title>Hello World!</title>
</head>

<body>

    {{$wisdom := .}} <!--Declaro uma variável no meu template e consigo usá-la em qualquer lugar -->
    <h1>The meaning of life is: {{$wisdom}}</h1>
</body>

</html>
```

Passando uma estrutura de dados compostas/agregadas para os templates (slices, structs, etc):

Using range:
```html
<!DOCTYPE html>
<head>
    <meta charset="UTF-8">
    <title>My Peeps</title>
</head>
<body>
    <ul>
        {{range .}}
        <li>{{.}}</li>
        {{end}}
    </ul>
</body>
</html>
```

Usando range e variables:
```html
        {{range $index, $element := .}}
        <li>{{$index}} - {{$element}}</li>
        {{end}}
```

Iterando em struct-slice-struct:
```html
<ul>
    {{range .Wisdom}}
    <li>{{.Name}} - {{.Motto}}</li>
    {{end}}
</ul>

<ul>
    {{range .Transport}}
    <li>{{.Manufacturer}} - {{.Model}} - {{.Doors}}</li>
    {{end}}
</ul>

```
Using go fmt para formatar o código (linter):
```bash
$ go fmt ./...
```

Funções em templates:
```golang
type FuncMap map[string]interface{}
//É um tipo de mapa definindo o mapeamento de nomes para funções. Cada função retorna um ou mais valores e erro. FuncMap tem a mesma base de tipo que o pacote "text/template", então você não precisaria importar o pacote.

var fm = template.FuncMap{
		"uc": strings.ToUpper,
		"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}
//Nesse caso uso a função "Funcs" para passar a função "fm" para o template e executo o template dessa forma, passando o arquivo html:
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
```

No html passo as variáveis necessárias:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Functions</title>
</head>
<body>

{{range .Wisdom}}
{{uc .Name}}
{{end}}

{{range .Wisdom}}
{{ft .Name}}
{{end}}

</body>
</html>
```

Atenção: Precisamos entender como o compilador chama tudo isso.
