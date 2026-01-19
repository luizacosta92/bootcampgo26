# 🧪 Guia de Testes com Postman - Supermarket API

## 🚀 Passo 1: Iniciar o Servidor

```bash
cd /Users/luizacosta/Documents/Estudos/AprendendoGo/bootcampgo26/go-web/supermarket-api
go run .
```

Você deve ver:
```
📚 Carregando produtos...
✅ 500 produtos carregados com sucesso!
🚀 Servidor rodando em http://localhost:8080
```

---

## 📋 Testes no Postman

### ✅ Teste 1: Verificar se o servidor está rodando

**Método:** `GET`  
**URL:** `http://localhost:8080/ping`

**Resposta esperada:**
```json
{
  "message": "pong"
}
```

---

### 📚 Teste 2: Listar todos os produtos

**Método:** `GET`  
**URL:** `http://localhost:8080/products`

**Resposta esperada:**
```json
[
  {
    "id": 1,
    "name": "Oil - Margarine",
    "quantity": 439,
    "code_value": "S82254D",
    "is_published": true,
    "expiration": "15/12/2021",
    "price": 71.42
  },
  ...
]
```

**Status:** `200 OK`

---

### 🔍 Teste 3: Buscar produto específico

**Método:** `GET`  
**URL:** `http://localhost:8080/products/1`

**Resposta esperada:**
```json
{
  "id": 1,
  "name": "Oil - Margarine",
  "quantity": 439,
  "code_value": "S82254D",
  "is_published": true,
  "expiration": "15/12/2021",
  "price": 71.42
}
```

**Status:** `200 OK`

**Teste de erro (produto não existe):**
- URL: `http://localhost:8080/products/99999`
- Status esperado: `404 Not Found`

---

### ➕ Teste 4: Criar novo produto

**Método:** `POST`  
**URL:** `http://localhost:8080/products`  
**Headers:** `Content-Type: application/json`

**Body (JSON):**
```json
{
  "name": "Café Pilão 500g",
  "quantity": 50,
  "code_value": "CAF001",
  "is_published": true,
  "expiration": "31/12/2026",
  "price": 18.90
}
```

**Resposta esperada:**
```json
{
  "id": 501,
  "name": "Café Pilão 500g",
  "quantity": 50,
  "code_value": "CAF001",
  "is_published": true,
  "expiration": "31/12/2026",
  "price": 18.90
}
```

**Status:** `201 Created`

**⚠️ IMPORTANTE:** O ID é gerado automaticamente! Não precisa enviar no JSON.

**🔍 Verificação:** Abra o arquivo `products.json` e veja o novo produto no final!

---

### ✏️ Teste 5: Atualizar produto existente

**Método:** `PUT`  
**URL:** `http://localhost:8080/products/501`  
**Headers:** `Content-Type: application/json`

**Body (JSON):**
```json
{
  "name": "Café Pilão 500g - PROMOÇÃO",
  "quantity": 45,
  "code_value": "CAF001",
  "is_published": true,
  "expiration": "31/12/2026",
  "price": 15.90
}
```

**Resposta esperada:**
```json
{
  "id": 501,
  "name": "Café Pilão 500g - PROMOÇÃO",
  "quantity": 45,
  "code_value": "CAF001",
  "is_published": true,
  "expiration": "31/12/2026",
  "price": 15.90
}
```

**Status:** `200 OK`

**🔍 Verificação:** Veja as mudanças no `products.json`!

---

### 🗑️ Teste 6: Deletar produto

**Método:** `DELETE`  
**URL:** `http://localhost:8080/products/501`

**Resposta esperada:** (sem body)

**Status:** `204 No Content`

**🔍 Verificação:** O produto 501 sumiu do `products.json`!

---

## 🎓 Conceitos Importantes Aplicados

### 1. **Persistência de Dados**

Toda mudança (POST, PUT, DELETE) chama `saveProducts()`, que:
- Pega os dados da memória (slice `products`)
- Transforma em JSON (`Marshal`)
- Salva no arquivo `products.json`

**Resultado:** Suas mudanças sobrevivem mesmo se reiniciar o servidor! 🎉

### 2. **RESTful API**

| Método   | Rota            | Ação              |
|----------|-----------------|-------------------|
| GET      | /products       | Listar todos      |
| GET      | /products/:id   | Buscar um         |
| POST     | /products       | Criar novo        |
| PUT      | /products/:id   | Atualizar         |
| DELETE   | /products/:id   | Deletar           |

**Padrão REST:** URLs representam recursos, métodos HTTP representam ações.

### 3. **Status Codes HTTP**

- `200 OK` - Sucesso
- `201 Created` - Criado com sucesso
- `204 No Content` - Sucesso sem retorno
- `400 Bad Request` - Erro na requisição
- `404 Not Found` - Não encontrado
- `500 Internal Server Error` - Erro no servidor

---

## 🐛 Troubleshooting

### Erro: "arquivo não encontrado"
- Verifique se `products.json` está na mesma pasta do `main.go`
- Rode `go run .` de dentro da pasta `supermarket-api`

### Erro: "address already in use"
- A porta 8080 já está em uso
- Pare o servidor anterior com `Ctrl+C`
- Ou mude a porta no código: `:8080` → `:8081`

### JSON inválido no Postman
- Confira aspas duplas (`"` não `'`)
- Confira vírgulas entre campos
- Use o formatter do Postman (botão "Beautify")

---

## 🎯 Exercícios Adicionais

1. **Filtro:** Adicione rota `GET /products?published=true` para listar só publicados
2. **Busca:** Adicione rota `GET /products?name=Coffee` para buscar por nome
3. **Validação:** Valide se o preço é maior que zero antes de criar produto
4. **Paginação:** Adicione parâmetros `?page=1&limit=10`

---

## 📖 Resumo Visual

```
POSTMAN (Você)          API (Go)              ARQUIVO
──────────────         ──────────            ──────────

POST /products  ──→  1. Unmarshal JSON    
   + JSON body         2. Adiciona ao slice  
                       3. saveProducts()  ──→  products.json
                       4. Marshal result        (atualizado)
                  ←──  JSON response

GET /products   ──→  1. Pega slice products
                       2. Marshal to JSON
                  ←──  JSON array

PUT /products/5 ──→  1. Encontra produto
   + JSON body         2. Unmarshal update
                       3. Substitui no slice
                       4. saveProducts()  ──→  products.json
                  ←──  JSON updated            (atualizado)
```

**Lembre-se:** JSON é só texto! Go precisa do `encoding/json` para transformar texto em structs e vice-versa. É como ter um tradutor universal! 🌍✨
