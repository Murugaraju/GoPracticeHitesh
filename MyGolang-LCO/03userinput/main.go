package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to userinput :"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating of IBM cloud: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')

	fmt.Println("printing the user input: ", input)

	fmt.Printf("Priting the type of variable: %T \n", input)

}
