package main

import "fmt"

/* Structs
Forma de cirar usa própria estrutura de dados
Personalizar de acordo com a sua necessidade
Podemos usar vários tipos diferentes
*/

type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	pessoa1 := Pessoa{"steph", 20}
	pessoa2 := Pessoa{Nome: "bento", Idade: 21}
	pessoa3 := Pessoa{Nome: "silva"}

	println(pessoa1.Nome, pessoa1.Idade)
	println(pessoa2.Nome, pessoa2.Idade)
	println(pessoa3.Nome, pessoa3.Idade)

	pessoas := []Pessoa{}
	pessoas = append(pessoas, pessoa1, pessoa2, pessoa3)
	fmt.Println(pessoas)

	alunos := map[string][]Pessoa{}
	alunos["programacao"] = pessoas
	fmt.Println(alunos)

	prof := Profissao{pessoa2, "dev"}
	fmt.Println(prof.Pessoa.Nome, prof.Pessoa.Idade, prof.Tipo)
	fmt.Println(prof)
}
