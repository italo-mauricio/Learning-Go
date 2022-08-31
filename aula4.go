package main


import (

    "fmt"

)

            /*

                trabalhando com tipos

            */

var x int = 10

func main(){

    x = 20
    fmt.Println(x)

}

// quando fazemos a declaração da variável sem atribuir valor, essa atribuição só pode ser feita dentro de um codeblock
// tipos básicos: int, string, bool
// tiips de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário, exemplo (slice, array, struct, map)