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
	}

	return "Saldo insuficiente"
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.Sacar(-100))
	fmt.Println(contaDaSilvia)

	fmt.Println(contaDaSilvia.Sacar(250))
	fmt.Println(contaDaSilvia)

	fmt.Println(contaDaSilvia.Sacar(250))
	fmt.Println(contaDaSilvia)

	fmt.Println(contaDaSilvia.Sacar(800))
	fmt.Println(contaDaSilvia.saldo)
}
