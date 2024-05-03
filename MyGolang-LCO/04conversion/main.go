package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	// add 1 + to user value
	numRating, err := strconv.ParseFloat(strings.TrimSpace(userInput), 64)

	if err != nil {
		fmt.Println("priting the error ", err)
	}

	fmt.Println("Thanks for rating ", numRating+1)

}
