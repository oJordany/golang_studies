package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/aoticombr/golang/lib"
)

func ProcessarArqComWorkers(nomeArquivo string, numWorkers int) error {
	fmt.Println("Iniciando Processamento")
	PortaChans := make([]chan string, numWorkers)

	var espera sync.WaitGroup //ponto
	/*
		Contrataçao de trabalhadores
	*/
	fmt.Println("Iniciando trabalhandores(workers)")
	for w := 0; w < numWorkers; w++ {
		fmt.Println("[" + lib.IntToStr(w) + "]Iniciando trabalhandores(workers) - 1")
		PortaChans[w] = make(chan string, 100)
		fmt.Println("[" + lib.IntToStr(w) + "]Iniciando trabalhandores(workers) - 2")
		espera.Add(1) //batendo o ponto de trabalho
		fmt.Println("[" + lib.IntToStr(w) + "]Iniciando trabalhandores(workers) - 3")
		go func(workerId int, ch <-chan string) {
			fmt.Println("[" + lib.IntToStr(workerId) + "]go trabalhandores(workers) - 3.1")
			defer espera.Done()
			fmt.Println("[" + lib.IntToStr(workerId) + "]go trabalhandores(workers) - 3.2")
			outputfile, err := os.Create(fmt.Sprintf("work%d.txt", workerId))
			if err != nil {
				fmt.Printf("["+lib.IntToStr(workerId)+"]erro ao criar arquivo do worker %d: %v\n", workerId+1, err)
				return
			}
			fmt.Println("[" + lib.IntToStr(workerId) + "]go trabalhandores(workers) - 3.3")
			defer outputfile.Close()
			fmt.Println("[" + lib.IntToStr(workerId) + "]go trabalhandores(workers) - 3.4")
			writer := bufio.NewWriter(outputfile)

			for linha := range ch {
				fmt.Println("[" + lib.IntToStr(workerId) + "]go[T]  trabalhandores(workers) - 3.4.1")
				writer.WriteString(linha + "\n")
			}

			fmt.Println("[" + lib.IntToStr(w) + "]Finalizado trabalhandores(workers) - 3.5")
			writer.Flush()
			fmt.Println("[" + lib.IntToStr(w) + "]Finalizando trabalhandores(workers) - 3.6")
		}(w, PortaChans[w])
	}
	fmt.Println("Iniciando leitura de arquivo")
	/*leitura de arquivo*/
	file, err := os.Open(nomeArquivo)
	if err != nil {

	}
	defer file.Close()
	fmt.Println("Pulando linha")
	scanner := bufio.NewScanner(file)
	linhaIdx := 0
	scanner.Scan() //pula o cabecalho (primeira linha)
	fmt.Println("Lendo dados/distribuindo dados")
	for scanner.Scan() {
		ch := PortaChans[linhaIdx%numWorkers]
		ch <- scanner.Text()
		linhaIdx++
	}
	fmt.Println("fechando channels")
	//fechar o contrato dos trabalhadores(demissao)
	for _, ch := range PortaChans {
		close(ch)
	}
	fmt.Println("aguardando")
	espera.Wait() //aguarda todos os trabalhadores terminarem o trabalho
	fmt.Println("termino de processamento")
	return nil
}

func main() {
	ProcessarArqComWorkers("aula4.txt", 4)
}
