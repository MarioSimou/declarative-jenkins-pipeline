package main

import (
	"net/http"
	"fmt"
	i "../../internal"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, i.GetHello())
}

func main(){
	var address = ":3000"

	http.HandleFunc("/", index)

	fmt.Printf("The app is listening on %s\n", address)
	http.ListenAndServe(address, nil)
}