package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world!")
	_, err := os.Open("")
	if err != nil {
		fmt.Println("aconteceu um erro aqui ->", err)
		return
	}
}

func init() {
	fmt.Println("Initiailizing...")
}
