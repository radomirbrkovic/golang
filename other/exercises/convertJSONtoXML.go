// Example of reading JSON file from filesystem and creating and writeing XML file 
package main  

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
    "io/ioutil"
)

const INPUT_FILE = "../json/parsing_file/users.json"
const OUTPUT_FILE = "/tmp/user.xml"

type Users struct {
    Users []User `json:"users"`
}

type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

func loadFromJson(key interface{}) error {
    in, err := os.Open(INPUT_FILE)
    defer in.Close()
    if err != nil {
        return err
    }

    decodeJSON := json.NewDecoder(in)
    err = decodeJSON.Decode(key)
    if err != nil {
        return err
    }

    return nil

}

func main() {
    var users Users
    err := loadFromJson(&users)

    if err == nil {
        fmt.Println("JSON:", users)
    } else {
        fmt.Println(err)
    }

    xmlData, _ := xml.MarshalIndent(users, "", "    ")
    err = ioutil.WriteFile(OUTPUT_FILE, xmlData, 0644)

    if err == nil {
        fmt.Println("XML file is created on ", OUTPUT_FILE)
    } else {
        fmt.Println(err)
    }
}