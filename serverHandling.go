package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)


// callback handling
func CALLBACK_HANDLER(w http.ResponseWriter, r *http.Request) {
    // Print request method
    fmt.Println("Request Method:", r.Method)

    // Print request headers
    fmt.Println("Request Headers:")
    for key, values := range r.Header {
        for _, value := range values {
            fmt.Printf("%s: %s\n", key, value)
        }
    }

    // Read the request body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

    // Example: Print the request body
    fmt.Println("Received callback:")
    fmt.Println(string(body))

    // Send a response (optional)
    w.WriteHeader(http.StatusOK)
	curtime := time.Now()
    fmt.Fprintf(w, "Callback received successfully: %s", curtime)
}

// local server
func HTTP_SERVER() {
    // Define the callback URL endpoint
    callbackURL := "/webhook/callback"

    // Register the callback handler function
    http.HandleFunc(callbackURL, CALLBACK_HANDLER)

    // Start the HTTP server
    port := ":8081" // Change this to your desired port
	host := "127.0.0.1"
    fmt.Printf("Listening on %s%s\n", host, port)
    if err := http.ListenAndServe(host+port, nil); err != nil {
        fmt.Printf("Server error: %s\n", err)
    }
}

