package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/aoticombr/golang/lib"
)

type Pessoa struct {
	Nome       string
	Idade      int
	DtaNasc    string
	Logradouro string
	Bairro     string
	Cidade     string
}

func (p *Pessoa) Imprimir() {
	fmt.Printf("%+v\n", p)
}

func ConverteLinhaEmPessoa(value []string) *Pessoa {
	p := &Pessoa{
		Nome:       value[0],
		Idade:      lib.StrToInt(strings.TrimSpace(value[1])),
		DtaNasc:    value[2],
		Logradouro: value[3],
		Bairro:     value[4],
		Cidade:     value[5],
	}
	return p
}

func main() {
	file, err := os.Open("aula4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'
	reader.ReuseRecord = true // economiza memória
	reader.TrimLeadingSpace = true

	linha := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		linha++
		if linha == 1 {
			continue
		}
		pessoa := ConverteLinhaEmPessoa(record)
		pessoa.Imprimir()
	}
}
