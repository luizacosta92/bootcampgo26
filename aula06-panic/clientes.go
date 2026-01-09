package main

import (
	"errors"
	"fmt"
	"os"
)

func readTxt() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Execução concluída\n" , r)
		}
	}()
	content, err := os.ReadFile("customers.txt")
	err = errors.New("This indicated file was not found")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}

func main() {
	fmt.Println("=== AULA 06 ===")
	readTxt()
}
