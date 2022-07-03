package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		name := os.Getenv("NAME")
		age := os.Getenv("AGE")

		fmt.Fprintf(w, "Hello I'm %s and I'm %s years old", name, age)

		w.Write([]byte("Hello, Kubernetes!!"))
	})
	http.ListenAndServe(":8080", nil)
}
