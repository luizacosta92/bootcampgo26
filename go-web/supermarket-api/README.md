# 🛒 Supermarket API - Projeto de Estudos

API RESTful para gerenciamento de produtos de supermercado, construída em Go.

## 📁 Estrutura do Projeto

```
supermarket-api/
├── main.go                      # Servidor HTTP e handlers
├── productmodel.go              # Estrutura do Product
├── products.json                # Dados persistentes (500 produtos)
├── GUIA_TESTE_POSTMAN.md       # Tutorial de testes
└── CONCEITOS_FUNDAMENTAIS.md   # Explicações didáticas
```

## 🚀 Como Rodar

```bash
# 1. Navegue até a pasta
cd go-web/supermarket-api

# 2. Execute o servidor
go run .

# 3. Veja a mensagem de sucesso
# 📚 Carregando produtos...
# ✅ 500 produtos carregados com sucesso!
# 🚀 Servidor rodando em http://localhost:8080
```

## 📍 Endpoints Disponíveis

| Método | Rota              | Descrição                |
|--------|-------------------|--------------------------|
| GET    | `/ping`           | Verificar se está ativo  |
| GET    | `/products`       | Listar todos os produtos |
| GET    | `/products/:id`   | Buscar produto por ID    |
| POST   | `/products`       | Criar novo produto       |
| PUT    | `/products/:id`   | Atualizar produto        |
| DELETE | `/products/:id`   | Deletar produto          |

## 🧪 Testar com Postman

Veja o arquivo `GUIA_TESTE_POSTMAN.md` para exemplos detalhados.

### Exemplo Rápido - Buscar Produto

```
GET http://localhost:8080/products/1
```

**Resposta:**
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

### Exemplo - Criar Produto

```
POST http://localhost:8080/products
Content-Type: application/json

{
  "name": "Café Pilão 500g",
  "quantity": 50,
  "code_value": "CAF001",
  "is_published": true,
  "expiration": "31/12/2026",
  "price": 18.90
}
```

## 🎓 Conceitos Aprendidos

1. **JSON Marshal/Unmarshal** - Converter entre Go structs e JSON
2. **HTTP Handlers** - Processar requisições HTTP
3. **REST Pattern** - Estrutura de API RESTful
4. **File I/O** - Ler e escrever arquivos
5. **Routing** - Direcionar requisições para handlers corretos
6. **Error Handling** - Tratar erros adequadamente

Leia `CONCEITOS_FUNDAMENTAIS.md` para explicações detalhadas com analogias!

## 📦 Dependências

Apenas biblioteca padrão do Go:
- `encoding/json` - Trabalhar com JSON
- `net/http` - Servidor HTTP
- `io` - Leitura de dados
- `os` - Operações com arquivos

## 🔧 Características

✅ CRUD completo (Create, Read, Update, Delete)  
✅ Persistência em arquivo JSON  
✅ Validação de dados  
✅ Códigos HTTP apropriados  
✅ Comentários explicativos no código  
✅ Error handling robusto  

## 📝 Notas Importantes

- **Arquivo JSON:** Todas as alterações são salvas em `products.json`
- **IDs:** Gerados automaticamente ao criar produtos
- **Memória:** Produtos carregados na RAM para performance
- **Sincronização:** `saveProducts()` garante persistência

## 🎯 Próximas Melhorias

- [ ] Adicionar filtros de busca
- [ ] Implementar paginação
- [ ] Adicionar validações mais robustas
- [ ] Criar testes unitários
- [ ] Adicionar middleware de logging
- [ ] Implementar autenticação

## 📚 Recursos Adicionais

- [Documentação oficial Go](https://go.dev/doc/)
- [encoding/json package](https://pkg.go.dev/encoding/json)
- [net/http package](https://pkg.go.dev/net/http)

---

**Feito com 💙 para aprendizado de Go**
