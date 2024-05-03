package main

import "fmt"

func main() {
	fmt.Println("Hellow array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Printing the fruitList value: ", fruitList)
	fmt.Println("Printing the fruitList value: ", len(fruitList))

	var vegList = [3]string{"potato", "tomato", "newveg"}

	fmt.Println("Printing the vegList value: ", vegList)
	fmt.Println("Printing the vegList value: ", len(vegList))

}
