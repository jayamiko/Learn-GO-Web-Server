package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
	Title    string
}

func main() {
	stringJsonToObject()
	arrayJsonToArrayObject()
	objectToStringJson()
}

func stringJsonToObject() {
	// String JSON to Object

	jsonString := `{"Name": "Jaya Miko", "Age": 24, "Title": "Programmer"}`
	jsonData := []byte(jsonString)

	var data User

	// json.Unmarshal = json string dikonversi menjadi bentuk objek
	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Full Name: ", data.FullName)
	fmt.Println("Age: ", data.Age)
	fmt.Println("Title: ", data.Title)
}

func arrayJsonToArrayObject() {
	// Array JSON to Array Object

	var jsonStringArray = `[
		{"Name": "Jaya Miko", "Age": 24, "Title": "Programmer"},
		{"Name": "Leanne Graham", "Age": 32}
	]`

	var dataArray []User

	e := json.Unmarshal([]byte(jsonStringArray), &dataArray)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println("FullName: ", dataArray[0].FullName) // Jaya Miko
	fmt.Println("Age: ", dataArray[0].Age)           // 24
	fmt.Println("Title: ", dataArray[0].Title)       // Programmer
}

func objectToStringJson() {
	// Object to JSON String

	object := []User{{"Jaya Miko", 24, "Programmer"}, {"Leanne Graham", 27, "Designer"}}

	// json.Marshal = encoding daa ke json string
	jsonDataObject, err := json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStringObject := string(jsonDataObject)
	fmt.Println(jsonStringObject)
}
