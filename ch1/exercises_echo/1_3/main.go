package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Exercício 1.3: experimente medir a diferença de tempo de execução entre os programas echo das versões ineficiente e eficiente.
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], ""))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
