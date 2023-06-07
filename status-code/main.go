package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	// This file does not exist.
	resp, _ := http.Get("https://www.dotnetperls.com/robots.txt")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	// Test status string.
	fmt.Println("Status:", resp.Status)
	if strings.Contains(resp.Status, "404") {
		fmt.Println("Is not found (string tested)")
	}

	// Test status code of the response.
	fmt.Println("Status code:", resp.StatusCode)
	if resp.StatusCode == 404 {
		fmt.Println("Is not found (int tested)")
	}
}
