package main


import "fmt"


//Synchronization with Channels:
//Channels are used for communication and synchronization between goroutines. 
//They provide a safe way to exchange data and coordinate the execution of goroutines. 



func toChannel(c chan int){
	c <- 10  // the value is sent into the channel
}

func main()  {
	c := make(chan int) // Create an unbuffered channel
    go toChannel(c) // Create a new goroutine to send value to the channel
    value := <-c // Receive the value from the channel
    fmt.Println(value) // Print the received value
}