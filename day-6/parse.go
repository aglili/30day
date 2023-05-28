package main

import (
	"encoding/json"
	"fmt"
)

/// Parsing and Encoding JSON Data

//Go provides the encoding/json package for working with JSON data

// Parse ---> to turn data into human readable form


type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}


func main(){

	jsonData := `{
		"name" : "Kenedy Kuma",
		"age" : 23,
		"email" : "kk@aolonline.com"
	}
	`
	// Parse JSON into a Person struct


	var user User

	err := json.Unmarshal([]byte(jsonData),&user) // Unmarshall takes in data in form of a byte and stores it in the user variable
	if err != nil{
		fmt.Printf("Error: %v",err)
		return
	}

	fmt.Println("Parsed JSON data:")
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Email:", user.Email)
}