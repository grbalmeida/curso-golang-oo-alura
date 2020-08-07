package main

import (
	"fmt"

	"github.com/grbalmeida/curso-golang-oo-alura/banco/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	fmt.Println(status, contaDaSilvia, contaDoGustavo)

	status = contaDoGustavo.Transferir(50, &contaDaSilvia)
	fmt.Println(status, contaDaSilvia, contaDoGustavo)

	status = contaDoGustavo.Transferir(250, &contaDaSilvia)
	fmt.Println(status, contaDaSilvia, contaDoGustavo)
}
