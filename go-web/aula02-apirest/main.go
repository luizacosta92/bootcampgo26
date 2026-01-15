package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("Iniciando servidor na porta 8080")
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", nil)
}