package main


import (

	"fmt"
)

// Exercicios em Go

// Utilizando o operador curto de declaração, atribua estes valores as variáveis com os indentificadores "x", "y" e "z"

/*
	1 - 42
	2 - "James Bond"
	3 - true

*/


func main (){

	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("Valor da variável x é %v\nE o seu tipo é %T\n", x, x)
	fmt.Printf("Valor da variável y é %v\nEo seu tipo é %T\n", y, y)
	fmt.Printf("Valor da variável z é %v\nEo seu tipo é %T\n", z, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)


}