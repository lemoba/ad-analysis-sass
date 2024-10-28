package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Me(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := map[string]any{
		"code": 0,
		"msg":  "success",
		"data": []map[string]any{
			{"name": "ranen", "age": 12},
			{"name": "kkk3", "age": 23},
			{"name": "lemoba", "age": 24},
		},
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	r := httprouter.New()
	r.POST("/", Me)
	fmt.Println("serve run in http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
