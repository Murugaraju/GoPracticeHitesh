package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	parseTime := time.Now()

	fmt.Println(parseTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.January, 18, 0, 0, 0, 0, time.UTC)

	fmt.Println("createdDate", createdDate.Format("2006-01-02 Monday"))

}
