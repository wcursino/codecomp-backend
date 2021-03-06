package main

import (
	"codecomp-backend/routers"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/login", login).Methods("POST")
	myRouter.HandleFunc("/register", register).Methods("POST")
	myRouter.HandleFunc("/code", register).Methods("POST")
	routers := routers.Init()
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(routers)))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}

func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you just hit this endpoint yo")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("request body", string(reqBody))
	fmt.Fprintf(w, "%+v", string(reqBody))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you just hit this endpoint yo")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("request body", string(reqBody))
	fmt.Fprintf(w, "%+v", string(reqBody))
}

func runCode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("you just hit this endpoint yo")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("request body", string(reqBody))
	fmt.Fprintf(w, "%+v", string(reqBody))
}
