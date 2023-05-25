package main

import "fmt"

//Creating a struct

type User struct {
	username,email,address string
}



func main() {
	// initializing a struct
	user1 := User{username: "csaglili",email: "cecilselorm34@gmail.com",address: "Akosombo"}

	user1.changeUserName("spanky")


	fmt.Println(user1.username)

}

//methods for structs

// method to change username

func (u *User) changeUserName(newName string){
	u.username = newName
}