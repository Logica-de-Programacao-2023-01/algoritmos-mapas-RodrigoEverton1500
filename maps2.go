package main

import "fmt"

func mesclarMapas(mapa1, mapa2 map[string]int) map[string]int {
	resultado := make(map[string]int)
	for chave, valor := range mapa1 {
		resultado[chave] = valor
	}
	for chave, valor := range mapa2 {
		resultado[chave] = valor
	}

	return resultado
}

func main() {
	mapa1 := map[string]int{
		"chave1": 1,
		"chave2": 2,
	}

	mapa2 := map[string]int{
		"chave2": 3,
		"chave3": 4,
	}

	resultado := mesclarMapas(mapa1, mapa2)
	for chave, valor := range resultado {
		fmt.Printf("%s: %d\n", chave, valor)
	}
}
