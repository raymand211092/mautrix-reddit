.PHONY: all build clean test run install deps fmt lint docker help

# Variables
BINARY_NAME=mautrix-reddit
GO=go
GOFLAGS=-v
LDFLAGS=-ldflags="-s -w"

# Colores para output
COLOR_RESET=\033[0m
COLOR_BOLD=\033[1m
COLOR_GREEN=\033[32m
COLOR_YELLOW=\033[33m

all: build ## Build el proyecto

help: ## Mostrar ayuda
	@echo "$(COLOR_BOLD)Comandos disponibles:$(COLOR_RESET)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(COLOR_GREEN)%-15s$(COLOR_RESET) %s\n", $$1, $$2}'

deps: ## Descargar dependencias
	@echo "$(COLOR_YELLOW)Descargando dependencias...$(COLOR_RESET)"
	$(GO) mod download
	$(GO) mod verify

build: deps ## Compilar el proyecto
	@echo "$(COLOR_YELLOW)Compilando $(BINARY_NAME)...$(COLOR_RESET)"
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BINARY_NAME) ./cmd/mautrix-reddit
	@echo "$(COLOR_GREEN)✓ Compilación exitosa!$(COLOR_RESET)"

clean: ## Limpiar archivos generados
	@echo "$(COLOR_YELLOW)Limpiando...$(COLOR_RESET)"
	$(GO) clean
	rm -f $(BINARY_NAME)
	rm -f *.db *.db-shm *.db-wal
	@echo "$(COLOR_GREEN)✓ Limpieza completa!$(COLOR_RESET)"

test: ## Ejecutar tests
	@echo "$(COLOR_YELLOW)Ejecutando tests...$(COLOR_RESET)"
	$(GO) test -v -race -coverprofile=coverage.txt ./...

test-coverage: test ## Ver cobertura de tests
	$(GO) tool cover -html=coverage.txt

run: build ## Compilar y ejecutar
	@echo "$(COLOR_YELLOW)Ejecutando $(BINARY_NAME)...$(COLOR_RESET)"
	./$(BINARY_NAME) -c config.yaml

install: build ## Instalar binario
	@echo "$(COLOR_YELLOW)Instalando $(BINARY_NAME)...$(COLOR_RESET)"
	$(GO) install ./cmd/mautrix-reddit
	@echo "$(COLOR_GREEN)✓ Instalado en $(GOPATH)/bin/$(BINARY_NAME)$(COLOR_RESET)"

fmt: ## Formatear código
	@echo "$(COLOR_YELLOW)Formateando código...$(COLOR_RESET)"
	$(GO) fmt ./...
	@echo "$(COLOR_GREEN)✓ Código formateado!$(COLOR_RESET)"

lint: ## Ejecutar linter
	@echo "$(COLOR_YELLOW)Ejecutando linter...$(COLOR_RESET)"
	golangci-lint run ./...

config: ## Generar configuración de ejemplo
	@echo "$(COLOR_YELLOW)Generando config.yaml...$(COLOR_RESET)"
	./$(BINARY_NAME) -e > config.yaml
	@echo "$(COLOR_GREEN)✓ Configuración generada!$(COLOR_RESET)"

register: ## Generar archivo de registro
	@echo "$(COLOR_YELLOW)Generando registration.yaml...$(COLOR_RESET)"
	./$(BINARY_NAME) -g -c config.yaml -r registration.yaml
	@echo "$(COLOR_GREEN)✓ Registro generado!$(COLOR_RESET)"

docker-build: ## Construir imagen Docker
	@echo "$(COLOR_YELLOW)Construyendo imagen Docker...$(COLOR_RESET)"
	docker build -t mautrix-reddit:latest .
	@echo "$(COLOR_GREEN)✓ Imagen Docker construida!$(COLOR_RESET)"

docker-run: docker-build ## Ejecutar en Docker
	docker run --rm -v $(PWD)/data:/data mautrix-reddit:latest

docker-compose-up: ## Iniciar con docker-compose
	docker-compose up -d

docker-compose-down: ## Detener docker-compose
	docker-compose down

docker-compose-logs: ## Ver logs de docker-compose
	docker-compose logs -f

dev: ## Ejecutar en modo desarrollo con auto-reload
	@echo "$(COLOR_YELLOW)Ejecutando en modo desarrollo...$(COLOR_RESET)"
	@command -v air > /dev/null || (echo "Instalando air..." && go install github.com/cosmtrek/air@latest)
	air
