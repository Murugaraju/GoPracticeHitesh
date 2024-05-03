package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hey, welcome to Maps")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println("The languages: ", languages)
	fmt.Println("js stands for: ", languages["js"])

	// delete a key value pair

	delete(languages, "js")
	fmt.Println("The languages: ", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For the key:%v -> value:%v \n",key, value )
	}
}
