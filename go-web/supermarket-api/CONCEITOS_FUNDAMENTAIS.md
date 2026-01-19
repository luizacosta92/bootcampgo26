# 🎓 Conceitos Fundamentais - JSON em Go

## 🌟 A Grande Ideia

**JSON é a língua universal da web.** Assim como o inglês conecta pessoas de diferentes países, JSON conecta diferentes sistemas e linguagens de programação.

---

## 🔄 O Ciclo Marshal/Unmarshal

### Visualização:

```
┌─────────────────────────────────────────────────┐
│              MUNDO GO (Estruturado)             │
│                                                 │
│   type Product struct {                         │
│       ID    int                                 │
│       Name  string                              │
│       Price float64                             │
│   }                                             │
│                                                 │
│   product := Product{                           │
│       ID: 1,                                    │
│       Name: "Café",                             │
│       Price: 15.90                              │
│   }                                             │
└─────────────────────────────────────────────────┘
                      │
                      │ json.Marshal() ↓
                      │ "Empacotar"
                      ↓
            ┌──────────────────┐
            │   {"id":1,       │
            │   "name":"Café", │
            │   "price":15.90} │
            └──────────────────┘
                      │
                      │ json.Unmarshal() ↑
                      │ "Desempacotar"
                      │
┌─────────────────────────────────────────────────┐
│           MUNDO JSON (Texto/String)             │
│                                                 │
│   - Fácil para humanos lerem                    │
│   - Enviado via HTTP                            │
│   - Salvo em arquivos                           │
│   - Usado por JavaScript, Python, etc           │
└─────────────────────────────────────────────────┘
```

---

## 🏷️ Tags JSON: Os Tradutores

### Por que precisamos delas?

```go
// ❌ SEM TAGS - Não funciona!
type Product struct {
    id    int     // Privado (minúsculo) = JSON não consegue acessar
    name  string  // JSON busca "name", Go tem "name" mas privado
}

// ✅ COM TAGS - Funciona perfeitamente!
type Product struct {
    ID    int    `json:"id"`      // Go: ID   (público) ↔ JSON: id
    Name  string `json:"name"`    // Go: Name (público) ↔ JSON: name
    Price float64 `json:"price"`  // Go: Price (público) ↔ JSON: price
}
```

### Anatomia de uma Tag JSON:

```go
Price float64 `json:"price"`
│     │        └─────────┬─────────┘
│     │                  └─ Tag: conecta Go ↔ JSON
│     └─ Tipo Go
└─ Nome Go (DEVE ser público - primeira letra maiúscula)
```

### Analogia 🎭

Imagine um intérprete simultâneo em uma conferência:
- **Palestrante Go** diz: "Price" (em Golês)
- **Tag JSON** traduz: "price" (em JSONês)
- **Plateia JavaScript/Postman** entende perfeitamente!

---

## 🧠 Memória vs Disco

### Por que carregamos tudo na memória?

```
╔══════════════════════╗     ╔══════════════════════╗
║   DISCO (HD/SSD)     ║     ║    MEMÓRIA (RAM)     ║
║   products.json      ║     ║   var products []    ║
╚══════════════════════╝     ╚══════════════════════╝
         🐌                            🚀
    Lento (ms)                   Rápido (ns)
    Permanente                   Temporário
    500KB                        Carregado na RAM
```

**Estratégia:**
1. **Startup:** Carrega JSON → Memória (1x, lento)
2. **Requisições:** Lê/modifica na Memória (1000x, super rápido)
3. **Mudanças:** Salva Memória → JSON (quando necessário)

**Analogia 📚:**
- **Disco** = Biblioteca (longe, precisa buscar)
- **Memória** = Mesa de trabalho (perto, acesso instantâneo)

Você não vai à biblioteca buscar o mesmo livro 100 vezes. Você pega uma vez, trabalha na mesa, e devolve quando terminar!

---

## 🎯 Ponteiros: O Endereço Real

### Por que `&` é importante?

```go
// Sem ponteiro (CÓPIA)
func modificar(produto Product) {
    produto.Name = "Novo Nome"
    // Modifica a CÓPIA, não o original!
}

// Com ponteiro (REFERÊNCIA)
func modificar(produto *Product) {
    produto.Name = "Novo Nome"
    // Modifica o ORIGINAL! ✅
}
```

### Analogia 🏠

**Sem ponteiro (`Product`):**
- Você dá uma **foto** da sua casa para o pintor
- Ele pinta a foto de azul
- Sua casa continua branca! 🏠 (nada mudou)

**Com ponteiro (`*Product` ou `&Product`):**
- Você dá o **endereço** da sua casa para o pintor
- Ele vai lá e pinta de azul
- Sua casa está azul! 🏠💙 (mudou de verdade!)

### Uso prático:

```go
// Unmarshal PRECISA de ponteiro para modificar a variável
var product Product
json.Unmarshal(data, &product)  // ← & dá o "endereço" para modificar
```

---

## 🌐 HTTP Methods: Os Verbos da Web

