// Parsing json file
package main 

import (
	"encoding/json"
    "io/ioutil"
    "os"
	"log"
)
// parse json file and return data
func getData() (map[string]interface {}) {
	jsonFile, err := os.Open("countries.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)
	
	return result
}