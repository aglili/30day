package main

import (
	"fmt"
	"io/ioutil" // standard package library for reading and writing to files
)

func main() {
	data, err := ioutil.ReadFile("day-6/hello.txt") // readfile returns a response in form of bytes and an error message incase one occurs

	if err != nil{
		fmt.Printf("Error Reading file: %v",err)
		return
	}

	fmt.Println("File Content: ")
	fmt.Println(string(data))


}