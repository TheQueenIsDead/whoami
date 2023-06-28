package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	jsonResponse := `{"hello": "world"}`
	http.HandleFunc("/whoami", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, jsonResponse)
	})

	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code
	//http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
