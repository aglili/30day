package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("app.log")

	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)


	log.Println("This is an informational message")
	log.Printf("This is a warning message: %s", "Something went wrong")
	log.Fatalf("This is a fatal error: %s", "Exiting application")
}