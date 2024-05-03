package main

import "fmt"

func main() {
	fmt.Println("Welcome to IfElse in Go lang")

	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "You are regular user"
	} else if loginCount > 20 {
		result = "You are awesome user"
	} else {
		result = "You are Awesome2 user"
	}
	fmt.Println("result", result)

	if num:= 3; num<10 {
		fmt.Println("Num is less than 3")
	} else{
		fmt.Println("Num is more than 3")
	}

}
