package main

import "fmt"

func main() {

	// Arrays > Um array em Go é uma sequência de elementos de um único tipo, com um tamanho fixo definido na declaração.
	var numeros [4]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 15
	numeros[3] = 05

	fmt.Println(numeros)

	var valores = [3]int{3, 10, 90}
	fmt.Println(valores[2])

	permissoes := [3]string{"usuario", "admin", "editor"}
	fmt.Println(permissoes[1])

}
