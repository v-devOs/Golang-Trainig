package main

import (
	"fmt"

	"rsc.io/quote"
)

// go mod init 'nombre'
// go get nombre 'paquete de terceros'

/*Declaración de constantes a nivel de paquete*/
const Pi = 3.14

const (
	x = 100
	y = 0b1010 // binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

// iota inicia en 0 y hereda el siguiente valor sucesivamente
const (
	Domingo = iota + 1
	Lunes
	Martes
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Go())

	// fullName := "Uriel Galindo \t(alias  \"vdeVos\")\n"

	// Podemos obtener el valor asscii de algun caracter
	// var a byte = 'a'

	// Representación de caracteres unicode

	var emoki rune = '❤'

	// Declaración de variables

	/*Forma uno de declarar variables*/
	// var firstName, lastName string
	// var age int

	/*Forma dos de declarar variables*/
	// Si inciamos una variable en su declaración no es necesario especificar el tipo de dato
	// var firstName, lastName, age = "Uriel", "Galindo", 20

	//Declaración e inicialisación en linea solo se pueden declarar dentro de funciones
	// firstName, lastName, age := "Uriel", "Galindo", 20

	// fmt.Println(firstName, lastName, age, Pi, x, y, z, w, Martes)

	fmt.Println(emoki)
}
