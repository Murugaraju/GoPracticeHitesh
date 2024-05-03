package main

import "fmt"

func main() {
	fmt.Println("This is for struct related method")
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

	fmt.Printf("Full detail of Murugaraju : %+v \n", raju)

	// Trying to execute the method

	raju.GetStatus()

	new_mail := raju.NewMail()

	fmt.Println("The new mail is: and the old email is: ", new_mail, raju.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() { // remember the u User is now the exact copy of the value being happening, later mode will see usage of pointer to change the original instance obj
	fmt.Println("Status of the user", u.Status)
}

func (u User) NewMail() string {
	u.Email = "test@go.dev"
	return u.Email

}
