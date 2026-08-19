package main

import (
	"fmt"
	"math/rand"
	"time"

	lib "github.com/aoticombr/golang/lib"
	cp "github.com/aoticombr/golang/stringlist"
)

func Cabecalho() string {
	return "Nome | Idade | Dta Nasci | Logradouro | Bairro | Cidade"
}

func Linha(i int) string {
	Idade := SortearIdade(50)
	Data := time.Now()
	Data = Data.AddDate(0, 0, -Idade)
	linha := fmt.Sprintf("Nome %d | %d | %s | Rua %d | Bairro %d | Cidade %d",
		i, Idade, Data.Format("02/01/2006"), i, i, i)
	return linha
}

func SortearIdade(ate int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ate)
}

func criarArquivoGb(value int) {
	var arquivo cp.Strings
	arquivo.Delimiter = "\n"

	arquivo.Add(Cabecalho())

	for i := 0; i < 10000000*value; i++ {
		arquivo.Add(Linha(i))
	}

	err := lib.ByteToSaveFile("aula4.txt", arquivo.Byte())
	if err != nil {
		fmt.Println("erro ao gravar o arquivo", err)
	}
}

func main() {
	criarArquivoGb(2)
}
