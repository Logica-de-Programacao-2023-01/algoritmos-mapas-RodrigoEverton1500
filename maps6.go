package main

import (
	"fmt"
)

func somarContagens(contagens []map[string]int) map[string]int {
	soma := make(map[string]int)

	for _, contagem := range contagens {

		for palavra, frequencia := range contagem {

			soma[palavra] += frequencia
		}
	}

	return soma
}

func main() {

	contagens := []map[string]int{
		{"hello": 2, "world": 1},
		{"hello": 3, "go": 2, "world": 1},
		{"hello": 1, "oi": 5},
	}
	resultado := somarContagens(contagens)

	for palavra, frequencia := range resultado {
		fmt.Printf("Palavra: %s, Frequência: %d\n", palavra, frequencia)
	}
}
