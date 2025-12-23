package main

import (
	"fmt"
	"net/http"
)

func main() {
	serverPort := "8080"
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
