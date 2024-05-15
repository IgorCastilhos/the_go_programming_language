package main

import (
	"fmt"
	"os"
)

func main() {
	// Exercício 1.2 Modifique o programa echo para exibir o índice e o valor de cada um de seus argumentos, um por linha.
	fmt.Println("Nome do comando: ", os.Args[0])
	for i, arg := range os.Args[1:] {
		s := fmt.Sprintf("Índice: %d, Valor do argumento: %s", i, arg)
		fmt.Println(s)
	}
}
