package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Simple app running...")
	msg := os.Getenv("SIMPLE_MSG")
	if msg == "" {
		msg = ":( SIMPLE_MSG variable not defined"
	}
	fmt.Fprintf(w, "<h1>%s again...</h1>", msg)
}

func main() {
	flag.Parse()
	log.Print("Simple app server started...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
