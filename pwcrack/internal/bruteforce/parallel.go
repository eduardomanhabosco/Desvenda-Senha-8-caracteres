package bruteforce

import (
	"context"
	"strconv"
	"sync"
	"time"
)

// BruteForceParallel divide o espaço de busca entre N threads,
// compara números diretamente e só converte para string ao encontrar a senha.
func BruteForceParallel(target string, threads int) Result {
	start := time.Now()

	
	targetInt, _ := strconv.Atoi(target) // converte senha p/ int

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	resultChan := make(chan Result, threads)

	chunk := 100_000_000 / threads
	wg.Add(threads)

	for t := 0; t < threads; t++ {
		begin := t * chunk
		end := begin + chunk
		if t == threads-1 {
			end = 100_000_000
		}

		go func(startRange, endRange int) {
			defer wg.Done()
			localAttempts := int64(0) //conta tentativas locais

			for i := startRange; i < endRange; i++ {
				select {
				case <-ctx.Done():
					return
				default:
				}

				localAttempts++

				if i == targetInt {
					finish := time.Now()
					resultChan <- Result{
						Found:      true,
						Password:   strconv.Itoa(i),
						Attempts:   localAttempts,
						Elapsed:    finish.Sub(start),
						StartedAt:  start,
						FinishedAt: finish,
					}
					cancel()
					return
				}
			}

			resultChan <- Result{
				Found:    false,
				Attempts: localAttempts,
			}
		}(begin, end)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var totalAttempts int64
	finalResult := Result{
		Found:     false,
		Password:  "",
		StartedAt: start,
	}

	for res := range resultChan {
		totalAttempts += res.Attempts
		if res.Found {
			finalResult = res
			finalResult.Attempts = totalAttempts
		}
	}

	if !finalResult.Found {
		finalResult.Attempts = totalAttempts
		finalResult.Elapsed = time.Since(start)
		finalResult.FinishedAt = time.Now()
	} else {
		finalResult.Elapsed = finalResult.FinishedAt.Sub(finalResult.StartedAt)
	}

	return finalResult
}
