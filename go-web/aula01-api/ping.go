package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("Aplicação de API em GO")
	
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8080", nil)
}
