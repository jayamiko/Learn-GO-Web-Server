package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "How are You?")
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		var data = map[string]string{
			"Name": "Jaya Miko",
			"Message": "Have a Nice Day!",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})
	// http.HandleFunc("/index", index)

	fmt.Println("Starting web Server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}