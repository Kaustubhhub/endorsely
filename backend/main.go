package main

import (
	"fmt"
	"net/http"

	"github.com/kaustubhhub/endorsely/routes"
)

func main() {
	fmt.Println("In main function")
	http.HandleFunc("/health", healthServer)      // Correct usage for handler registration
	http.HandleFunc("/api/signup", routes.Signup) // Pass the handler function reference
	http.HandleFunc("/api/login", routes.Login)   // Pass the handler function reference
	http.ListenAndServe(":8080", nil)
}

func healthServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Health check of server")
	fmt.Fprintf(w, "Server is running fine!")
}
