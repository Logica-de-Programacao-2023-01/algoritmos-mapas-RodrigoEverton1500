package main

import (
	"fmt"
	"strings"
)

func contarLetrasPorPalavra(frase string) map[string]map[rune]int {

	palavras := strings.Fields(frase)

	resultado := make(map[string]map[rune]int)

	for _, palavra := range palavras {

		contagem := make(map[rune]int)
		for _, letra := range palavra {

			contagem[letra]++
		}

		resultado[palavra] = contagem
	}

	return resultado
}

func main() {
	frase := "Olá. Tudo bem?"

	resultado := contarLetrasPorPalavra(frase)

	for palavra, contagem := range resultado {
		fmt.Printf("Palavra: %s\n", palavra)
		for letra, quantidade := range contagem {
			fmt.Printf("  Letra: %c, Quantidade: %d\n", letra, quantidade)
		}
		fmt.Println()
	}
}
