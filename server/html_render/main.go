// This is example of rendering simple html page

package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Student struct {
	Name string
	College string
}


func home(w http.ResponseWriter, r *http.Request)  {
	request := r.URL.Query()

	if (request["name"] != nil) {
		fmt.Fprintf(w, "Welcome " + request["name"][0] +"!")
	}
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func temp(w http.ResponseWriter, r *http.Request)  { 
	student := Student{
        Name:       "John Doe",
        College:    "Stanford University",
    }

	parsedTemplate, _ := template.ParseFiles("template.html")
    err := parsedTemplate.Execute(w, student)
    if err != nil {
        log.Println("Error executing template :", err)
        return
    }
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/template", temp)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}