```
┌─────────┬─────────────┬──────────────────────────┐
│ Método  │ Ação        │ Analogia                 │
├─────────┼─────────────┼──────────────────────────┤
│ GET     │ Buscar      │ "Me MOSTRA o produto 5"  │
│ POST    │ Criar       │ "ADICIONA este produto"  │
│ PUT     │ Atualizar   │ "TROCA o produto 5"      │
│ DELETE  │ Deletar     │ "REMOVE o produto 5"     │
└─────────┴─────────────┴──────────────────────────┘
```

### Características importantes:

**GET - O mais usado**
- Não tem body (dados na URL)
- Seguro (só lê, não muda nada)
- Pode ser cacheado

**POST - Criação**
- Tem body (JSON com dados)
- Cria recurso novo
- Retorna 201 Created

**PUT - Substituição completa**
- Tem body (JSON completo)
- Substitui o recurso inteiro
- Retorna 200 OK

**DELETE - Remoção**
- Não tem body
- Remove o recurso
- Retorna 204 No Content

---

## 📦 Fluxo Completo de uma Requisição POST

```
1. POSTMAN                          2. GO API
   ↓                                   ↓
   POST /products                      func createProductHandler(...)
   Body: {                             ↓
     "name": "Café",                   body, _ := io.ReadAll(r.Body)
     "price": 15.90                    ↓
   }                                   var newProduct Product
   │                                   ↓
   │                                   json.Unmarshal(body, &newProduct)
   │                                   ↓
   │                           ┌──────────────────────┐
   │                           │  Product struct      │
   │                           │  Name: "Café"        │
   │                           │  Price: 15.90        │
   │                           └──────────────────────┘
   │                                   ↓
   │                           newProduct.ID = 501
   │                                   ↓
   │                           products = append(products, newProduct)
   │                                   ↓
   │                           saveProducts()
   │                                   ↓
   │                           json.MarshalIndent(products, ...)
   │                                   ↓
   │                           os.WriteFile("products.json", ...)
   │                                   ↓
   ├───────────────────────────  w.WriteHeader(201)
   │                                   ↓
   ↓                                json.Marshal(newProduct)
   Response 201 Created                ↓
   Body: {                          w.Write(data)
     "id": 501,
     "name": "Café",
     "price": 15.90
   }

3. products.json (ATUALIZADO!)
   [...499 produtos anteriores..., 
    {"id":501,"name":"Café","price":15.90}]
```

---

## 🎨 Padrões de Design

### 1. Repository Pattern (Implícito)

```go
// Funções que gerenciam dados = "Repository"
loadProducts()      // READ do arquivo
saveProducts()      // WRITE no arquivo
findProductByID()   // QUERY na memória
```

**Benefício:** Se mudar de arquivo JSON para banco de dados, só muda essas 3 funções!

### 2. Handler Pattern

```go
// Cada rota tem seu handler
http.HandleFunc("/products", productsRouter)
```

**Benefício:** Organização clara, fácil adicionar rotas novas.

### 3. Error Handling

```go
if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"error": "mensagem clara"}`))
    return  // ← SEMPRE retorne após erro!
}
```

**Regra de Ouro:** Trate erros imediatamente e retorne!

---

## 💡 Dicas Práticas

### 1. Debugging JSON

```go
// Ver JSON formatado no terminal
data, _ := json.MarshalIndent(product, "", "  ")
fmt.Println(string(data))
```

### 2. Validação de Dados

```go
if newProduct.Price <= 0 {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(`{"error": "Preço deve ser positivo"}`))
    return
}
```

### 3. Headers sempre!

```go
w.Header().Set("Content-Type", "application/json")
// ↑ Diz ao cliente: "Estou enviando JSON!"
```

---

## 🚀 Próximos Passos

Agora que você entende os fundamentos, pode explorar:

1. **Banco de Dados:** MySQL, PostgreSQL (ao invés de JSON)
2. **Frameworks:** Gin, Echo (ao invés de http padrão)
3. **ORM:** GORM (mapeia structs → tabelas)
4. **Validação:** go-validator (valida structs automaticamente)
5. **Testes:** testing package (testar handlers)
6. **Middleware:** Autenticação, logging, CORS
7. **Docker:** Containerizar a aplicação

---

## 📚 Glossário Rápido

- **Marshal:** Go struct → JSON (empacotar)
- **Unmarshal:** JSON → Go struct (desempacotar)
- **Handler:** Função que responde a uma requisição HTTP
- **Router:** Decide qual handler chamar
- **Body:** Corpo da requisição/resposta HTTP
- **Header:** Metadados da requisição/resposta
- **Status Code:** Número que indica resultado (200, 404, etc)
- **Endpoint:** URL da API (ex: `/products`)
- **REST:** Estilo de arquitetura para APIs
- **CRUD:** Create, Read, Update, Delete

---

## 🎯 Resumo em uma frase

> **JSON é texto que representa dados, Go precisa de `encoding/json` + structs com tags para converter esse texto em objetos manipuláveis e vice-versa!**

Pronto! Agora você tem a base sólida para construir APIs RESTful em Go! 🚀
