package main

import (
	"fmt"
	"log"
	"net/http"
)

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
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
