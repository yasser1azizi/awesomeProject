/*
package main

import (
	"errors"
	"fmt"
)

func main() {
	port := "8080"
	mux := http.NewServeMux()

	// Example endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!")) //
	})

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("Starting server on port %s...", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
*/
package main

import (
	"errors"
	"fmt"
)

func riskyOperation() error {
	return errors.New("something went wrong")
}

func main() {
	err := riskyOperation()
	if err != nil {
		// Proper error handling
		fmt.Printf("Operation failed: %v\n", err)
		return
	}
	fmt.Println("Operation succeeded")
}
