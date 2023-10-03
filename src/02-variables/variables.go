package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	// Declaración de variables

	/*Forma uno de declarar variables*/
	// var firstName, lastName string
	// var age int

	/*Forma dos de declarar variables*/
	// Si inciamos una variable en su declaración no es necesario especificar el tipo de dato
	// var firstName, lastName, age = "Uriel", "Galindo", 20

	//Declaración e inicialisación en linea solo se pueden declarar dentro de funciones
	firstName, lastName, age := "Uriel", "Galindo", 20

	fmt.Println(firstName, lastName, age)
}
