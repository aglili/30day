package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id int `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Adress string `json:"adress"`
}

var users []User

func main() {

	user1 := User{Id:1,Name: "jake",Email: "jake@gmail.com",Adress: "Tema Comm 21" }
	user2 := User{Id:3,Name: "Selorm",Email: "selorm@gmail.com",Adress: "Tema Comm 24" }
	user3 := User{Id:2,Name: "lee",Email: "lee@gmail.com",Adress: "Lagos Town" }

	users = append(users, user1,user2,user3)

	fmt.Println(users)

	http.HandleFunc("/users",getUsers)
	http.HandleFunc("/user/create",createUser)
	http.HandleFunc("/user/",getUser)
	http.HandleFunc("/user/delete/",deleteUser)
	http.HandleFunc("/user/update/",updateUser)


	log.Fatal(http.ListenAndServe(":8000",nil))

}


func getUsers(w http.ResponseWriter, r *http.Request){

	// set the content type
	w.Header().Set("content-type","application/json")


	err:= json.NewEncoder(w).Encode(users)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func createUser(w http.ResponseWriter,r *http.Request){
	// set content type
	w.Header().Set("content-type","application/json")

	//create a new user

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil{
		fmt.Fprintf(w,"Error: %v",http.StatusBadRequest)
		return
	}
	//create a user id
	user.Id = len(users) + 1

	users = append(users, user)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)


}

func getUser(w http.ResponseWriter, r *http.Request){
	// set content type
	w.Header().Set("Content-type","application/json")

	//get userid for url path

	user_id := r.URL.Path[len("/user/"):]

	//convert user_id(string) ==> integer

	idNum,err := strconv.Atoi(user_id)

	fmt.Println(idNum)

	if err != nil{
		fmt.Fprintf(w,"Error: %v",err)
		return
	}

	// create targeted user

	var targetUser User

	for _,user := range users{
		if user.Id == idNum{
			targetUser = user
			break
		}
	}

	if targetUser.Id == 0{
		fmt.Fprintf(w,"Error : User Not Found ")
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(targetUser)
	if err != nil{
		fmt.Fprintf(w,"Error: %v",err)
		return
	}

	
}

func deleteUser(w http.ResponseWriter,r *http.Request){
	// set content type
	w.Header().Set("Content-type","application/json")

	user_id := r.URL.Path[len("/user/delete/"):]

	idNum,err := strconv.Atoi(user_id)
	if err != nil{
		fmt.Fprintf(w,"Error: %v",err)
		return
	}


	index := -1
	for i,user := range users{
		if user.Id == idNum{
			index = i
			break
		}
	}

	if index == -1{
		fmt.Fprintf(w,"Error : User Not Found ")
		return
	}

	users = append(users[:index],users[index+1:]... )


	fmt.Fprintf(w,"User Deleted Sucessfully")
}


func updateUser(w http.ResponseWriter, r *http.Request){

	// set content type

	w.Header().Set("Content-type","applecation/json")

	user_id := r.URL.Path[len("/user/update/"):]

	idNum, err := strconv.Atoi(user_id)

	if err != nil {
		fmt.Fprintf(w,"Error",err)
		return
	}


	// find the index of the usr in the slides
	index := -1

	for i,user := range users{
		if user.Id == idNum{
			index = i
			break
		}
	}

	// check if user exists

	if index == -1{
		fmt.Fprintf(w,"User Not Found")
		return
	}

	var updatedUser User

	err = json.NewDecoder(r.Body).Decode(&updatedUser)

	if err != nil {
		fmt.Fprintf(w,"Error: %v",err)
		return
	}

	users[index].Name = updatedUser.Name
	users[index].Adress = updatedUser.Adress
	users[index].Email = updatedUser.Email


	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users[index])
	if err != nil {
		fmt.Fprintf(w,"Error: %v",err)
		return
	}

}