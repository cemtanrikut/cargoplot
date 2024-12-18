package main

import (
	"cmd/main.go/internal/handlers"
	"fmt"
	"net/http"
)

func RunServer() {
	router := http.NewServeMux()
	handler := handlers.NewHandler()

	router.HandleFunc("/", handler.HandleRequests)

	fmt.Println("Server is running on :3142")
	http.ListenAndServe(":3142", router)
}

func main() {
	RunServer()
}
