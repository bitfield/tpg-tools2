package main

import (
	"fmt"
	"net/http"

	hello "hello3"
)

func main() {
	fmt.Println("Listening on http://localhost:9001")
	http.ListenAndServe(":9001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hello.PrintTo(w)
	}))
}
