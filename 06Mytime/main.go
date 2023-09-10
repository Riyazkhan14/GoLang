package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study")
	curTime := time.Now()
	fmt.Println(curTime.Format("02-01-2006 Monday"))

	createDate := time.Date(1990, time.October, 14, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
