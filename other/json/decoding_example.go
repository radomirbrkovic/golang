package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func main()  {
	url := "https://xkcd.com/info.0.json"
	result := getJson(url)

	fmt.Printf("%v\n", result)
}


func getJson(url string) map[string]interface{} {
    r, err := http.Get(url)
	if err != nil {
        panic(err)
    }
	defer r.Body.Close()
	byteValue, _ := ioutil.ReadAll(r.Body)
	var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

	return result
}