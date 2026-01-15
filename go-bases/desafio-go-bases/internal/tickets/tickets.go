package tickets

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	Time        string
	Price       float64
}

// Função auxiliar para ler o CSV e retornar as linhas
func readCSVLines() ([]string, error) {
	data, err := os.ReadFile("tickets.csv")
	if err != nil {
		return nil, errors.New("error reading file")
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

// isInPeriod verifica se a hora pertence ao período
func isInPeriod(hour int, period string) bool {
	switch strings.ToLower(period) {
	case "madrugada":
		return hour >= 0 && hour <= 6
	case "manha", "manhã":
		return hour >= 7 && hour <= 12
	case "tarde":
		return hour >= 13 && hour <= 19
	case "noite":
		return hour >= 20 && hour <= 23
	default:
		return false
	}
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	var total int = 0

	lines, err := readCSVLines()
	if err != nil {
		return 0, err
	}

	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) < 6 {
			continue
		}
		if strings.EqualFold(fields[3], destination) {
			total++
		}
	}
	return total, nil
}

// GetCountByPeriod conta tickets por período do dia
// Períodos: "madrugada" (0-6), "manhã" (7-12), "tarde" (13-19), "noite" (20-23)
func GetCountByPeriod(period string) (int, error) {
	var total = 0

	lines, err := readCSVLines()
	if err != nil {
		return 0, err
	}

	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) < 6 {
			continue
		}

		// Extrai a hora do formato "17:11"
		timeParts := strings.Split(fields[4], ":")
		if len(timeParts) < 2 {
			continue
		}

		hour, err := strconv.Atoi(timeParts[0])
		if err != nil {
			continue
		}

		// Verifica se a hora está no período solicitado
		if isInPeriod(hour, period) {
			total++
		}
	}
	return total, nil
}

func AverageDestination(destination string, total int) (float64, error) {
	ticketsForDestination, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if total == 0 {
		return 0, errors.New("total cannot be zero")
	}
	return float64(ticketsForDestination) / float64(total), nil
}
