package relatorio

import (
	testecarga "desafio-stress-test/pkg/testecarga"
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
	fmt.Printf("Esse é Tempo Médio de Requisição: %.2f segundos\n", avgDuration)
	fmt.Printf("Essa é a Quantidade Total de Requisições: %d\n", len(results))
	fmt.Println("Distribuição dos Códigos de Status:")
	for code, count := range statusCodeCount {
		fmt.Printf("O numero total de Status Code %d: %d vezes\n", code, count)
	}
}
