package main

import (
	"fmt"

	"github.com/grbalmeida/curso-golang-oo-alura/banco/contas"
)

func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}
