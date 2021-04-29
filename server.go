package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/secret" {
		http.Error(rw, "404: This is not the page you are looking for.", http.StatusFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "GET out of here... please.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(rw, "You have found me! I don't know how but you found me.")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	fmt.Printf("Server will be starting soon\n")
	http.HandleFunc("/secret", homeHandler)
	http.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
