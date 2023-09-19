package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	langs := make(map[string]string)

	langs["JS"] = "JavaScript"
	langs["RB"] = "Ruby"
	langs["PY"] = "Python"

	fmt.Println("List of all languages: ", langs)

	fmt.Println("JS shorts for : ", langs["JS"])

	// Delete

	delete(langs, "RB")
	fmt.Println("List of all Languages: ", langs)

	// Loops
	for key, value := range langs {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}