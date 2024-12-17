package main

import "fmt"

func main(){
	nome1 := "matheus"
	nome2 := "lucas"

	numero1 := 5
	numero2 := 7

	fmt.Println(nome1 == nome2)
	fmt.Println(numero1 >= numero2)
	fmt.Println(10 > 20)
	fmt.Println(10 < 20)
	fmt.Println(10 <= 20)

	fmt.Println("Matheus é maior que Lucas?", len(nome1) >= len(nome2))
}