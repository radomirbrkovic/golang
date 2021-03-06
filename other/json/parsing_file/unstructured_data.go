//JSON parse file - Unstructured data 

package main

import (
	"fmt"
	"encoding/json"
    "io/ioutil"
    "os"
)

func main()  {
	// Open our jsonFile
	jsonFile, err := os.Open("./users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		 fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    fmt.Println(result["users"])
}