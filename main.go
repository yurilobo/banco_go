package main

import (
	"banco_oo/clientes"
	"banco_oo/contas"
	"fmt"
)

func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.254.426.85",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}
	fmt.Println(contaDoBruno)

}
