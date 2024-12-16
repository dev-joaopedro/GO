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

	//nil = nulo, se tem algo dentro de erro ele entra no if pq é diferente de nulo 0
	if err != nil {
		fmt.Println("ERRO AO CONVERTER")
	}else{
		fmt.Println(valorInteiro)
	}

	//String para float
	var pi string = "3.14159"

	piConvertido, error2 := strconv.ParseFloat(pi, 64)

	if error2 != nil {
		fmt.Println("ERRO AO CONVERTER PARA FLOAT")
	}else{
		fmt.Println(piConvertido)
	}
}