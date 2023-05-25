package main

import (
	"fmt"
	"math"
	"net"
) // math package has been imported


func main(){

	answer := math.Cos(90) * math.Log(1)

	fmt.Println(answer) // returns the result of the cosine of 90 x the log of 1

	_,conn :=net.Dial("tcp","golang.org:http")

	fmt.Println(conn)
	
}