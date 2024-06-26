package main

import (
	"desafio-stress-test/pkg/relatorio"
	"desafio-stress-test/pkg/testecarga"
	"flag"
	"fmt"
)

var (
	url         string
	totalReq    int
	concurrency int
)

func init() {
	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&totalReq, "requests", 100, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 10, "Número de chamadas simultâneas")
}

func main() {

	flag.Parse()

	if url == "" {
		fmt.Println("URL é um parâmetro obrigatório")
		return
	}

	results := testecarga.TesteCarga(url, totalReq, concurrency)
	relatorio.GenerateRepost(results)
}
