// Parsing json file into structured data  
package main

import (
	"fmt"
	"encoding/json"
    "io/ioutil"
    "os"
    "strconv"
)

// Users struct which contains
// an array of users
type Users struct {
    Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("./users.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()

    var users Users
    byteValue, _ := ioutil.ReadAll(jsonFile)

    json.Unmarshal(byteValue, &users)

    for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
        fmt.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Social URLs: ", users.Users[i].Social)
        fmt.Printf("\n\n")
    }

}