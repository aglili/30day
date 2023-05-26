package main

import (
	"errors"
	"fmt"
)

func main() {

	//basic error checking

	result,err := divide(10,5)
	if err != nil{
		fmt.Printf("Error: %v",err)
	}else{
		fmt.Printf("Result: %v",result)
	}

}






func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0,errors.New("Division by Zero")
	}
	return a/b,nil
}