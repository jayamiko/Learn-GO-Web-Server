package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "How are You?")
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello!")
	})
	http.HandleFunc("/index", index)

	fmt.Println("Starting web Server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}