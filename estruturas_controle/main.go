package main

import "fmt"

	//if else | else if
	//switch

func main() {
	nota1 := 7.0
	nota2 := 7.0

	media := (nota1 + nota2) / 2

	if media >= 7 {
	fmt.Println("Aprovado")	
	} else{
		fmt.Println("Reprovado")
	}

	//**********************************

	dia := 6

	switch dia{
		case 1:
			fmt.Println("Segunda")
		case 2:
			fmt.Println("Terça")
		case 3:
			fmt.Println("Quarta")
		case 4:
			fmt.Println("Quinta")
		case 5:
			fmt.Println("Sexta")
		default:
			fmt.Println("Fim de semana")
	}

	redeSocial := "blog"

	switch redeSocial {
	case "Instagram":
		fmt.Println("@Joao")
	case "youtube":
		fmt.Println("@joao")
	case "blog":
		fmt.Println("https://meusite.com.br")
	default:
		fmt.Println("Você errou o acesso!")
	}
}