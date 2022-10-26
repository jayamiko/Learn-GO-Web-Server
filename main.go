package main

import (
	"fmt"
	"net/url"
)


func main() {
	var urlString = "https://jsonplaceholder.typicode.com/users?name=Jaya%20Miko"
	var u, e = url.Parse(urlString)
    if e != nil {
        fmt.Println(e.Error())
return
    }
    fmt.Printf("url: %s\n", urlString)
    fmt.Printf("protocol: %s\n", u.Scheme) // https
    fmt.Printf("host: %s\n", u.Host)       // jsonplaceholder.typicode.com
    fmt.Printf("path: %s\n", u.Path)       // /users
    var name = u.Query()["name"][0] // name: Leanne Graham
    fmt.Printf("name: %s", name) 
}