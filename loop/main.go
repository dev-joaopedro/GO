package main

import "fmt"

func main() {
	
	// FOR
	total := 10

	for i := 0; i < total; i++ {
		fmt.Println(i)
	}

	// WHILE
	numero := 0

	for numero < 5 {
		fmt.Println("PASSANDO NO FOR:", numero)
		numero++
	}

	for {
		fmt.Println("LOOP INFINITO")
		break
	}
	
	for i := 0; i < 100; i++ {
		if i == 50 {
			fmt.Println("Passou dos 50")
			continue
		}

		if i == 60 {
			fmt.Println(i)
			break
		}
	}
}