package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func ProcessarArquivoComWorkers(nomeArquivo string, numWorkers int) error {
	// Contratação de trabalhadores

	PortaChans := make([]chan string, numWorkers)

	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ {
		PortaChans[w] = make(chan string, 100)
		wg.Add(1)
		go func(workerId int, ch <-chan string) {
			defer wg.Done()
			outputfile, err := os.Create(fmt.Sprintf("work%d.txt", workerId))
			if err != nil {
				fmt.Printf("erro ao criar arquivo do worker %d: %v\n", workerId+1, err)
				return
			}
			defer outputfile.Close()
			writer := bufio.NewWriter(outputfile)
			for linha := range ch {
				writer.WriteString(linha + "\n")
			}
			writer.Flush()
		}(w, PortaChans[w])
	}

	// Leitura do arquivo
	file, err := os.Open(nomeArquivo)

	if err != nil {
		fmt.Println("erro ao abrir o arquivo %s: %v", nomeArquivo, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linhaIdx := 0
	scanner.Scan() // pula o cabeçalho
	/*
		resto de 0/4 = 0
		resto de 1/4 = 1
		resto de 2/4 = 2
		resto de 3/4 = 3
		resto de 4/4 = 0
		resto de 5/4 = 1
		resto de 6/4 = 2
		....
	*/
	for scanner.Scan() {
		ch := PortaChans[linhaIdx%numWorkers]
		ch <- scanner.Text()
		linhaIdx++
	}

	// fechar o contrato dos trabalhadores(demissão)
	for _, ch := range PortaChans {
		close(ch)
	}
	wg.Wait()
	return nil
}

func main() {
	ProcessarArquivoComWorkers("aula4.txt", 4)
}
