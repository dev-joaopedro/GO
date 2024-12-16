package main

import "fmt"

func calculadora() {
    fmt.Println("DESAFIO 1")

    nascimento := 2000
    dataatual := 2024

    idade := dataatual - nascimento

    fmt.Println("Idade atual é:", idade)

    fmt.Println("DESAFIO 2")

    hora := 4
    var hrminutos int = hora * 60

    fmt.Println(hrminutos)
}
