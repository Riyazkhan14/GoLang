package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Kiwi"

	fmt.Println("this is the fruit basket", fruitList)
	fmt.Println("Type of this is the fruit basket", len(fruitList))

	var vegList = [3]string{"Tomato", "Beans", "Pulses"}
	fmt.Println("The veggies list is :", vegList)
	fmt.Println("Type of this veggies list is :", len(vegList))

}
