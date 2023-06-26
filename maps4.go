package main

import (
	"fmt"
	"sort"
	"strings"
)

func obterAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {

		sorted := sortString(palavra)

		anagramas[sorted] = append(anagramas[sorted], palavra)
	}

	return anagramas
}
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func main() {

	palavras := []string{"amor", "roma", "carro", "corra", "arco"}

	resultado := obterAnagramas(palavras)

	for _, grupo := range resultado {
		fmt.Println(grupo)
	}
}
