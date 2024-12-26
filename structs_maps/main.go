package main

import "fmt"

type Usuario struct{
	Nome string
	Idade int
	status bool
	permissao string
}

func main(){

	usuarios := map[int]Usuario{
		1:{Nome: "Johny", Idade: 24, status: true, permissao: "admin"},
		2:{Nome: "Kamille", Idade: 21, status: true, permissao: "rainha"},
	}

	usuarios[3] = Usuario{Nome: "Jose", Idade: 60, status: true, permissao: "admin"}

	usuarios[len(usuarios)+1] = Usuario{Nome: "Henrique", Idade: 30, status: true, permissao: "admin"}

	i := 1
	for i <= len(usuarios){
		fmt.Printf("Nome: %s , Idade %d , Permissão atual: %s \n", usuarios[i].Nome, usuarios[i].Idade, usuarios[i].permissao)
		i++
	}

	fmt.Println("============")

	usuarios[len(usuarios)+1] = Usuario{Nome: "Pedro", Idade: 20, status: true, permissao: "admin"}

	usuarios[1] = Usuario{Nome: "Pedro Programador", Idade: 25, status: false, permissao: "admin"}

	fmt.Println(usuarios)
}