package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to our application"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our pizza : ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating : ", input)

	fmt.Printf("Type of string is %T", input)

}