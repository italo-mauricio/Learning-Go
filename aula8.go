package main


// conversão, não coerção

import (

	"fmt"

)
type hotdog int
var b hotdog = 102

func main() {

	x := 10
	fmt.Printf("%v\n", x)
	
	x = int (b)
	fmt.Printf("%v\n", x)

}


/*
	Conversão, não coerção:
		Conversão de tipos é o que soa.
		Em Go não se diz casting, se diz conversion.
		a = int(b)
		ref/spec#Conversions
		Fim da sessão. Parabéns! Dicas, motivação e exercícios.


		by: https://github.com/vkorbes/aprendago/blob/master/OUTLINE.md

*/