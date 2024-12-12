package main

import "fmt"

func main (){
	var nome string = "Lucas Silva"
	var idade int = 28
	var texto string			

	fmt.Println("Variaveis")
	fmt.Println(nome)
	fmt.Println(idade)

	texto = "Meu primeiro projeto em GO"
	idade = 29
	nome = "Sujeiro programador"

	fmt.Println("Agora a idade é ", idade)
	fmt.Println(texto)
	fmt.Println(nome)

	//Variavel do tipo CONST

	const pi = 3.14159

	fmt.Println(pi)
}