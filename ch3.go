package main

import (
	"fmt"
)

func main() {
	selamatMalam := []string{"s", "e", "l", "a", "m", "a", "t", " ", "m", "a", "l", "a", "m"}

	for i := 0; i < len(selamatMalam); i++ {
		fmt.Println(string(selamatMalam[i]))
	}

	hasilKata := handleKata(selamatMalam)
	fmt.Println(hasilKata)
}

func handleKata(selamatMalam []string) map[string]int {
	var result map[string]int = map[string]int{}

	for _, eachKata := range selamatMalam {
		result[eachKata]++
	}

	return result
}
