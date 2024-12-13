package main

import (
	"fmt"
	"strconv"
)

func main(){
	// Conversão de Inteiro para float

	var valor int = 42
	var valorConvertido float64 = float64(valor)

	fmt.Println("Inteiro:", valor)
	fmt.Println("Valor em Float:", valorConvertido)

	// Converter int para string
	var valorString string = strconv.Itoa(valor)

	fmt.Println(valorString)

	// Converter de string para int
	var valor2 string = "230"
	
	valorInteiro,err := strconv.Atoi(valor2)

	if err != nil {
		fmt.Println("ERRO AO CONVERTER")
	}else{
		fmt.Println(valorInteiro)
	}

	fmt.Println(valorInteiro)
}