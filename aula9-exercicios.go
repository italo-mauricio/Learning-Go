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

var x int = 42
var y string = "James Bond"
var z bool = true

func main (){

	fmt.Printf("%v, %T", x, x, "%v, %T\n", y, y, "%v, %T\n", z, z)


}