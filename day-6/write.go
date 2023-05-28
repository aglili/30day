package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte("This is written to the file again")

	err := ioutil.WriteFile("day-6/write.txt",data,0644)

	//0644 value represents the file permissions when using the ioutil.WriteFile function to write data to a file.

	if err != nil{
		fmt.Printf("Error: %v ",err)
		return 
	}

	fmt.Println("Writing to file sucessful")
}