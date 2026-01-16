package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct que representa o usuário
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	fmt.Println("Iniciando servidor na porta 8080")

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	})

	http.HandleFunc("/welcome", Welcome)

	http.ListenAndServe(":8080", nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	// a) Verifica se é POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
		return
	}

	// b) Lê o body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao ler o body"))
		return
	}

	// c) Cria variável user e d) Converte JSON para struct
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON inválido"))
		return
	}

	// e) Monta a frase
	response := fmt.Sprintf("Boas vindas %s, o que você mais gosta do seu %s?", user.Name, user.Country)

	// f) Escreve a resposta
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
