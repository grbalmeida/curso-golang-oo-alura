package main

import (
	"fmt"

	"github.com/grbalmeida/curso-golang-oo-alura/banco/clientes"
	"github.com/grbalmeida/curso-golang-oo-alura/banco/contas"
)

func main() {
	clienteBruno := clientes.Titular{Nome: "Bruno", CPF: "123.123.123.12", Profissao: "Desenvolvedor Go"}
	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}
	fmt.Println(contaDoBruno)
}
