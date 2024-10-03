package contas

import "github.com/deadman360/cursoAlura2/clientes"

type ContaPoupança struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (conta *ContaPoupança) Saque(valor float64) bool {
	if conta.saldo >= valor && valor > 0 {
		conta.saldo -= valor
		return true
	} else {
		return false
	}
}
func (conta *ContaPoupança) Deposito(valor float64) (string, float64) {
	if valor > 0 {
		conta.saldo += valor
		return "Deposito realizado com sucesso!", conta.saldo
	} else {
		return "Valor inválido!", valor
	}
}
func (conta *ContaPoupança) Extrato() float64 {
	return conta.saldo
}
