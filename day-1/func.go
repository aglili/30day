package main

import "fmt"

type User struct{
	userAge int
	UserName,userPassword string
	isAdmin bool
}

func main() {
	person1 := User{
		userAge: 23,userPassword: "abakobi",UserName: "Jake",isAdmin: false,
	}

	
	person1.changeUserName("kwaku")
	person1.changePassword("jhdsjhiseuvnsivd")

	fmt.Println(person1.getUserName())
	fmt.Println(person1.userPassword)

	
}

// function to find the area of a square

func squareArea(length int) int {
	return length * length
}


// function that calculates density

func calcDensity(mass,volume float32) float32{
	return mass/volume
}

// function to return the username from a struct
func (user User) getUserName() (string){
	return user.UserName
}

// function to change the username

func (u *User) changeUserName(newName string) {
	u.UserName = newName
}

func (u *User) changePassword(newPass string) {
	u.userPassword = newPass
}