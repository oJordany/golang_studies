package main

import (
	"fmt"
	"os"
	"github.com/aoticombr/golang/stringlist"
)

func Soma(a, b int) int {
	return a + b
}

func Dividir(a, b int) float32 {
	if b == 0 {
		return 0
	}
	return float32(a) / float32(b)
}

// LerArquivo
//
// essa função lê arquivos que você indicar
func LerArquivo(nome string) (arq *os.File, err error) {
	arquivo, err := os.Open(nome)
	if err != nil {
		return nil, err
	}
	return arquivo, nil
}

func main() {
	var texto stringlist.Strings
	texto.Add("asasdasdas")
	texto.Add("asdasd")
	fmt.Println(texto.Text())

	a := 2
	b := 5
	fmt.Printf("A soma de %d + %d é igual a %d\n", a, b, Soma(a, b))
	fmt.Printf("A divisao de %d / %d é igual a %.2f\n", b, a, Dividir(b, a))

	teste := "teste "

	defer func(x string) {
		fmt.Println("Função anônima " + x + teste)
	}(teste)

	teste = "testando "

	fmt.Println("iiiiiiiiii")
	fmt.Println("hhhhhhhhhh")

	arq, err := LerArquivo("arquivo.csv")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo", err)
	}
	defer arq.Close() // Serve para executar por último antes de fechar o programa
}
