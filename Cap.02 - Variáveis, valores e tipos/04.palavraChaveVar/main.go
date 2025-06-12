// A PALAVRA-CHAVE VAR

// Lição principal:

package main

import "fmt"

var z = 60

func main() {

	y := 10
	funcaoAleatoria(y)
}

func funcaoAleatoria(x int) {
	// Para esse escopo, y não existe
	// fmt.Println(y) - undefined: y

	// Exibindo a var passada como parâmetro na func
	fmt.Println("x: ", x)

	// Para esse escopo, z existe, pois está como global
	fmt.Println("z: ", z)
}
