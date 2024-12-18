package tests

import (
	"cmd/main.go/internal/handlers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func RunServer() {
	router := http.NewServeMux()
	handler := handlers.NewHandler()

	router.HandleFunc("/", handler.HandleRequests)

	fmt.Println("Server is running on :3142")
	http.ListenAndServe(":3142", router)
}

func TestMainIntegration(t *testing.T) {
	// Start server
	go func() {
		RunServer()
	}()

	// POST
	postBody := `{"Company":1,"Price":1000,"Origin":"CNSGH","Date":"2023-01-01"}`
	resp, err := http.Post("http://localhost:3142/", "application/json", strings.NewReader(postBody))
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed POST request, error: %v, status: %v", err, resp.StatusCode)
	}

	// GET
	resp, err = http.Get("http://localhost:3142/")
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Failed GET request, error: %v, status: %v", err, resp.StatusCode)
	}

	// check response
	expectedResponse := map[string]int{"CNSGH": 1000}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var actualResponse map[string]int
	if err := json.Unmarshal(body, &actualResponse); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if !reflect.DeepEqual(expectedResponse, actualResponse) {
		t.Errorf("Expected response %v, got %v", expectedResponse, actualResponse)
	}
}
