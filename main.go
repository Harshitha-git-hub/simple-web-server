package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Printf("starting server at port 8020\n")
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
