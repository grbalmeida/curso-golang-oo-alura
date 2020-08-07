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

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	}

	return "Valor do depósito menor que zero", c.saldo
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	status, valor := contaDaSilvia.Depositar(500)
	fmt.Println("Status:", status, "Valor:", valor)

	status, valor = contaDaSilvia.Depositar(-300)
	fmt.Println("Status:", status, "Valor:", valor)

	status, valor = contaDaSilvia.Depositar(200)
	fmt.Println("Status:", status, "Valor:", valor)
}
