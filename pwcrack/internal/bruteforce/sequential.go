package bruteforce

import (
	"fmt"
	"time"
)

// Result encapsula informações da execução sequencial.
type Result struct {
	Found      bool          // True se encontrou a senha
	Password   string        // A senha encontrada (se Found == true)
	Attempts   int64         // Quantidade de tentativas realizadas
	Elapsed    time.Duration // Tempo total de execução
	StartedAt  time.Time     // Timestamp de início
	FinishedAt time.Time     // Timestamp de término
}

// BruteForceSequential faz a busca de "00000000" até "99999999".
// Assim que encontrar o alvo, retorna imediatamente.
func BruteForceSequential(target string) Result {
	start := time.Now()
	var attempts int64

	for i := 0; i <= 99999999; i++ {
		attempts++
		candidate := fmt.Sprintf("%08d", i)
		if candidate == target {
			finish := time.Now()
			return Result{
				Found:      true,
				Password:   candidate,
				Attempts:   attempts,
				Elapsed:    finish.Sub(start),
				StartedAt:  start,
				FinishedAt: finish,
			}
		}
	}

	finish := time.Now()
	return Result{
		Found:      false,
		Password:   "",
		Attempts:   attempts,
		Elapsed:    finish.Sub(start),
		StartedAt:  start,
		FinishedAt: finish,
	}
}
