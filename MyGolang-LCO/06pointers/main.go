package main

import "fmt"

func main() {
	fmt.Println("welcome to Pointers")

	// var ptr *int

	// fmt.Println("Value of the pointer is: ", ptr)

	myNumber := 23

	var ptr *int = &myNumber
	fmt.Println("Value of the actual pointer is: ", ptr)
	fmt.Println("Value of the actual pointer is: ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("New value of myNumber is: ", myNumber)

}
