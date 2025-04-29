package auth

import (
	"fmt"
	"net/http"
)

// SignUpHandler Handles user signup requests - placeholder
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Signup endpoint hit!")
}

// LoginHandler Handles user login requests - placeholder
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login endpoint hit!")
}
