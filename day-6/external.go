package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://catfact.ninja/fact"
	response, err := http.Get(url)

	if err != nil{
		fmt.Printf("Error: %v",err)
		return
	}

	defer response.Body.Close()

	// read response from the body

	body,err := ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Printf("Error: %v",err)
		return
	}


	// parse json response 
	var data map[string]interface{}

	err = json.Unmarshal(body,&data)

	if err != nil{
		fmt.Printf("Error: %v",err)
	}

	fmt.Println(data["fact"]) // accessing a specific data from the the json

	// change golang data into json

	payload := map[string]interface{}{
		"name": "Kwadjo Momoni",
		"Adress": "Ofankor",
		"email":"bodie@aol.com",
	}

	jsonPayload,err := json.Marshal(payload)

	if err != nil{
		fmt.Printf("error: %v",err)
		return
	}

	post_url := "https://httpbin.org/post"

	postResponse,err := http.Post(post_url,"application/json",strings.NewReader(string(jsonPayload)))


	if err != nil{
		fmt.Printf("Error: %v",err)
	}
	defer postResponse.Body.Close()

	postBody,_ := ioutil.ReadAll(postResponse.Body)


	fmt.Println(string(postBody))


}