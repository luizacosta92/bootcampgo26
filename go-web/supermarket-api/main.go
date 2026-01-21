package main

import (
	"encoding/json" // 📦 O "tradutor" entre Go e JSON
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// 🚀 FUNÇÃO PRINCIPAL - O "maestro" que inicia tudo

func main() {
	// 1. Carrega os produtos do JSON quando o servidor inicia
	fmt.Println("📚 Carregando produtos...")
	if err := loadProducts(); err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}

	// 2. Registra as rotas (como colocar placas indicativas)
	router := gin.Default()
	Routes(router)
	// 3. Inicia o servidor
	fmt.Println("🚀 Servidor rodando em http://localhost:8080")
	fmt.Println("📍 Rotas disponíveis:")
	fmt.Println("   GET    /products")
	fmt.Println("   GET    /products/:id")
	fmt.Println("   POST   /products")
	fmt.Println("   PUT    /products/:id")
	fmt.Println("   DELETE /products/:id")

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("❌ Erro ao iniciar servidor: %v\n", err)
	}
}

// 📚 VARIÁVEL GLOBAL:
// Quando o servidor inicia, carregamos tudo do JSON aqui
var products []Product

// 🔑 Funções auxiliares

// loadProducts carrega os produtos do arquivo JSON
// É como abrir a estante e ler todos os livros de uma vez
func loadProducts() error {
	// 1. Abre o arquivo (como abrir uma gaveta)
	file, err := os.Open("products.json")
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close() // Sempre feche a gaveta quando terminar! 🚪

	// 2. Lê todo o conteúdo (como ler todas as páginas)
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo: %w", err)
	}

	// 3. MÁGICA DO JSON! 🎩✨
	// json.Unmarshal = "Desempacotar" JSON → Go structs
	// É como o tradutor pegando o texto em JSONês e falando em Golês
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		return fmt.Errorf("erro ao fazer unmarshal: %w", err)
	}

	fmt.Printf("✅ %d produtos carregados com sucesso!\n", len(products))
	return nil
}

// saveProducts salva os produtos de volta no arquivo JSON
// É como guardar os livros atualizados na estante
func saveProducts() error {
	// 1. MÁGICA REVERSA! 🎩✨
	// json.MarshalIndent = "Empacotar" Go structs → JSON bonito
	// MarshalIndent deixa formatado (com indentação)
	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao fazer marshal: %w", err)
	}

	// 2. Escreve no arquivo (guarda na gaveta)
	err = os.WriteFile("products.json", data, 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever arquivo: %w", err)
	}

	return nil
}

// findProductByID busca um produto pelo ID
// Como procurar um livro específico na biblioteca
func findProductByID(id int) (*Product, int) {
	for index, product := range products {
		if product.ID == id {
			return &product, index // Achamos! Retorna o produto e sua posição
		}
	}
	return nil, -1 // Não encontrado 😢
}
