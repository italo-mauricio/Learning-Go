// aula 2, operador curto de demonstração

package main


import (

	"fmt"

)

var y = "ola bom dia"

func main(){

	x := 10.0
	y := "olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T", y, y)

	x, z := 20,  30

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
	
	k := 10 < 10

	fmt.Printf("k: %v, %T\n", k, k)

}



// marmota só funciona dentro de um codeblock
/*
	nomes reservados não podem ser usados como variavel
	exemplos (package, fun, var, break, case, chan, const, defaut, map, interface
	continue, defer, else, falithrough, for, go, goto, if, import, range, return
	select, struct, switch, type.)

*/

