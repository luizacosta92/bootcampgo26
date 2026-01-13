package main

import (
	"errors"
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

type Client struct {
	File    string
	Name    string
	ID      int
	Phone   string
	Address string
}

var (
	ErrClientExists = errors.New("client already exists")
	ErrEmptyFile    = errors.New("file is empty")
	ErrEmptyName    = errors.New("name is empty")
	ErrInvalidyID   = errors.New("Id must be greater than 0")
	ErrEmptyPhone   = errors.New("phone is empty")
	ErrEmptyAddress = errors.New("address is empty")
)

var clients = []Client{
	{File: "001", Name: "Machado de Assis", ID: 1, Phone: "11999999999", Address: "Rua A"},
	{File: "002", Name: "Carolina", ID: 2, Phone: "11888888888", Address: "Rua B"},
}

func checkClientExists(id int) {

	for _, client := range clients {
		if client.ID == id {
			panic(ErrClientExists)
		}
	}
}

func validateClientData(file, name string, id int, phone, address string) (bool, error) {

	if file == "" {
		return false, ErrEmptyFile
	}
	if name == "" {
		return false, ErrEmptyName
	}
	if id <= 0 {
		return false, ErrInvalidyID
	}
	if phone == "" {
		return false, ErrEmptyPhone
	}
	if address == "" {
		return false, ErrEmptyAddress
	}
	return true, nil
}

func registerClient(file, name string, id int, phone, address string) {
	// Variável para rastrear se houve erro
	hadError := false

	// DEFER: Mensagem de erros (só se houve erro)
	defer func() {
		if hadError {
			fmt.Println("Several errors were detected at runtime")
		}
	}()

	// DEFER: Sempre imprime "End of execution"
	defer fmt.Println("End of execution")

	// DEFER: Recover - captura panic e marca que houve erro
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Erro capturado:", r)
			hadError = true
		}
	}()

	fmt.Println("Registering client...")

	checkClientExists(id)

	valid, err := validateClientData(file, name, id, phone, address)
	if err != nil {
		panic(err)
	}
	if !valid {
		panic(ErrClientExists)
	}

	if valid {
		newClient := Client{File: file, Name: name, ID: id, Phone: phone, Address: address}
		clients = append(clients, newClient)
		fmt.Println("Client registered successfully")
	} else {
		panic(ErrClientExists)
	}
}

func showClients() {
	for _, client := range clients {
		fmt.Println(client)
	}
}

func main() {
	fmt.Println("=== AULA 06 ===")
	//readTxt()
	registerClient("003", "José de Alencar", 3, "11999999999", "Rua C")
	showClients()
}
