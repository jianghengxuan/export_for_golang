package main

import (
	"net/http"
	"testing"
)

func TestHelloHandle(t *testing.T) {
	http.HandleFunc("/hello", helloHandle)

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	go server.ListenAndServe()

	// Here we use the http.Get() function to send the request.
	// We also use the t.Errorf() function to report any errors
	// that occur during the test.

	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		t.Errorf("Error sending request: %v", err)
	}
}
