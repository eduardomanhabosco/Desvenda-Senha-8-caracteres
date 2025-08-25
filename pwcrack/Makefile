APP_NAME = pwcrack

# Pastas
BIN_DIR = bin
SEQ_DIR = ./cmd/sequential
PAR_DIR = ./cmd/parallel

# Go configs
GO = go
GOFLAGS = 

# Arquivos de saída
SEQ_BIN = $(BIN_DIR)/sequential.exe
PAR_BIN = $(BIN_DIR)/parallel.exe

# Variáveis que podem ser passadas pelo usuário
PASSWORD ?= 12345678
THREADS ?= 8

.PHONY: all build sequential parallel clean

# compila tudo
all: build

# compila os dois binários
build: $(SEQ_BIN) $(PAR_BIN)

# compila versão sequencial
$(SEQ_BIN):
	@if not exist $(BIN_DIR) mkdir $(BIN_DIR)
	$(GO) build $(GOFLAGS) -o $(SEQ_BIN) $(SEQ_DIR)

# compila versão paralela
$(PAR_BIN):
	@if not exist $(BIN_DIR) mkdir $(BIN_DIR)
	$(GO) build $(GOFLAGS) -o $(PAR_BIN) $(PAR_DIR)

# roda versão sequencial com senha personalizável
sequential: $(SEQ_BIN)
	$(SEQ_BIN) -password $(PASSWORD)

# roda versão paralela com senha e threads personalizáveis
parallel: $(PAR_BIN)
	$(PAR_BIN) -password $(PASSWORD) -threads $(THREADS)

# limpa os binários (Windows)
clean:
	if exist $(BIN_DIR) rmdir /s /q $(BIN_DIR)
