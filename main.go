package main

import (
	"banco_oo/clientes"
	"banco_oo/contas"
	"fmt"
)

func main() {
	clienteBruno := clientes.Titular{"Yuri", "123.254.426.85", "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 12645, 1000}
	fmt.Println(contaDoBruno)

}
