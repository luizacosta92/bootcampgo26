package main

import (
	"errors"
	"strings"
	"time"
)

func validateProduct(product Product) error {
	if strings.TrimSpace(product.Name) == "" {
		return errors.New("Nome do produto é obrigatório")
	}
	if product.Quantity <= 0 {
		return errors.New("Quantidade do produto deve ser maior que 0")
	}
	if product.CodeValue == "" {
		return errors.New("Código do produto é obrigatório")
	}
	if product.Expiration == "" {
		return errors.New("Data de validade é obrigatória")
	}
	if product.Price <= 0 {
		return errors.New("Preço do produto deve ser maior que 0")
	}
	return nil
}

func codeValueIsUnique(codeValue string, excludeID int) bool {
	for _, product := range products {
		if product.CodeValue == codeValue && product.ID != excludeID {
			return false
		}
	}
	return true
}
func validateExpirationDate(expiration string) error {
	expiration = strings.TrimSpace(expiration)

	//Pegar a data e verificar se esta no formato DD/MM/AAAA
	_, err := time.Parse("02/01/2006", expiration)
	if err != nil {
		return errors.New("Data de validade deve estar no formato DD/MM/AAAA")
	}
	return nil
}
