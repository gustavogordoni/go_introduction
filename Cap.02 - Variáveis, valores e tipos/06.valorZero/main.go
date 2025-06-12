// VALOR ZERO

/*
	Zeros:

	ints: 0
	floats: 0.0
	booleans: false
	strings: ""
	pointers, functions, interfacesm slices, channels, maps: nil
*/

package main

import "fmt"

var a int
var b float64
var c string

func main() {
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T", c, c)
}
