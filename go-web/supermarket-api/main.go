package main

import (
	"encoding/json" // 📦 O "tradutor" entre Go e JSON
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 📚 VARIÁVEL GLOBAL: Nossa "biblioteca" de produtos na memória
// Quando o servidor inicia, carregamos tudo do JSON aqui
var products []Product

// 🔑 Funções auxiliares - Nossos "ajudantes da biblioteca"

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

// 🌐 HANDLERS - Os "atendentes" que respondem às requisições HTTP

// pingHandler - Simples verificação se o servidor está vivo
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "pong"}`))
}

// getAllProductsHandler - GET /products (Listar todos)
// Como pedir ao bibliotecário: "Me mostra todos os livros!"
func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Marshal = transformar struct Go em JSON
	data, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Erro ao processar produtos"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// getProductByIDHandler - GET /products/:id (Buscar um específico)
// Como pedir: "Me mostra o livro número 5!"
func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extrai o ID da URL: /products/5 → "5"
	// Usamos strings.TrimPrefix para remover "/products/"
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "ID inválido"}`))
		return
	}

	product, _ := findProductByID(id)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Produto não encontrado"}`))
		return
	}

	data, _ := json.Marshal(product)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// createProductHandler - POST /products (Criar novo)
// Como entregar um livro novo para o bibliotecário catalogar
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. Lê o corpo da requisição (o JSON que vem do Postman)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Erro ao ler corpo da requisição"}`))
		return
	}

	// 2. Unmarshal = converte JSON → struct Go
	var newProduct Product
	err = json.Unmarshal(body, &newProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "JSON inválido"}`))
		return
	}

	// 3. Gera um novo ID (pega o maior ID + 1)
	maxID := 0
	for _, p := range products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	newProduct.ID = maxID + 1

	// 4. Adiciona à lista
	products = append(products, newProduct)

	// 5. Salva no arquivo
	if err := saveProducts(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Erro ao salvar produto"}`))
		return
	}

	// 6. Retorna o produto criado
	data, _ := json.Marshal(newProduct)
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// updateProductHandler - PUT /products/:id (Atualizar existente)
// Como pedir para atualizar as informações de um livro
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "ID inválido"}`))
		return
	}

	_, index := findProductByID(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Produto não encontrado"}`))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	err = json.Unmarshal(body, &updatedProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "JSON inválido"}`))
		return
	}

	updatedProduct.ID = id // Garante que o ID não mude
	products[index] = updatedProduct

	if err := saveProducts(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(updatedProduct)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// deleteProductHandler - DELETE /products/:id (Deletar)
// Como pedir para remover um livro da biblioteca
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "ID inválido"}`))
		return
	}

	_, index := findProductByID(id)
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Produto não encontrado"}`))
		return
	}

	// Remove o produto da slice (técnica comum em Go)
	products = append(products[:index], products[index+1:]...)

	if err := saveProducts(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 = sucesso sem conteúdo
}

// productsRouter - Roteador que decide qual handler chamar
// Como a recepcionista que direciona visitantes para diferentes departamentos
func productsRouter(w http.ResponseWriter, r *http.Request) {
	// Verifica se é /products exato ou /products/123
	if r.URL.Path == "/products" {
		switch r.Method {
		case http.MethodGet:
			getAllProductsHandler(w, r)
		case http.MethodPost:
			createProductHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else if strings.HasPrefix(r.URL.Path, "/products/") {
		switch r.Method {
		case http.MethodGet:
			getProductByIDHandler(w, r)
		case http.MethodPut:
			updateProductHandler(w, r)
		case http.MethodDelete:
			deleteProductHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// 🚀 FUNÇÃO PRINCIPAL - O "maestro" que inicia tudo
func main() {
	// 1. Carrega os produtos do JSON quando o servidor inicia
	fmt.Println("📚 Carregando produtos...")
	if err := loadProducts(); err != nil {
		fmt.Printf("❌ Erro: %v\n", err)
		return
	}

	// 2. Registra as rotas (como colocar placas indicativas)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/products", productsRouter)
	http.HandleFunc("/products/", productsRouter) // Com barra para capturar /products/:id

	// 3. Inicia o servidor
	fmt.Println("🚀 Servidor rodando em http://localhost:8080")
	fmt.Println("📍 Rotas disponíveis:")
	fmt.Println("   GET    /ping")
	fmt.Println("   GET    /products")
	fmt.Println("   GET    /products/:id")
	fmt.Println("   POST   /products")
	fmt.Println("   PUT    /products/:id")
	fmt.Println("   DELETE /products/:id")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("❌ Erro ao iniciar servidor: %v\n", err)
	}
}
