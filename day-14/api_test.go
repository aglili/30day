package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)



func helloHandler(w http.ResponseWriter, r *http.Request){
	username := r.URL.Path[len("/me"):]
	fmt.Fprintf(w,"Hello %v",username)	
}



func testHelloHandler(t *testing.T){

	req,err := http.NewRequest("GET","/me",nil)

	if err != nil{
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr,req)

	
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v",
			rr.Code, http.StatusOK)
	}

	expectedResponse := "Hello,Selorm"
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v, want %v",
			rr.Body.String(), expectedResponse)
	}

}