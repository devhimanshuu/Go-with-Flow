package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Module in Golang")
	Greeter()
	r := mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000",r))
}

func Greeter(){
	fmt.Println("hey there mod users")
}

func serveHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcome To Golang</h1>"))
}