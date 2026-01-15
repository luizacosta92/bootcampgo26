package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	defer fmt.Println("Execução concluída")
	reader := bufio.NewReader(os.Stdin)
	var destination string

	fmt.Println("Informe o país de destino: ")
	destination, _ = reader.ReadString('\n')
	destination = strings.TrimSpace(destination)


	total, err := tickets.GetTotalTickets(destination)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Printf("Total de tickets para %s: %d\n", destination, total)

	var period string
	fmt.Println("Informe o período do dia: ")
	fmt.Scan(&period)
	
	total, err = tickets.GetCountByPeriod(period)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Printf("Total de tickets para %s: %d\n", period, total)

	average, err := tickets.AverageDestination(destination, total)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Printf("Média de tickets para %s: %.2f\n", destination, average)
}
