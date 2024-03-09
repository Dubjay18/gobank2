package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "hello world!!")
	})
	http.ListenAndServe(":8080", nil)
}
