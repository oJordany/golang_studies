package main

import (
	"fmt"
	funcoes "golang_studies/basic"
	steph "strings"
)

func main() {
	fmt.Println("Hello, world!")
	result := steph.Split("steph", "")
	fmt.Println(result)
	fmt.Println(funcoes.PrintaNomeCompleto("steph", "silva"))
	fmt.Println("EXERCICIO 01")
	fmt.Println(funcoes.MakeNegative(3))
	fmt.Println(funcoes.MakeNegative(-3))
	fmt.Println("EXERCICIO 02")
	fmt.Println(funcoes.Summation(2))
	fmt.Println(funcoes.Summation(8))
	fmt.Println("EXERCICIO 03")
	fmt.Println(funcoes.RepeatStr(3, "steph"))
	fmt.Println(funcoes.RepeatStr(5, "I"))
	println("EXERCICIO 04")
	fmt.Println(funcoes.CalculateYears(1))
	fmt.Println(funcoes.CalculateYears(2))
	fmt.Println(funcoes.CalculateYears(3))
	fmt.Println(funcoes.CalculateYears(4))
}
