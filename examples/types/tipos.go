package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Type: %T - Value: %v\n", true, true)
	fmt.Printf("Type: %T - Value: %v\n", "steph", "steph")
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	fmt.Printf("Type: %T - Value: %v\n", 1.5, 1.5)
}

// Tipos:
// bool (true or false)
// string (sequence of bytes)
// int
// float (float64/float32) - decimal
