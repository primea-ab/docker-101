package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	/*data, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(data))*/
	fmt.Fprintf(w, "Hello, Docker! %s", uuid.NewString())
}

func main() {
	http.HandleFunc("/", handler)
	port := "1337"
	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
