package main

import "fmt"

func main(){
	valor1 := 10
	valor2 := 7

	// Adição ( + )
	resultado := valor1 + valor2
	fmt.Println("Valor da soma:", resultado)

	// Subtração ( - )
	subtracao := valor1 - valor2
	fmt.Println("Valor da subtração:", subtracao)

	// Multiplicação ( * )
	resultadoMultiplicacao := valor1 + valor2
	fmt.Println("Valor da Multiplicação:", resultadoMultiplicacao)

	// Divisão ( / )
	nota1 := 10.00
	nota2 := 2.00

	resultadoNota := nota1 / nota2
	
	fmt.Println("Valor da divisão:", resultadoNota)

	x := 10
	y := 3

	resto := x % y

	fmt.Println(resto)

}