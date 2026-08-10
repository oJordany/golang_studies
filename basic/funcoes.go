package basic

// import "fmt"

// func main() {
// 	somaDosValores := soma(1, 2)
// 	subtracaoDosValores := subtracao(1, 2)
// 	fmt.Println(somaDosValores, " ", subtracaoDosValores)
// 	nome1, nome2 := getNomeDuplicado("steph")
// 	fmt.Println(nome1, " ", nome2)
// 	nomeCompleto := PrintaNomeCompleto("steph", "silva")
// 	fmt.Println(nomeCompleto)
// }

// Função começando com letra minúscula:
// Função é PRIVADA
// Função só pode ser usada no próprio pacote

// Função começando com letra maiúscula:
// Função é PÚBLICA
// Função pode ser usada em qualquer lugar

func soma(a int, b int) int {
	return a + b
}

func subtracao(a int, b int) int {
	return a - b
}

func getNomeDuplicado(nome string) (string, string) {
	return nome, nome
}

func PrintaNomeCompleto(nome, sobrenome string) string {
	return nome + " " + sobrenome
}
