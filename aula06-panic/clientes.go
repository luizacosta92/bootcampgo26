package main

import (
	"fmt"
	"os"
)

func readTxt() {
	//um defer para imprimir a mensagem de execução concluída
	defer fmt.Println("Execução concluída\n")

	//e outrodefer para capturar o erro e imprimir a mensagem de erro
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro capturado:", r)
		}
	}()

	//usa biblioteca "os" para ler o arquivo "customers.txt"
	content, err := os.ReadFile("aula06-panic/customers.txt")

	//se der erro, vai ter um panic > que vai prum recover > e vai imprimir a mensagem de erro
	if err != nil {
		panic("The indicated file was not found")
	}

	// Imprime o conteúdo
	fmt.Println("\n📄 Conteúdo do arquivo:")
	fmt.Println(string(content))
}

func main() {
	fmt.Println("=== AULA 06 ===")
	readTxt()
}
