package main

import (
	// "fmt"
	// "Practice/practice1"
	// "Practice/practice2"
	// "Practice/practice3"
	// "Practice/practice4"
	"Practice/practice5"
	"log"
	"net/http"
	// "github.com/gorilla/mux"
)

func main() {
	port := ":9776"

	log.Fatal(http.ListenAndServe(port, practice5.Mux(port)))
}