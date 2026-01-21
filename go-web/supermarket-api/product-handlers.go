package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Handlers lidam com as requisições HTTP e retornam as respostas

// getAllProductsHandler - GET /products (Listar todos)
func getAllProductsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// getProductByIDHandler - GET /products/:id (Buscar um específico)
func getProductByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	product, _ := findProductByID(id)
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// createProductHandler - POST /products (Criar novo)
func createProductHandler(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// 1. Validar campos obrigatórios
	if err := validateProduct(newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Validar code_value único
	if !codeValueIsUnique(newProduct.CodeValue, 0) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code_value já existe"})
		return
	}

	// 3. Validar data de validade
	if err := validateExpirationDate(newProduct.Expiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 4. Gera um novo ID (pega o maior ID + 1)
	maxID := 0
	for _, p := range products {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	newProduct.ID = maxID + 1

	// 5. Adiciona à lista
	products = append(products, newProduct)

	// 6. Salva no arquivo
	if err := saveProducts(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar produto"})
		return
	}

	c.JSON(http.StatusCreated, newProduct)

}

// updateProductHandler - PUT /products/:id (Atualizar existente)
func updateProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, index := findProductByID(id)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	updatedProduct.ID = id
	products[index] = updatedProduct

	if err := saveProducts(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar produto"})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)

}

// deleteProductHandler - DELETE /products/:id (Deletar)
func deleteProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, index := findProductByID(id)
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	products = append(products[:index], products[index+1:]...)

	if err := saveProducts(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar produto"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
