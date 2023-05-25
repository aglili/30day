package main

import "fmt"

func main() {
	//control structures determine how a portion of code will run
	// if,for statements are all control statements

	var input int
	var guess int = 10

	
	// for loop
	for i:=0;i<10;i++{
		fmt.Print("Enter The Guess(1-20): ")
		fmt.Scan(&input)
		if (guess==input) {
			fmt.Println("You've Guessed the right number")
			break
		}else {
			fmt.Println("Wrong Guess Try Again")
		}
		
	}



}