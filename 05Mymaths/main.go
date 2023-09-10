package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to random number generation.")
	//var num1 int = 25
	//var num2 float64 = 4.5
	//fmt.Println("sum of it ", num1+int(num2))

	//fmt.Println(rand.Intn(100))
	myRandom, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandom)
}
