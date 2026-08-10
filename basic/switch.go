package basic

import "fmt"

func main() {
	posicao := 1

	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	default:
		fmt.Println("Não está no pódio")
	}
}
