package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		name := os.Getenv("NAME")
		age := os.Getenv("AGE")

		fmt.Fprintf(w, "Hello I'm %s and I'm %s years old", name, age)
	})
	http.HandleFunc("/ConfigMap", func(w http.ResponseWriter, r *http.Request) {

		data, err := ioutil.ReadFile("myfamily/family.txt")
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		fmt.Fprintf(w, "My family: %s", string(data))
	})

	http.ListenAndServe(":8080", nil)
}
