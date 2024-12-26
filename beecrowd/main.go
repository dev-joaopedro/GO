package main

import "fmt"

func main() {
    var A, B int

    // Lê os dois valores inteiros da entrada
    fmt.Scan(&A, &B)

    // Calcula a soma de A e B
    X := A + B

    // Imprime o resultado no formato exigido
    fmt.Printf("X = %d\n", X)
}