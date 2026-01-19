package main

type Product struct {
	ID          int     `json:"id"`           // Campo público (maiúscula) + tag JSON
	Name        string  `json:"name"`         // Go: Name → JSON: name
	Quantity    int     `json:"quantity"`     // Quantidade em estoque
	CodeValue   string  `json:"code_value"`   // Código do produto
	IsPublished bool    `json:"is_published"` // Se está disponível
	Expiration  string  `json:"expiration"`   // Data de validade
	Price       float64 `json:"price"`        // Preço em reais
}
