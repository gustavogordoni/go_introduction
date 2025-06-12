// CRIANDO O SEU PRÓPRIO TIPO

/*
	Revisando: tipos são fixos.
		Uma vez declarada uma variável como de um certo tipo, isso é imutável.
*/

package main

import "fmt"

type hotdog int

var a hotdog = 23

func main() {
	b := 10
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)

	// b = a  cannot use a (variable of type hotdog) as type int in assignment

}
