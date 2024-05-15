package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Exercício 1.1: Modifique o programa echo para exibir também o nome do comando que o chamou.
	fmt.Println("Nome do comando: ", os.Args[0])
	fmt.Println(strings.Join(os.Args[0:], ""))
}
