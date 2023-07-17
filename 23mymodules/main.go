package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	fmt.Println("Hello MOD")
	r := mux.NewRouter();
	r.HandleFunc("/home", serveHome).Methods("GET")
	
	log.Fatal(http.ListenAndServe("http://localhost:8000", r))
}


func serveHome(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("<h1>Welcome to Golang Learningsss....</h1>"))
}