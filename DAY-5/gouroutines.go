package main

import (
	"fmt"
	"time"
)

/// Goroutines are lightweight threads of execution in Go.
//They allow us to perform concurrent operations concurrently,
///with minimal overhead. Goroutines are created using the go keyword followed by a function call or function literal.



func helloWorld()  {
	fmt.Println("Hello World")
}


func getNumbers()  {

	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	
}

func main()  {
	go helloWorld() // Create a new goroutine
	fmt.Println("Today") // this "main gouroutine is executed"
	time.Sleep(time.Second) // pauses the main go routine to execute the "helloWorld goroutine"


	//example 2
	go getNumbers() // goroutine created
	time.Sleep(time.Minute) // this pauses the main goroutine until the "getNumbers" goroutine is executed
	fmt.Println("starting") // <--main goroutine
	
}