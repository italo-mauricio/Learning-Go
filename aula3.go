/*

	Aula 3, a respeito da palavra-chave Var


*/

package main


import (

	"fmt"

)
var y = 10   // o var funciona fora do codeblock

func main(){
	
	z := 20

	random (z)

}

func random (x int){
	fmt.Println(y)
	fmt.Println(x)



}