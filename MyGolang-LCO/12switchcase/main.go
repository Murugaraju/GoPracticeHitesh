package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch case progaramming!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("diceNumber", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value 1 open")
		fallthrough
	case 2:
		fmt.Println("Falthrou 2")
	default:
		fmt.Println("What was that")
	}

}
