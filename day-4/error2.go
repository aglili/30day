package main

import "fmt"

//creating custom errors

type CustomError struct{
	message string

}

// function to return error mesage 

func (e *CustomError) Error() string {
	return e.message
}


//function to process the data 

func processs(data []int) (error,int){
	if len(data) == 0{
		return &CustomError{"empty data"},0
	}

	// process the data
	sum := 0
	for _,number := range data{
		sum += number
	}
	return nil,sum

}

func main()  {

	var data = []int{}

	err,result := processs(data)

	if err != nil{
		fmt.Printf("Error: %v ",err)
	}else{
		fmt.Printf("Result: %v",result)
	}
	
}