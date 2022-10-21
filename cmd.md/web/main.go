package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/teaching", Teaching)
	http.HandleFunc("/entreneurship", Entreneurship)

	fmt.Println(fmt.Printf("Starting Application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
