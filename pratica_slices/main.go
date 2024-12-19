package main

import "fmt"

func main(){
	var tarefas []string

	// Adicionar itens no slice

	tarefas = append(tarefas, "Estudar GO")
	tarefas = append(tarefas, "Seguir o canal do Johny")
	tarefas = append(tarefas, "Ler um livro")

	fmt.Println("Tarefas:", tarefas)
	fmt.Println("Tamanho do slice atual:", len(tarefas))

	// Slicing / Cortar

	//tarefas = tarefas[1:]

	//fmt.Println("Removendo a 1 tarefa:", tarefas)

	//tarefas = tarefas[:len(tarefas) - 1]
	//fmt.Println("Removendo ultimo item:", tarefas)

	
	//tarefas = append(tarefas[:1], tarefas[2:]...)
	//fmt.Println("Remover o segundo item:", tarefas)

}