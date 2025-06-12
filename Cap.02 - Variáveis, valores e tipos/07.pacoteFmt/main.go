// O PACOTE FMT

// Format printing
/*
	#1 - Print -> standard out
	Print()
	Println()
	Printf() - (%v, %T)

	#2 - Sprint -> string, pode ser usado como variável
	Sprint()
	Sprintln()
	Sprintf()

	#3 - Fprint -> fle, writer interface, e.g. arquivo ou resposta de servidor
	Fprint()
	Fprintln()
	Fprintf()
*/

package main

import "fmt"

func main() {

	// interpreted string literals
	x := "Oi bom dia \ncomo vai?\t espero \"que\" tudo bem."
	fmt.Println("interpreted: ", x, "\n")

	// raw string literals
	y := `"Oi bom dia \ncomo vai?\t espero \"que\" tudo bem."`
	fmt.Println("raw: ", y, "\n")

	// fmt.Println() 	Adiciona uma linha ao final

	a := "Oi"
	b := "Boa tarde"
	c := fmt.Sprint(a, b) // Valor vai ser salvo em uma string

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)

}
