package main

import "fmt"

type Pessoa struct {
	nome string
	idade int
	nacionalidade string
}

func main(){
	
	pessoa1 := Pessoa{nome: "Matheus Fraga", idade: 18, nacionalidade: "Brasileiro"}
	pessoa2 := Pessoa{nome: "João", idade: 30, nacionalidade: "Brasileiro"}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	pessoas := map[int]Pessoa{
		1:{nome: "Johny", idade: 24, nacionalidade: "Americano"},
		2:{nome: "Ka", idade: 21, nacionalidade: "Brasileiro"},
	}

	fmt.Println(pessoas)

}