package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Exercício 1: Hello World
	helloWorld()

	// Exercício 2: Soma de dois números
	somaDoisNumeros()

	// Exercício 3: Ímpar ou Par
	imparOuPar()

	// Exercício 4: Tabuada
	tabuada()

	// Exercício 5: Inverter String
	inverterString()

}

// Exercício 1: Hello World
func helloWorld() {
	fmt.Println("Exercício 1: Hello, World!")
}

// Exercício 2: Soma de dois números
func somaDoisNumeros() {
	fmt.Println("Exercício 2: Soma de dois números")
	var num1, num2 int
	fmt.Print("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Printf("A soma de %d e %d é %d\n", num1, num2, num1+num2)
}

// Exercício 3: Ímpar ou Par
func imparOuPar() {
	fmt.Println("Exercício 3: Verificar se um número é ímpar ou par")
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("O número é par.")
	} else {
		fmt.Println("O número é ímpar.")
	}
}

// Exercício 4: Tabuada
func tabuada() {
	fmt.Println("Exercício 4: Tabuada")
	var num int
	fmt.Print("Digite um número para a tabuada: ")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}

// Exercício 5: Inverter String
func inverterString() {
	fmt.Println("Exercício 5: Inverter uma string")
	var texto string
	fmt.Print("Digite uma string: ")
	reader := bufio.NewReader(os.Stdin)
	texto, _ = reader.ReadString('\n')
	texto = strings.TrimSpace(texto)
	
	runes := []rune(texto)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println("String invertida:", string(runes))
}
