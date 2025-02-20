package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Docker! %s", uuid.NewString())
}

func main() {
	http.HandleFunc("/", handler)
	port := "1337"
	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
