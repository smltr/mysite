package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	http.Handle("/",  http.FileServer(http.Dir("site")))
	log.Println("Listening on port 80.....")
	l, err := net.Listen("tcp4", ":80")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, nil))
}