package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	fmt.Println("serve run in http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
