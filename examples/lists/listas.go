package main

// LISTAS

// 1 - Arrays e Slices: Homogêneos
// todos os elementos tem o mesmo tipo
// [1,2,3,4,5,6] - []int
// ["steph", "silva", "bento"] - []string

// 2 - Maps: Heterogêneos
// pode misturar tipos diferentes
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro tipo
// map[string]int
// {"steph": 1, "silva": 2, "bento": 3}
// map[string]string
// {"steph": "silva", "bento": "sousa"}

// ARRAYS
// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamos os valoresa com índices: a[0], a[1], a[2], ...
// função embutida len() retorna o tamanho do array
// Por conta dopp tmaanho fixo, não é tão usado só em casos específicos

// Slice
// Tipo o Array, mas com tamanho dinâmico
// acessamos os valores com índices: s[0], s[1], s[2], ...
// função embutida len() retorna o tamanho do slice
// função append() usada para adicionar elementos ao slice

import "fmt"

func main() {
	// Array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(len(array))
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos[0:3])
	fmt.Println(len(numPrimos))

	// Slice - tamanho dinâmico
	// slice := []string{"Hello", "World"}
	// slice := []string{}
	slice := make([]string, 2) //cria um slice com tamanho 2
	slice[0] = "Hello"
	slice[1] = "World"
	slice = append(slice, "!", "How", "are", "you?")
	fmt.Println(slice[0], slice[1], slice[2])
	fmt.Println(slice)
	fmt.Println(len(slice))
}
