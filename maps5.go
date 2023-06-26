package main

import (
	"fmt"
)

func contarCaracteres(str string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range str {

		frequencia[char]++
	}

	return frequencia
}

func main() {

	str := "abracadabra"

	resultado := contarCaracteres(str)

	for char, count := range resultado {
		fmt.Printf("Caractere: %c, Frequência: %d\n", char, count)
	}
}
