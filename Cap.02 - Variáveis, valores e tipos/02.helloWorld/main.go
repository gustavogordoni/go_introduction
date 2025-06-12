// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	x := 16
	y := "strings"
	z := true

	numeroBytes, erros := fmt.Println("Hello, World!", "Olá pessoal!", 100)
	fmt.Println(numeroBytes, erros)
	fmt.Println(x, y, z)
}
