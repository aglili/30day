package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

//encode ---> to turn human readble data into a form the computer can understand and use

func main() {

	user_one := User{Name: "Kwame Adade", Age: 34, Email: "adadekwame@gmail.com"}

	user_byte,err :=json.Marshal(user_one)

	if err != nil{
		fmt.Printf("Error: %v",err)
		return
	}


	fmt.Println(string(user_byte))
	fmt.Println("Encoded")

}