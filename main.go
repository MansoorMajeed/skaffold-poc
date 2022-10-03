package main

import (
	"fmt"
	"log"
	"net/http"
)

var appVersion string = "v2"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, Path : %s! \n", r.URL.Path[1:])
	fmt.Fprintf(w, "This is version : %s", appVersion)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
