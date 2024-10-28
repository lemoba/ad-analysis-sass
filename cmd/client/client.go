package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data []map[string]any `json:"data"`
}

func main() {
	userInfo := map[string]any{
		"name": `ranen1024`,
		"age":  25,
		"city": "hangzhou",
	}

	reqBody, _ := json.Marshal(userInfo)

	resp, _ := http.Post("http://127.0.0.1:8080/", "application/json", bytes.NewBuffer(reqBody))

	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var result Response

	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, v := range result.Data {
		for key, value := range v {
			fmt.Println("key:", key, "value:", value)
		}
	}
	fmt.Println(result.Data)
}
