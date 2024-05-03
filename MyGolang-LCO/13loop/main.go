package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Loop")

	days := []string{"Monday", "TuesDay", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Print(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 5 {
			break
		}

		fmt.Println("value is ->  ", rogueValue)
		rogueValue++
	}

	rogueValue = 1
	for rogueValue < 10 {
		rogueValue++

		if rogueValue == 6 {
			goto lco
		}

		if rogueValue%2 == 0 {
			goto helo
			continue
		}

		fmt.Println("value is ->  ", rogueValue)
	}

lco:
	fmt.Println("This is goto statement ")

helo:
	fmt.Println("asdfasdf")

}
