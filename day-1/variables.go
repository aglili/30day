package main

import "fmt"

func main() {

	//variables
	// variables are initialised by using the keey word "var"[variable_name] and type
	// example
	var userAge int = 45

	// if the value of the varible is not known before hand, the variable can be decaled and later initialised 
	//example 
	var userName string

	userName = "Jake"
	
	// variables can also be created by inference, that is witout stating the variable type. this is done wth ":="
	//example
	isActive := true
	// go lang inferes the value of this variable, which is a boolean


	fmt.Printf("UserName: %v, UserAge: %v, IsActive: %v\n",userName,userAge,isActive)


	// Constants
	// constants are variables which do not change or are inmutable. they are initialized by the "const" keyboard

	const dateOB = 2001

	// if we try to change the value of dateOB like this "dateOB = 2002" we get an error

}