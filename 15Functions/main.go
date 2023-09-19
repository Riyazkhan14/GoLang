package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greetings()

	result := adder(3, 5)
	fmt.Println("Result is : ", result)

	proRes := proAdder(2, 5, 6, 7, 8, 8, 0)
	fmt.Println("Pro Result is : ", proRes)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val

	}
	return total
}

func greetings() {
	fmt.Println("Salam from Go Lang")
}
