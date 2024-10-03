package contas

import "github.com/deadman360/cursoAlura2/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Saque(valor float64) bool {
	if conta.saldo >= valor && valor > 0 {
		conta.saldo -= valor
		return true
	} else {
		return false
	}
}
func (conta *ContaCorrente) Deposito(valor float64) (string, float64) {
	if valor > 0 {
		conta.saldo += valor
		return "Deposito realizado com sucesso!", conta.saldo
	} else {
		return "Valor inválido!", valor
	}
}
func (contaT *ContaCorrente) Transfere(contaR *ContaCorrente, valor float64) bool {
	if contaT.saldo >= valor && valor > 0 {
		contaT.saldo -= valor
		contaR.Deposito(valor)
		return true
	} else {
		return false
	}
}
func (conta *ContaCorrente) Extrato() float64 {
	return conta.saldo
}
