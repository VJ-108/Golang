package main

import "fmt"

func main() {
	fmt.Println("Methods")
	//No inheritance in golang; No super or parent

	hello := User{"Hello", "Hello@gmail.com", true, 16}
	hello.GetStatus()
	hello.NewMail() //New mail will not be reflected outside NewMail function
}
func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewMail() {
	u.Email = "Hii@gmail.com"
	fmt.Println("Email: ", u.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
