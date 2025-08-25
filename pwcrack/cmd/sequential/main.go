package main

import (
	"flag"
	"fmt"
	"os"

	"pwcrack/internal/bruteforce"
	"pwcrack/internal/validate"
)

func main() {
	// Flags de execução
	password := flag.String("password", "", "Senha de 8 dígitos a ser quebrada (ex: 12345678)")
	quiet := flag.Bool("quiet", false, "Silencia o progresso (saída somente do resultado)")
	flag.Parse()

	// Checagem básica das flags
	if *password == "" {
		fmt.Fprintln(os.Stderr, "erro: use -password para informar a senha alvo (ex: -password 12345678)")
		os.Exit(2)
	}

	// Validação de formato (8 dígitos numéricos)
	if err := validate.ValidateNumeric8(*password); err != nil {
		fmt.Fprintln(os.Stderr, "erro de validação:", err)
		os.Exit(2)
	}

	if !*quiet {
		fmt.Println(">>> Modo: Sequencial")
		fmt.Println(">>> Alvo:", *password)
		fmt.Println(">>> Iniciando brute force...")
	}

	// Execução sequencial
	result := bruteforce.BruteForceSequential(*password)

	// Impressão de estatísticas
	if result.Found {
		fmt.Println("--------------------------------------------------")
		fmt.Println("Senha encontrada!")
		fmt.Println("Alvo:          ", *password)
		fmt.Println("Encontrada:    ", result.Password)
		fmt.Println("Tentativas:    ", result.Attempts)
		fmt.Println("Início:        ", result.StartedAt.Format("2006-01-02 15:04:05.000"))
		fmt.Println("Término:       ", result.FinishedAt.Format("2006-01-02 15:04:05.000"))
		fmt.Println("Tempo total:   ", result.Elapsed)
		// Taxa aproximada (tentativas por segundo)
		secs := result.Elapsed.Seconds()
		if secs > 0 {
			fmt.Printf("Taxa média:    %.0f tentativas/seg\n", float64(result.Attempts)/secs)
		}
		fmt.Println("--------------------------------------------------")
	} else {
		fmt.Println("--------------------------------------------------")
		fmt.Println("Senha NÃO encontrada (verifique o formato/alvo).")
		fmt.Println("Tentativas:    ", result.Attempts)
		fmt.Println("Tempo total:   ", result.Elapsed)
		fmt.Println("--------------------------------------------------")
	}
}
