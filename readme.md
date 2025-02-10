# 🚀 Load Tester CLI - Teste de Carga com Golang

## 📌 Introdução

O **Load Tester CLI** é uma ferramenta em Go para realizar testes de carga em serviços web. Ele permite enviar um número configurável de requisições HTTP concorrentes, coletando estatísticas detalhadas sobre os tempos de resposta e status HTTP.

## 🛠 Instalação

### **Requisitos**

- Go **1.21+**
- Docker (opcional, para execução em container)

### **Clonar o Repositório**

```sh
git clone https://github.com/pietronirod/pg_desafio2.git
cd load-tester
```

### **Instalar Dependências**

```sh
go mod tidy
```

### **Compilar a Aplicação**

```sh
go build -o load-tester main.go
```

## 🚀 Uso

### **Executar o Teste de Carga via CLI**

```sh
./load-tester test --url=http://example.com --requests=100 --concurrency=5
```

### **Parâmetros**

| Parâmetro     | Descrição                                         | Padrão |
|--------------|-------------------------------------------------|--------|
| `--url`       | URL do serviço a ser testado                      | Obrigatório |
| `--requests`  | Número total de requisições a serem enviadas     | `10` |
| `--concurrency` | Número de chamadas simultâneas                  | `2` |
| `--output`    | Formato do relatório (`console`, `json`, `csv`) | `console` |

### **Exportar Relatórios**

#### JSON

```sh
./load-tester test --url=http://example.com --requests=100 --concurrency=5 --output=json
```

📂 Gera o arquivo `report.json`

#### CSV

```sh
./load-tester test --url=http://example.com --requests=100 --concurrency=5 --output=csv
```

📂 Gera o arquivo `report.csv`

## 📦 Executar via Docker

### **Construir a Imagem**

```sh
docker build -t load-tester .
```

### **Executar o Teste de Carga no Docker**

```sh
docker run --rm load-tester test --url=http://example.com --requests=100 --concurrency=5
```

### **Executar com Docker Compose (Opcional)**

Rodar o teste com:

```sh
docker-compose up --build
```

## 📊 Exemplo de Saída

```sh
Iniciando teste de carga...
URL: http://example.com
Total de Requests: 100
Concorrência: 5

--- Relatório Final ---
Tempo total: 2.5s
Total de requests: 100
Tempo médio de resposta: 200ms
Timeouts: 0

Distribuição de Status HTTP:
  - HTTP 200: 100 vezes

Teste concluído!
```

## 📜 Licença

Este projeto é distribuído sob a licença MIT. Sinta-se à vontade para usar e contribuir!

---
🚀 **Desenvolvido com Go para testes de carga eficientes e rápidos!**
