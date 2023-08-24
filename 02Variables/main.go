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

	var smallFloat float64 = 255.455245584584564444565
	fmt.Println(smallFloat)
	fmt.Printf("Type of Variable is : %T \n\n", smallFloat)

	// implicit type

	var website = "www.google.com"
	fmt.Println(website)

	// no var style

	var numberOfUser := 35220.0
	fmt.Println(numberOfUser)

	var LoginToken = "Hello World"
	fmt.Println(LoginToken)
	fmt.Printf("Varibale is of type : %T \n", LoginToken)

}
