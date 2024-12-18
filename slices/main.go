package main

import "fmt"

//Slices são mais flexiveis e poderoso do que arrays. Um slice não tem tamanho fixo, permitindo que você adicione ou remova elementos dinamicamente.

func main() {

	var lista []int //Declarando slices de inteiro

	fmt.Println("Slice vazio", lista)

	lista = append(lista, 10, 50, 30, 20)

	fmt.Println("Slice com valores:", lista)

	myArray := [7]int{1,2,3,4,5,6,7}

	fmt.Println("Meu array:", myArray)

	//Criar um slice a partir de um array existente.

	mySlice := myArray[1:5] // Criar um slice incluindo elementos do incide 1 a 4

	fmt.Println("Slice a partir do Array:", mySlice)
}
