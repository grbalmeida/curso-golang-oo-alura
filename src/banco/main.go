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

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	}

	return false
}

func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	fmt.Println(status, contaDaSilvia, contaDoGustavo)

	status = contaDoGustavo.Transferir(50, &contaDaSilvia)
	fmt.Println(status, contaDaSilvia, contaDoGustavo)

	status = contaDoGustavo.Transferir(250, &contaDaSilvia)
	fmt.Println(status, contaDaSilvia, contaDoGustavo)
}
