/*

	Aula 3, a respeito da palavra-chave Var


*/

package main


import (

	"fmt"

)
var y = 10

func main(){

	y := 10
	random (y)

}

func random (x int){
	fmt.Println(y)
	fmt.Println(x)



}