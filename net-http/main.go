package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func hello(w http.ResponseWriter, _ *http.Request) {
	type Employee struct {
		Name    string `json:"name"`
		Salary  int    `json:"salary"`
		Married bool   `json:"married"`
	}

	emp := Employee{
		Name:    "User",
		Salary:  1000,
		Married: false,
	}

	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(emp)
	if err != nil {
		return
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				return
			}
		}
	}
}
