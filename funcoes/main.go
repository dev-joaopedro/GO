package main

import (
	"errors"
	"fmt"
)

func boasvindas(nome string){
	fmt.Println("Bem vindo ao meu programa", nome)
}

func soma (numero1 int, numero2 int)int{
	resultado := numero1 + numero2

	return resultado
}

type Usuario struct{
	Nome string
	Senha string
}

func Autentica(user Usuario, senha string) (string, error) {

	if user.Senha != senha {
		return "", errors.New("Senha Inválida!")
	}

	return user.Nome, nil
}

func main(){
	boasvindas("João")

	soma1 := soma (30,23)
	soma2 := soma (10, 20)

	fmt.Println("Resultado da soma1:", soma1)
	fmt.Println("Resultado da soma2:", soma2)

	usuario1 := Usuario{"João", "1234"}
	usuario2 := Usuario{"Kami", "123456"}

	nome, err := Autentica(usuario1, "1234")

	if err != nil {
		fmt.Println("Erro na autenticação", err)
	}else {
		fmt.Println("Autenticação realizada com sucesso", nome)
	}

	nome, err = Autentica(usuario2, "12353")

	if err != nil {
		fmt.Println("Erro na autenticação", err)
	}else {
		fmt.Println("Autenticação realizada com sucesso", nome)
	}
}