package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for a.com
func helloA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, a.com")
}

// Handler for sub.a.com
func helloSubA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, sub.a.com")
}

// Handler for b.com
func helloB(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, b.com")
}

func main() {
	http.HandleFunc("a.com/", helloA)
	http.HandleFunc("sub.a.com/", helloSubA)
	http.HandleFunc("b.com/", helloB)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
