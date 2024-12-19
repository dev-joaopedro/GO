package main

import "fmt"

func main(){
	// Declarar um map e iniciar ela

	var capitais map[string]string = make(map[string]string)

	capitais["França"] = "Paris"
	capitais["Itália"] = "Roma"
	capitais["Brasil"] = "Brasília"
	capitais["Japão"] = "Tóquio"

	fmt.Println(capitais["Brasil"])
	fmt.Println("Capital do Brasil:", capitais["Brasil"])

	capitais["Japão"] = "Esquina"

	fmt.Println(capitais)
}