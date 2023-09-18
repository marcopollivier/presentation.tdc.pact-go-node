package main

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"regexp"
)

type User struct {
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

var users []User

func main() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
			return
		}

		if !IsValidEmail(user.Email) {
			http.Error(w, "Email inválido", http.StatusBadRequest)
			return
		}

		users = append(users, user)

		response := ResponseMessage{
			Message: "User created successfully!",
		}

		responseData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Erro ao codificar a resposta", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseData)
		return
	}

	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return false
	}

	rgx := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return rgx.MatchString(email)
}
