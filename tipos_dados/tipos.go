package main

import "fmt"

// String
// bool > True False
// int | int8 | int16 | int32 | int64
// float32 float64
// Bytes, Runes


func main(){
	var nome string = "johny"
	var isAdmin bool = true
	idade := 28
	
	var valor float64 = 900.90
	valor2 := 350.20

	fmt.Println(valor)
	fmt.Printf("Valor em float é %.2f" ,valor2)
	
	var numero int8 = 127

	fmt.Println(nome)
	fmt.Println(isAdmin)
	fmt.Println(idade)
	fmt.Println(numero)

	var teste byte = '?'
	fmt.Println(teste)
}