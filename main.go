package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A struct to map our Pokemon's Species which includes it's name
type Users struct {
	ID       string `json:"id"`
	Name     string
	Username string
	Email    string
}

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []Users
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject)

	for i := 0; i < len(responseObject); i++ {
		fmt.Println(i, responseObject[i].Email)
	}
}
