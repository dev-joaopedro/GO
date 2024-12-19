package main

import "fmt"

type Pessoa struct {
	Nome string
	Idade int
	Nacionalidade string
}

func main(){
	
	pessoa1 := Pessoa{Nome: "Matheus Fraga", Idade: 18, Nacionalidade: "Brasileiro"}
	pessoa2 := Pessoa{Nome: "João", Idade: 30, Nacionalidade: "Brasileiro"}

	fmt.Println(pessoa1.Idade)
	fmt.Println(pessoa2.Nome)

	pessoas := map[int]Pessoa{
		1:{Nome: "Johny", Idade: 24, Nacionalidade: "Americano"},
		2:{Nome: "Ka", Idade: 21, Nacionalidade: "Brasileiro"},
	}

	fmt.Println(pessoas)
	
}