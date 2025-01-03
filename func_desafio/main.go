package main

import "fmt"

type Aluno struct{
	Nome string
	Nota1 float64
	Nota2 float64
}

func calculaMedia(aluno Aluno) (string, float64){
	media := (aluno.Nota1 + aluno.Nota2) / 2

	return aluno.Nome, media
}

func verificaAprovacao(media float64) string{
	if(media >= 7){
		return "Aprovado!"
	}

	return "Você ficou de recuperação!"
}

func main(){

	aluno1 := Aluno{Nome: "Joao", Nota1: 6.0, Nota2: 5.0}

	fmt.Printf("%s, Nota1: %.2f, Nota2: %.2f \n", aluno1.Nome, aluno1.Nota1,aluno1.Nota2)

	nome, media := calculaMedia(aluno1)

	fmt.Printf("Media do %s: %.2f \n", nome, media)

	statusAluno1 := verificaAprovacao(media)
	fmt.Println(statusAluno1)

	fmt.Println("==============================================================")

	aluno2 := Aluno{Nome: "Kamille", Nota1: 9.0, Nota2: 8.0}
	fmt.Printf("%s, Nota1: %.2f, Nota2: %.2f \n", aluno2.Nome, aluno2.Nota1,aluno2.Nota2)
	
	nome2, media2 := calculaMedia(aluno2)

	fmt.Printf("Media do %s: %.2f \n", nome2, media2)
	statusAluno2 := verificaAprovacao(media2)
	fmt.Println(statusAluno2)
}