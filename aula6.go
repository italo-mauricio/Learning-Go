package main

import (

	"fmt"
)

/*
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

/* func main() {

	x := `"ola bom dia\n como vai?\nespero que tudo bem"`
	fmt.Println(x)



	// com o acento ` a formatação fica literal, não importando de qual jeito eu digite

}

*/

func main(){

	x := "oi"
	y := "bom dia"

	z := fmt.Sprint(x, " ", y)

	fmt.Println(z)



}