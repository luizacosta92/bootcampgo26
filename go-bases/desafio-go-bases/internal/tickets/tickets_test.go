package tickets

import (
	"os"
	"testing"
)

// Dado uma tabela de tickets, quando recebe um pais, deve retornar o total de tickets para esse pais
func TestGetTotalTickets(t *testing.T) {
	// Muda para o diretório raiz do projeto onde está o tickets.csv
	os.Chdir("../../")

	total, err := GetTotalTickets("Brazil")
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if total != 45 {
		t.Errorf("Esperado 45, obteve %d", total)
	}
}

//Dado uma tabela de tickets, quando recebe um periodo, deve retornar o total de tickets para esse periodo
func TestGetCountByPeriod(t *testing.T) {
	os.Chdir("../../")

	total, err := GetCountByPeriod("manhã")
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if total != 256 {
		t.Errorf("Esperado 256, obteve %d", total)
	}
}

//Dado uma tabela de tickets, quando recebe um pais e um periodo, deve retornar a média de tickets para esse pais e periodo
func TestAverageDestination(t *testing.T) {
	os.Chdir("../../")

	average, err := AverageDestination("Brazil", 15)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if average != 3 {
		t.Errorf("Esperado 3, obteve %.2f", average)
	}
}