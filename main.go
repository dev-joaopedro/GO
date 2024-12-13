package main

import "fmt"

func main() {
	var nome string
	var idade int
	var salario float32

	fmt.Println("Entre com seu nome: ")
	fmt.Scanln(&nome)
	fmt.Println("Entre com sua idade: ")
	fmt.Scanln(&idade)
	fmt.Println("Entre com seu salario: ")
	fmt.Scanln(&salario)

	fmt.Println("Seu nome é", nome)
	fmt.Println("Sua idade é", idade)
	fmt.Println("Seu salario é", salario)
	
}