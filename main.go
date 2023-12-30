package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// var titular string = "Guilherme"
	// var numeroAgencia int = 589
	// var numeroConta int = 12346
	// var saldo float64 = 125.50
	contaDoYuri := ContaCorrente{"Yuri", 323, 11222, 311}

	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	fmt.Println(contaDoYuri)
	fmt.Println(contaDoGuilherme)

}
