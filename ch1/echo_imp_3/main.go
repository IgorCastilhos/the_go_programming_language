/* 3ª Forma, mais eficiente e simples, porém, coloca todas as strings juntas em uma única linha */
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}
