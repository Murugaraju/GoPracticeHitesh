package main

import "fmt"

func main() {
	// fmt.Println("Welcome to git related Defer")ß
	/*
		Official definition: A `defer` statement involves a function whose execution is deferred to the moment the surrounding function returns,
		either because the suttounding function executed a return statement, reached the end of the its function.


		Analogy or normal way of understanding `defer` working in go.
		When ever a function executes it executes line by line

		Imagine all the defer functions appended for execution and at the end of function
		the execution of those function start from LIFO

	*/

	defer fmt.Println("Hello world ---> 1")
	fmt.Println("Hello")
	myDefer()

	defer fmt.Println("One more hello world ---->2 ")

	// myDefer()

}

func myDefer() {
	for i := range 10 {
		defer fmt.Println("i", i)
	}
}

//
