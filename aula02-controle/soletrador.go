package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Soletrador() {
	
	leitor := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma palavra: ")
		// 1 - receber a palavra
	palavra, _ := leitor.ReadString('\n')
	palavra = strings.TrimSpace(palavra)


	// 2 - contar o numero de letras
	letras := []rune(palavra) // converte a palavra em um slice de runes

	numeroDeLetras := len(letras) // conta o numero de letras

	fmt.Println("A palavra", palavra, "tem", numeroDeLetras, "letras")
	// 3 - Imprimir cada letra em uma nova linha

	for i, letra := range letras {
		fmt.Printf("Letra %d: %c\n", i+1, letra)
	}
}
