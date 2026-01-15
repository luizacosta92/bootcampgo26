package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Person struct {
	Name    string `json:"firstName"`
	Surname string `json:"lastName"`
}

func main() {
	fmt.Println("Iniciando servidor na porta 8080")
	http.HandleFunc("/greetings", Greetings)
	http.ListenAndServe(":8080", nil)
}

func Greetings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading body"))
		return
	}

	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error unmarshalling body"))
		return
	}

	response := fmt.Sprintf("Hello %s %s", person.Name, person.Surname)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))	
}
