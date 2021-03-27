// This is example of rendering simple html page

package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request)  {
	request := r.URL.Query()

	if (request["name"] != nil) {
		fmt.Fprintf(w, "Welcome " + request["name"][0] +"!")
	}
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func handleRequests() {
	http.HandleFunc("/", home)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}