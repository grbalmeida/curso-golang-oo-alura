package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.50}
	contaDoGuilherme2 := ContaCorrente{titular: "Guilherme", saldo: 125.50}
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	contaDaBruna2 := ContaCorrente{"Bruna", 222, 111222, 200}

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"

	fmt.Println(contaDoGuilherme == contaDoGuilherme2)
	fmt.Println(contaDaBruna == contaDaBruna2)
	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(*contaDaCris, *contaDaCris2)
}
