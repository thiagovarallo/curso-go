package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Escrever no arquivo
	tamanho, err := f.Write([]byte("escrevendo.............................."))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Tamanho em bytes: %d\n", tamanho)

	// Lendo o arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Lendo o arquivo de pouco a pouco
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	ler := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		l, err := ler.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:l]))
	}

	// Removendo o arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
