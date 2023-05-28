package main

import (
	"fmt"
	"log"
	"net/http"
)

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form error : %v", err)
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Post request succesfull \n")

	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Form submitted with name %v and email %v", name, email)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/subscribe", SubscribeHandler)

	fmt.Println("Starting server at Port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
