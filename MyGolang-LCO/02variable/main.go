package main

import "fmt"

const LoginToken = "asdfasdfasd"

func main() {

	var username string = "Murugaraju"

	var isLoggedIn bool = true

	fmt.Println(username)
	fmt.Printf("Variable is of type is: %T \n", username)
	fmt.Printf("isLogged in type is: %T \n", isLoggedIn)

	var smallValue int = 256
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type is: %T \n", smallValue)

	var smallFloat float64 = 255.4554533434242
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type is: %T \n", smallFloat)

	// default values and aliases

	var defaultForInt int
	fmt.Println(defaultForInt)
	fmt.Printf("Variable is of type is: %T \n", defaultForInt)

	var defaultForString string
	fmt.Println(defaultForString)
	fmt.Printf("Variable is of type is: %T \n", defaultForString)

	// implicit type
	var website = "http://Murugaraju.com"
	fmt.Println(website)
	// website = 3

	// no var style

	nuberOfUser := 30000
	fmt.Println(nuberOfUser)

	fmt.Println(LoginToken)

}
