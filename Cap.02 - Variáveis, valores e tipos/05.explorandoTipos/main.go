// EXPLORANDO TIPOS

// Tipagem estática, uma variável não muda seu tipo.

package main

import "fmt"

var x int = 10

func main() {

	// x = 20.5 	cannot use 20.5 (untyped float constant) as int value in assignment (truncated)
	fmt.Printf("x: %v, %T\n", x, x)
}
