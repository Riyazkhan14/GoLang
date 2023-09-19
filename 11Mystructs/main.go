package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	riyaz := User{"Riyaz", "linux.riyaz@gmail.com", true, 32}
	fmt.Println(riyaz)
	fmt.Printf("Riyaz details are : %+v\n", riyaz)
	fmt.Printf("Name is %v and email is %v.", riyaz.Name, riyaz.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
