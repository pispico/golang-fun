package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/teste", handler2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handler 2 = %q\n", r.Method)
	fmt.Fprintf(w, "Handler 2 = %q\n", r.Host)
	fmt.Fprintf(w, "Handler 2 = %q\n", r.Proto)
	fmt.Fprintf(w, "Handler 2 = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Handler 2 = %q\n", r.RequestURI)
	fmt.Fprintf(w, "Handler 2 = %q\n", r.URL)
}
