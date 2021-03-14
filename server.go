package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("listening on 8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:80", http.FileServer(http.Dir("site/"))))
}