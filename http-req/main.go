package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// HTTP request
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	// Read body
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// Write to file
	err = os.WriteFile("resposta.txt", res, 0644)
	if err != nil {
		panic(err)
	}

	// Read from file
	ler, err := os.Open("resposta.txt")
	if err != nil {
		panic(err)
	}
	defer ler.Close()

	read := bufio.NewReader(ler)
	buffer := make([]byte, 10)

	for {
		n, err := read.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buffer[:n]))
	}
}
