package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDoYuri := ContaCorrente{"Yuri", 323, 11222, 311}

	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	fmt.Println(contaDoYuri)
	fmt.Println(contaDoGuilherme)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(contaDaCris)
}
