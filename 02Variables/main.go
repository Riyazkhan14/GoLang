package main

import "fmt"

func main() {
	var username string = "Riyaz"
	fmt.Println("Welcome : ", username)
	fmt.Printf("Type of Varible is : %T \n\n", username)

	var isNum int = 20
	fmt.Println(isNum)
	fmt.Printf("Type of Variable is : %T \n\n", isNum)

	var isitBool bool = true
	fmt.Println(isitBool)
	fmt.Printf("Type of Variable is : %T \n\n", isitBool)

	var isFloat float32 = 2000.222020292
	fmt.Println(isFloat)
	fmt.Printf("Type of Variable is : %T \n\n", isFloat)
}
