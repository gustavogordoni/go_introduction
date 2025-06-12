// OPERADOR CURTO DE DECLARAÇÃO

// Lição principal:
// := utilizado par criar novas variáveis, dentro de code blocks (escopos)
// = para atribuir valor a variáveis já existentes

package main

import "fmt"

// Escopo global
var w = "Vugnaes Sreo"

func main() {

	// := 	Operador de declaração,
	// Apenas funciona em um escopo (nesse caso main)
	x := 10
	y := "Olá, bom dia!"
	z := 11.1
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	// = 	Operador de atribuição
	x, z = 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n\n", z, z)

	// Escopo global
	fmt.Printf("w: %v, %T\n\n", w, w)

	// Valor por operação
	x = 100 - 70
	fmt.Printf("x: %v, %T\n\n", x, x)

	// Valor por comparação
	// a := 10 == 10
	a := 10 >= 10
	fmt.Printf("a: %v, %T\n", a, a)
}
