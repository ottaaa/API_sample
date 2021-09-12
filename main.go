package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", handler)
	// 3001ポートで起動
	http.ListenAndServe(":3001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
