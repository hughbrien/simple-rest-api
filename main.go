package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{'first_name':'Hugh '")
	})
	if err := http.ListenAndServe("localhost:8090", mux); err != nil {
		fmt.Println(err.Error())
	}
}
