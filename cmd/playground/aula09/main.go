package main

import "fmt"

type Pessoa struct {
	Peso float32
	Cor  string
}

func Inc(value *int) *int {
	newValue := 3
	value = &newValue
	return value
}

func ModPessoaA(value *Pessoa) {
	value.Cor = "amarelo"
	value.Peso = 300
}

func ModPessoaB(value *Pessoa) *Pessoa {
	value.Cor = "amarelo"
	value.Peso = 300

	value = &Pessoa{Peso: 90, Cor: "Branca"}
	return value
}

func main() {

	var (
		// string, float32, int -> tipos primitivos
		v1 string
		v2 *string
		// struct: objeto com múltiplas propriedades
		v3 Pessoa
		v4 *Pessoa
	)
	v1 = "pedro"

	nome := "zezinho"
	v2 = &nome

	v3.Cor = "vermelho"
	v3.Peso = 1.68

	v4 = &Pessoa{
		Peso: 2.10,
		Cor:  "Branco",
	}
	fmt.Println("=======begin====")
	fmt.Println(fmt.Sprintf("v1 = %v", v1))
	fmt.Println("================")
	fmt.Println(fmt.Sprintf("v2 = %v", *v2))
	fmt.Println("================")
	fmt.Println(fmt.Sprintf("v3 = %v", v3))
	fmt.Println("================")
	fmt.Println(fmt.Sprintf("v3.Peso = %f", v3.Peso))
	fmt.Println("================")
	fmt.Println(fmt.Sprintf("v3.Cor = %s", v3.Cor))
	if v4 != nil {
		fmt.Println("================")
		fmt.Println(fmt.Sprintf("v4 = %v", v4))
		fmt.Println("================")
		fmt.Println(fmt.Sprintf("v4.Peso = %f", v4.Peso))
		fmt.Println("================")
		fmt.Println(fmt.Sprintf("v4.Cor = %s", v4.Cor))
	}
	fmt.Println("=======end======")

	var contador *int
	value := 0
	contador = &value
	fmt.Println(*contador)
	contador = Inc(contador)
	fmt.Println(*contador)
	Inc(contador)
	fmt.Println(*contador)
	Inc(contador)
	fmt.Println(*contador)
	Inc(contador)
	fmt.Println(*contador)
	Inc(contador)

	var pessoa *Pessoa
	pessoa = &Pessoa{}
	fmt.Println(pessoa)

	ModPessoaA(pessoa)
	fmt.Println(pessoa)
	pessoa = ModPessoaB(pessoa)
	fmt.Println(pessoa)
}
