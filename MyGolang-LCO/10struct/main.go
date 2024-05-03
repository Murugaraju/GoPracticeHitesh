package main

import "fmt"

func main() {
	fmt.Println("This is for struct,")
	// Struct in go lang is alternate for preparing the Class similar to other languages
	// no inheritance or parent or super class

	var raju User = User{
		"Murugaraju",
		"murugaraju.perumalla@ibm.com",
		true,
		30,
	}
	fmt.Println("Murugaraju user printing: ", raju)
	fmt.Printf("Murugaraju user printing: %T \n", raju)

	fmt.Printf("Full detail of Murugaraju : %+v", raju)

}

type User struct {
	Name   string
	Emain  string
	Status bool
	Age    int
}
