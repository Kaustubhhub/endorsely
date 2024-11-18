package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SignupInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.method : ", r.Method)
	fmt.Println("http.MethodPost : ", http.MethodPost)
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
		return
	}
	var signupBody SignupInput

	err := json.NewDecoder(r.Body).Decode(&signupBody)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("these are the inputs : ", signupBody.Username)
	fmt.Println("these are the inputs : ", signupBody.Password)
	fmt.Fprintf(w, "You are in the signup route")
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "You are in the login route")
}
