package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")
	name, _ := reader.ReadString('\n')

	fmt.Println("Enter your age:")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr) // Remove the newline character
	age, _ := strconv.Atoi(ageStr)     // Convert string to int

	fmt.Printf("Your name is: %s", name)
	fmt.Printf("Your age is: %d\n", age)

	// Add a pause at the end of the program
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
