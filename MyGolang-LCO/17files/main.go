package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files, and working on file related operation")
	content := "This is some sample context for working on file related - LearnByDoing.com"

	fmt.Println("content", content)

	// file related os module/package

	file, err := os.Create("./mylocalfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	// fmt.Printf("printing the file value and type of the file ob, object type: %T value: %v", file, file)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("length is: ", length)

	defer file.Close()
	readFile("./mylocalfile.txt")

}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Text data inisedis \n", data)
	fmt.Println("Text data inisedis \n", string(data))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
