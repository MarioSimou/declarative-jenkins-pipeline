package main

import (
	"net/http"
	"fmt"
	i "../../internal"
)

func ping(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World\n")
}

func hello(w http.ResponseWriter, r *http.Request){
	var queries = r.URL.Query()
	var name = queries.Get("name")

	if name == "" {
		name = "John"
	}

	fmt.Fprintf(w, "%s\n", i.GetGreeting(name))
}

func main() {
	var address = ":3000"
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", hello)

	fmt.Printf("The app listens on %s\n", address)
	http.ListenAndServe(address, nil)
}