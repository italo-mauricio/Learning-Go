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



Pacote fmt:
	Setup: strings, ints, bools
	Striings: interpreted string literal vs. raw string literals. 
		Rune Literals. 
		Em ciência da computação, um literal é uma notação para representar um valor fixo no código fonte. 
		Format printing: documentação.
			Grupo #1: Print -> standard out
			func Print(a ...interface{}) (n int, err error)
			func Println(a ...interface{}) (n int, err error)
			func Printf(format string, a ...interface{}) (n int, err error)
			Format verbs. (%v %T)
			Grupo #2: Print -> string, pode ser usado como variável
			func Sprint(a ...interface{}) string
			func Sprintf(format string, a ...interface{}) string
			func Sprintln(a ...interface{}) string
			Grupo #3: Print -> file, writer interface, e.g. arquivo ou resposta de servidor
			func Fprint(w io.Writer, a ...interface{}) (n int, err error)
			func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
			func Fprintln(w io.Writer, a ...interface{}) (n int, err error)


			by: https://github.com/vkorbes/aprendago/blob/master/OUTLINE.md
*/


func main(){

	x = 10
	fmt.Printf("%v, %T", x, x)
	x = 20
	fmt.Printf("%v, %T", x, x)




}