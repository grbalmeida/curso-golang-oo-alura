package main

import (
	"fmt"

	"github.com/grbalmeida/curso-golang-oo-alura/banco/contas"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(55)
	contaDoDenis.Sacar(5000)

	fmt.Println(contaDoDenis.ObterSaldo())
}
