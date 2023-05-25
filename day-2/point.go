package main

import "fmt"

func main() {
	//  a pointer is a variable that stores the memory address of another variable

	var name string = "jake"

	pointer := &name // "pointer" now stores the memory location of the variable name

	// to get the value of the "pointer" variable
	// we dereference it with the asterisk *

	fmt.Println(*pointer) // == "jake"


	///////

	var age *int // initialize a pointer

	age = new(int) // allocate memory 

	*age = 50 // store a value in the memory address 

	fmt.Println(*age)

}