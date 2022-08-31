package main


import (

	"fmt"

)

// se sua varável for um string, sua valor zer vai ser 1
//   sobre o valor zero

/*
	Valores de zero:
	int = 0
	floats = 0.0
	booleans = false
	strings = ""
	pointers, functions, interfaces, slices, channels, maps: nil

Use := sempre que possível. 
Use var para package-level scope



			by: https://github.com/vkorbes/aprendago/blob/master/OUTLINE.md
*/

var a int
var b float64
var c string
var d bool

func main(){

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)	
	fmt.Printf("%v, %T\n", d, d)
	

}