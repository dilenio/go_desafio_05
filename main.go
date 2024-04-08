package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	url         string
	requests    int
	concurrency int
)

func init() {
	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&requests, "requests", 100, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 10, "Número de chamadas simultâneas")
}

func main() {
	flag.Parse()

	if url == "" {
		fmt.Println("URL é obrigatório")
		return
	}

	fmt.Printf("Testando %s com %d requests e concorrência de %d...\n", url, requests, concurrency)

	start := time.Now()
	statusCodes := make(map[int]int)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	semaphore := make(chan struct{}, concurrency)

	for i := 0; i < requests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			mutex.Lock()
			statusCodes[resp.StatusCode]++
			mutex.Unlock()

			<-semaphore
		}()
	}

	wg.Wait()
	totalTime := time.Since(start)

	fmt.Println("Relatório de Teste de Carga")
	fmt.Printf("Tempo Total: %v\n", totalTime)
	fmt.Printf("Requests Totais: %d\n", requests)

	mutex.Lock()
	for statusCode, count := range statusCodes {
		fmt.Printf("Status %d: %d vezes\n", statusCode, count)
	}
	mutex.Unlock()
}
