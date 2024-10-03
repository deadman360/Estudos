package main

import (
	"fmt"

	"github.com/deadman360/cursoAlura2/clientes"
	"github.com/deadman360/cursoAlura2/contas"
)

func PagaBoleto(conta contaBancaria, valor float64) {
	conta.Saque(valor)
}

type contaBancaria interface {
	Saque(valor float64) bool
}

func main() {
	clienteBea := clientes.Titular{
		Nome:      "Bea",
		CPF:       "000.000.000-00",
		Profissao: "Musicista",
	}
	contaBea := contas.ContaCorrente{
		Titular:       clienteBea,
		NumeroAgencia: 549,
		NumeroConta:   0001,
	}

	clienteVictor := clientes.Titular{
		Nome:      "Victor",
		CPF:       "000.000.000-00",
		Profissao: "Dev",
	}
	contaVictor := contas.ContaCorrente{
		Titular:       clienteVictor,
		NumeroAgencia: 548,
		NumeroConta:   0002,
	}
	clienteGabs := clientes.Titular{
		Nome:      "Gabriel",
		CPF:       "000.000.000-00",
		Profissao: "Gestor de eventos",
	}
	contaGabs := contas.ContaPoupança{
		Titular:       clienteGabs,
		NumeroAgencia: 598,
		NumeroConta:   0003,
		Operacao:      1,
	}
	contaBea.Deposito(4000)
	contaVictor.Deposito(347)
	contaGabs.Deposito(20000)

	fmt.Println(contaBea.Extrato(), contaVictor.Extrato())

	PagaBoleto(&contaBea, 500)

	fmt.Println(contaBea.Extrato(), contaVictor.Extrato())

}
