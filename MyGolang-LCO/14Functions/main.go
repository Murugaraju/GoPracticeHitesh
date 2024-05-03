package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	greeter()

	result, someString := add(3, 4)
	fmt.Println("Resutl is: ", result, someString)

	// func greeterTwo()  {
	// 	fmt.Println("Another function")
	// }

	// greeterTwo()

	proresult := proAdder(3, 4, 5, 343, 3342, 22)

	fmt.Println("proAdder result -> ", proresult)

}

func greeter() {
	fmt.Println("Namaste from me Raju")
}

// function signature are imp
func add(a int, b int) (int, string) {
	return a + b, "hellow"
}

// ... are variadic function
func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
