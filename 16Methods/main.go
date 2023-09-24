package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	riyaz := User{"Riyaz", "linux.riyaz@gmail.com", true, 32}
	fmt.Println(riyaz)
	fmt.Printf("Riyaz details are : %+v\n", riyaz)
	fmt.Printf("Name is %v and email is %v.\n", riyaz.Name, riyaz.Email)
	riyaz.GetStatus()
	riyaz.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is use active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "riyazuddin@live.com"
	fmt.Println("Email of this user has changed, New Email is", u.Email)

}
