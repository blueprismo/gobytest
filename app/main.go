// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	var store PlayerStore
	server := NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}