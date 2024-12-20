package main

import "fmt"

func boasvindas(nome string){
	fmt.Println("Bem vindo ao meu programa", nome)
}

func soma (numero1 int, numero2 int)int{
	resultado := numero1 + numero2

	return resultado
}

func main(){
	boasvindas("João")

	soma1 := soma (30,23)
	soma2 := soma (10, 20)

	fmt.Println("Resultado da soma1:", soma1)
	fmt.Println("Resultado da soma2:", soma2)
}