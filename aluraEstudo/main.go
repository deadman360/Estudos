package main

import "fmt"

type estrutura struct {
	nome, data, sobre string
}

func main() {
	est := estrutura{"marcos", "kakakak", "papa"}
	est = trocaEstrutura(est)
	est2 := estrutura{"marcos", "kakakak", "papa"}
	trocaComPonteiro(&est2)

	fmt.Println(est, "\n", est2)
}

func trocaEstrutura(est estrutura) estrutura {
	est = estrutura{"jao", "hoej", "bab"}
	return est
}

func trocaComPonteiro(est *estrutura) {
	*est = estrutura{"jao", "hoje", "bab"}
}
