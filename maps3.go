package main

import "fmt"

func somarValoresMapa(m mapaInt) int {
	soma := 0

	for _, valor := range m {
		soma += valor
	}

	return soma
}

type mapaInt map[string]int

func main() {
	mapa := mapaInt{
		"chave1": 10,
		"chave2": 20,
		"chave3": 30,
	}
	resultado := somarValoresMapa(mapa)

	fmt.Println("Soma dos valores:", resultado)
}
