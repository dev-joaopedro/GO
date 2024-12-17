package main

import "fmt"

func main(){
	// and (&&)
	// or (||)
	// not (!)

	estoque := true
	vendaLiberada := true
	freteGratis := false
	
	nome1 := "matheus"
	nome2 := "lucas"

	fmt.Println("AND:",estoque && vendaLiberada)
	fmt.Println("OR:", estoque || vendaLiberada)
	fmt.Println("NOT:", !freteGratis)

	fmt.Println("Os nomes são diferente!", nome1 != nome2)

}