package main

import (
	"fmt"
	"net/http"
)

func makeRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}

	bs := make([]byte, 999999)

	resp.Body.Read(bs)

	return string(bs)
}
