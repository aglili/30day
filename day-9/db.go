package main

// run go get -u github.com/lib/pq to install postgresql drivers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)



type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Balance float64 `json:"balance"`
}
func checkErr(err error){
	if err != nil{
		fmt.Printf("Error: %v",err)
		return
	}
}


func postgresConnect() (*sql.DB,error){
	db ,err := sql.Open("postgres","postgresql://postgres:0244052298aglili@db.adhqafefzfxmdmcxsokd.supabase.co:5432/postgres")
	checkErr(err)

	//test connection
	err = db.Ping()

	checkErr(err)
	return db,nil
}


func main()  {

	http.HandleFunc("/add/",addUser)


	log.Fatal(http.ListenAndServe(":8000",nil))
	
}


func addUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")


	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	checkErr(err)

	db,err := postgresConnect()
	checkErr(err)
	defer db.Close()

	query := "INSERT INTO users (name,email,balance) VALUES ($1,$2,$3) RETURNING id"

	var userID int

	err = db.QueryRow(query,user.Username,user.Email,user.Balance).Scan(&userID)

	checkErr(err)

	w.WriteHeader(http.StatusCreated)
	response := map[string]int{"id":userID} 

	json.NewEncoder(w).Encode(response)

}


