package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculadora() float64 {
	leitor := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o primeiro número: ")
	entrada1, _ := leitor.ReadString('\n')
	entrada1 = strings.TrimSpace(entrada1)
	numero1, _ := strconv.Atoi(entrada1)

	fmt.Print("Digite a operação (+, -, *, /): ")
	operacao, _ := leitor.ReadString('\n')
	operacao = strings.TrimSpace(operacao)

	fmt.Print("Digite o segundo número: ")
	entrada2, _ := leitor.ReadString('\n')
	entrada2 = strings.TrimSpace(entrada2)
	numero2, _ := strconv.Atoi(entrada2)

	switch operacao {
	case "+":
		return float64(numero1 + numero2)
	case "-":
		return float64(numero1 - numero2)
	case "*":
		return float64(numero1 * numero2)
	case "/":
		if numero2 != 0 {
			return float64(numero1) / float64(numero2)
		}
		fmt.Println("Erro: divisão por zero")
		return 0
	default:
		fmt.Println("Operação inválida")
		return 0
	}
}
