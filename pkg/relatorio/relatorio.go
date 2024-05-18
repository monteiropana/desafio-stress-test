package relatorio

import (
	testecarga "desafio-stress-test/pkg/teste-carga"
	"fmt"
)

func GenerateRepost(results []testecarga.Result) {
	var totalDuration float64
	statusCodeCount := make(map[int]int)

	for _, result := range results {
		totalDuration += result.ResponseTime
		statusCodeCount[result.StatusCode]++
	}

	avgDuration := totalDuration / float64(len(results))
	fmt.Printf("Tempo Médio de Requisição: %.2f segundos\n", avgDuration)
	fmt.Printf("Quantidade Total de Requisições: %d\n", len(results))
	fmt.Println("Distribuição dos Códigos de Status:")
	for code, count := range statusCodeCount {
		fmt.Printf("Status Code %d: %d vezes\n", code, count)
	}
}
