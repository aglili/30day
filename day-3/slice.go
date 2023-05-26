package main

import (
	"fmt"
)

func main() {
	// slices and arrays are used for storing and manipulating elements

	//arrays have a fixed size which is alocated when the array is initialized
	//Example
	var names = [5]string{"jake", "mark", "eliza"} // this creates an array which can hold 5 elements

	// the elements off an array can be accesed by indexing

	fmt.Println(names[0]) //returns the first element in the array ==> "jake"

	//slices are dynamic so you can add elements as you go along
	// initialization
	var fruits = []string{"pineapple","cherry"}

	//to add to a string we use the "append"
	fruits = append(fruits,"apple" )  // adds the element apple to the slice "fruits"  

	//methods can be performed on strings
	numFruits := len(fruits)  // find the length of the array or slice


	//looping through slices with "range" . returns the index and the value of the particular iteration
	for index,fruit := range fruits{ 
		fmt.Printf("%v:%v \n",index,fruit)
	}


	// same way of doing the above
	for i:=0;i<len(fruits);i++ {
		fmt.Printf("%v:%v \n",i,fruits[i])
	}


	fmt.Printf("there are %v fruits",numFruits)
}