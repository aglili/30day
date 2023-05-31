package main

import (
	"crypto/rsa"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var privateKey *rsa.PrivateKey

// Load the private key from a file
func loadPrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile("private_key.pem")
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// database connection
func postgresConnect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:0244052298aglili@db.adhqafefzfxmdmcxsokd.supabase.co:5432/postgres")
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error pinging the database: %v", err)
	}

	return db, nil
}

// Function to register user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
		return
	}

	db, err := postgresConnect()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error connecting to the database: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := "INSERT INTO users(username, password) VALUES($1, $2) RETURNING id"

	var userID int
	err = db.QueryRow(query, user.UserName, user.Password).Scan(&userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing query: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	response := map[string]int{"id": userID}
	json.NewEncoder(w).Encode(response)
}

// Function to validate user credentials against the database
func ValidateCredentials(username, password string) bool {
	db, err := postgresConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := "SELECT COUNT(*) FROM users WHERE username = $1 AND password = $2"

	var count int
	err = db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	return count > 0
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate user
	if !ValidateCredentials(user.UserName, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Load the private key
	privateKey, err := loadPrivateKey()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error loading private key: %v", err), http.StatusInternalServerError)
		return
	}

	// Create a new JWT token with the username as a claim
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"username": user.UserName,
	})

	// Sign the token with the private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error signing token: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": tokenString}
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register", RegisterUser).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
