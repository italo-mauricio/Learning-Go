// aula 2, operador curto de demonstração

package main


import (

	"fmt"

)

func main(){

	x := 10.0
	y := "olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T", y, y)

	x, z := 20,  30

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T", z, z)

}