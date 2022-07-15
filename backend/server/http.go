package main

import (
	"fmt"
	"net/http"
)

func setupRoutes1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "this is server")
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":3100", nil)

}
