package main

import (
	"encoding/json"
	"fmt"
)

type Conta struct {
	Nome  string
	Saldo float64
}

func main() {
	conta := Conta{
		Nome:  "varallo",
		Saldo: 15.78,
	}

	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

}
