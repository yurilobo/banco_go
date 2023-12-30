package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", saldo

	} else {
		return "Valor do deposito menor que zero", saldo
	}
}
func main() {

	contaDoYuri := ContaCorrente{"Yuri", 323, 11222, 311}

	fmt.Println(contaDoYuri.saldo)

	fmt.Println(contaDoYuri.Sacar(5))
	fmt.Println(contaDoYuri.saldo)

	fmt.Println(contaDoYuri.Depositar(100))

	fmt.Println(contaDoYuri.saldo)
}
