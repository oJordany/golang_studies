package main

import (
	"fmt"
	"time"
)

func main() {
	var nome string
	fmt.Println("aaaaaaa")
	go func() {
		nome = "sms 1"
		time.Sleep(1 * time.Second)
		fmt.Println(nome)
	}()
	go func() {
		nome = "email 1"
		time.Sleep(3 * time.Second)
		fmt.Println(nome)
	}()
	go func() {
		nome = "push notification 1"
		time.Sleep(2 * time.Second)
		fmt.Println(nome)
	}()

	select {}
	// fmt.Println("Hello, World! 4")
}
