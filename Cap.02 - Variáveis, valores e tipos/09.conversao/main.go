// CONVERSÃO, NÃO COERÇÃO

/*

 */

package main

import "fmt"

type hotdog int

var a hotdog = 23

func main() {
	b := 10
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n\n", b, b)

	// b = a  cannot use a (variable of type hotdog) as type int in assignment

	b = int(a)
	fmt.Print("int(b): conversão de hotdog para int\n")
	fmt.Printf("b: %v, %T\n", b, b)

}
