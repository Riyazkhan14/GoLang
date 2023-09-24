package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in goLang")

	content := "this needs to go in file - riyaz.com"

	file, err := os.Create("./myfirstfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is : ", length)
	defer file.Close()
	readFile("./myfirstfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	panic(err)
	//}
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
