package main

import (
	"fmt"
)

func novoUsuario(username string) (string, bool){
	
	if username != ""{
		return username, true
	}

	return "", false
}



func main() {
	
	usuario1 := "Joao Pedro"

	resultado, status := novoUsuario(usuario1)
	fmt.Printf("Usuario atual: %s Status atual: %t \n", resultado, status)

	resultado2, status2 := novoUsuario("")
	fmt.Printf("Usuario atual: %s Status atual: %t \n", resultado2, status2)

	func(){
		fmt.Println("FUNÇÃO ANÔNIMA CHAMADA")
	}()
}