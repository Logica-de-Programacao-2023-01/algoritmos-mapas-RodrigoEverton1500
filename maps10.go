package main

import "fmt"

func contarPares(slice []int) map[[2]int]int {
	frequencia := make(map[[2]int]int)

	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			frequencia[pair]++
		}
	}

	return frequencia
}

func main() {
	numeros := []int{1, 2, 3, 2, 4, 1, 5, 3, 2}

	resultado := contarPares(numeros)

	for par, frequencia := range resultado {
		fmt.Printf("%v: %d\n", par, frequencia)
	}
}
