package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello!\r\n"))
	})
	log.Println("Listening on port 8080.....")
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, sm))
}