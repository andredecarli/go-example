# Exemplo de Projeto Go

Este projeto demonstra padrões comuns e boas práticas para estruturar aplicações em Go, incluindo interfaces, injeção de dependência, princípios de arquitetura limpa e padrões de design comuns.

## 🎯 Propósito

Este projeto serve como um **exemplo de aprendizado** para iniciar desenvolvimento em Go, demonstrando:

- **Design baseado em interfaces** - Definindo interfaces nos consumidores (inversão de dependência)
- **Arquitetura Limpa** - Separação de responsabilidades com camadas de domínio, aplicação e infraestrutura
- **Injeção de Dependência** - Baixo acoplamento através de composição de interfaces
- **Padrão Repository** - Abstraindo acesso a dados através de interfaces
- **Padrão Service Layer** - Encapsulamento da lógica de negócio
- **Padrão DTO** - Objetos de transferência de dados para entrada/saída
- **Tratamento de Erros** - Propagação adequada de erros e tipos de erro customizados

## 🏗️ Estrutura do Projeto

```
/
├── cmd/                  # Ponto de entrada da aplicação
│   └── main.go             # Função main com wiring de dependências
├── internal/             # Código privado da aplicação
│   ├── domain/             # Camada de domínio (entidades de negócio)
│   │   └── entities/         # Entidades principais de negócio
│   ├── infra/              # Camada de infraestrutura
│   │   └── db/               # Implementações de banco de dados
│   │       └── memory/         # Repository em memória
│   ├── application/        # Camada de aplicação (casos de uso)
│   │   ├── dto/              # Objetos de Transferência de Dados
│   │   └── services/         # Serviços de lógica de negócio
│   └── controllers/        # Camada de interface (controllers, handlers)
└── pkg/                  # Pacotes utilitários públicos
    └── util/               # Funções utilitárias
```

## 🔧 Padrões Go Demonstrados

### 1. Definição de Interface nos Consumidores
```go
// Em services/customer_service.go - Interface definida onde é usada
type CustomerRepository interface {
    Create(customer *entities.Customer) (*entities.Customer, error)
    FindAll() ([]*entities.Customer, error)
    // ... outros métodos
}
```

### 2. Injeção de Dependência
```go
// Padrão construtor com injeção de interface
func NewCustomerService(customerRepository CustomerRepository) *customerService {
    return &customerService{
        customerRepository: customerRepository,
    }
}
```

### 3. Composição ao invés de Herança
```go
// Customer composto por Entity
type Customer struct {
    Entity  // Struct embutido
    Name  string
    Email string
}
```

### 4. Receivers de Métodos
```go
// Receiver por valor para operações somente leitura
func (c Customer) String() string {
    return fmt.Sprintf("Customer{ID: %s, Name: %s, Email: %s}", ...)
}

// Receiver por ponteiro para modificações
func (r *customerRepository) Create(customer *entities.Customer) (*entities.Customer, error) {
    // Implementação
}
```

### 5. Tratamento de Erros
```go
// Tipos de erro customizados
var ErrCustomerAlreadyExists = errors.New("customer already exists")
var ErrCustomerNotFound = errors.New("customer not found")
```

## 🚀 Como Executar

### Pré-requisitos
- Go 1.24.4 ou superior
- Docker (opcional, para execução containerizada)

### Desenvolvimento Local

1. **Instale as dependências**
   ```bash
   go mod download
   ```

2. **Execute a aplicação**
   ```bash
   go run cmd/main.go
   ```

### Usando Docker

1. **Construa a imagem Docker**
   ```bash
   docker build -t app:latest .
   ```

2. **Execute o container**
   ```bash
   docker run --rm -it app:latest
   ```

### Usando Makefile

O projeto inclui um `Makefile` com comandos comuns:

```bash
make build    # Construir imagem Docker
make run      # Executar container Docker
make up       # Construir e executar
make clean    # Remover imagem Docker
```

## 📋 O que a Aplicação Faz

Quando você executa a aplicação, ela demonstra um fluxo completo de CRUD:

1. **Cria** um cliente
2. **Lista** todos os clientes
3. **Encontra** um cliente por ID
4. **Atualiza** as informações do cliente
5. **Remove** o cliente
6. **Verifica** a remoção tentando encontrar o cliente novamente

A aplicação usa um repositório em memória, então os dados são perdidos quando o programa termina.

## 🧪 Testando os Padrões

A aplicação demonstra vários conceitos importantes do Go:

- **Segregação de Interface**: Cada camada define apenas as interfaces que precisa
- **Inversão de Dependência**: Módulos de alto nível não dependem de módulos de baixo nível
- **Responsabilidade Única**: Cada struct tem um propósito claro
- **Princípio Aberto/Fechado**: Fácil de estender com novas implementações de repositório

## 🔄 Estendendo o Projeto

Para adicionar uma nova implementação de banco de dados (ex: PostgreSQL):

1. Crie um novo arquivo em `internal/infra/db/postgresql/customer.go`
2. Implemente a interface `CustomerRepository`
3. Atualize `cmd/main.go` para usar a nova implementação
4. Nenhuma mudança necessária nos serviços ou controllers!

---

**Nota**: Este é um projeto de aprendizado. Para uso em produção, considere adicionar logging adequado, gerenciamento de configuração, tratamento de erros e testes abrangentes.

