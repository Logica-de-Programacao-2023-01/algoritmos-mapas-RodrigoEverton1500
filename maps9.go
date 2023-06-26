package main

import (
	"fmt"
)

func gerarSequenciaFibonacci(n int) map[int]int {
	sequencia := make(map[int]int)

	sequencia[0] = 0
	sequencia[1] = 1

	for i := 2; ; i++ {

		numero := sequencia[i-1] + sequencia[i-2]

		if numero > n {
			break
		}

		sequencia[i] = numero
	}

	return sequencia
}

func main() {
	n := 100

	sequencia := gerarSequenciaFibonacci(n)

	for indice, valor := range sequencia {
		fmt.Printf("Fibonacci(%d) = %d\n", indice, valor)
	}
}
