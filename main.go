package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Web Service API adalah sebuah web yang menerima request dari client dan menghasilkan response, biasa berupa JSON/XML.

type student struct {
	ID    string
	Name  string
	Age   int
	Title string
}

// Json Data
var data = []student{
	student{"1", "Jaya Miko", 24, "Programmer"},
	student{"2", "Leanne Graham", 24, "Designer"},
	student{"3", "John Wick", 24, "Supervisor"},
	student{"4", "Ethan Bond", 24, "HRD"},
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

// Mengembalikan semua data yang ada di array
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

// Mengembalikan satu data, diambil dari data berdasarkan ID-nya
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id := r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}
