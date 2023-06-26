package main

import (
	"fmt"
)

func calcularDivisaoConta(despesas map[string]float64) map[string]float64 {
	totalDespesas := 0.0
	numPessoas := 0

	for _, valor := range despesas {
		totalDespesas += valor
		numPessoas++
	}

	mediaDespesas := totalDespesas / float64(numPessoas)
	resultado := make(map[string]float64)

	for nome, valor := range despesas {
		valorDevido := valor - mediaDespesas
		resultado[nome] = valorDevido
	}

	return resultado
}

func main() {
	despesas := map[string]float64{
		"João":     120.50,
		"Maria":    80.25,
		"Pedro":    50.75,
		"Carolina": 95.0,
	}

	resultado := calcularDivisaoConta(despesas)

	for nome, valor := range resultado {
		fmt.Printf("%s deve receber/pagar R$ %.2f\n", nome, valor)
	}
}
