package main

import (
	"fmt"
	"strings"
)

func contarPalavras(s string) map[string]int {
	palavras := strings.Fields(s) // Separa a string em palavras
	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	texto := "Olá, isso é um exemplo. Olá, mundo!"
	resultado := contarPalavras(texto)

	for palavra, ocorrencias := range resultado {
		fmt.Printf("%s: %d\n", palavra, ocorrencias)
	}
}
