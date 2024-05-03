package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello Slices")

	var fruitList = []string{"Aplle", "Tomato", "Banana"}

	fmt.Printf("The type of fruitList is: %T \n", fruitList)
	fmt.Println("The type of fruitList is: ", fruitList)

	fruitList = append(fruitList, "Mango")

	fmt.Println("The value in fruitList is: ", fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println("The value in fruitList is: ", fruitList)

	// look hitesh video again for more info

	highScores := make([]int, 4)
	highScores[0] = 232
	highScores[1] = 34
	highScores[2] = 345
	highScores[3] = 345
	// highScores[4] = 232

	highScores = append(highScores, 793, 343)

	fmt.Println("Printing high scores: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove value from slice based on index

	var courses = []string{"reactjs", "golang", "python", "ruby"}

	fmt.Println("courses", courses)

	index := 2

	courses = append(courses[:1], courses[index+1:]...)
	fmt.Println("courses", courses)

}
