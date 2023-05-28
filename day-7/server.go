package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/",helloFunc) // Define the route and corresponding handler function for GET requests
	http.HandleFunc("/about",aboutFunc)
	http.HandleFunc("/user",getUser)


	log.Fatal(http.ListenAndServe(":8000",nil)) // Start the server on port 8080

}



func helloFunc(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprintf(w,"Hello World")
	
}


func aboutFunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"This is the about page of the site")
	
}

func getUser(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case "GET":
		fmt.Fprintf(w,"You Receive a GET Request")
	case "POST":
		fmt.Fprintf(w,"You Sent a POST Request")
	default:
		fmt.Fprintf(w,"Method Not Allowed")
	} 
	
}