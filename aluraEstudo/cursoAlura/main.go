package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var nome string

func main() {

	introdução()
	comando := menu()
	handler(comando)

}

func introdução() {
	fmt.Println("Bom dia, com quem eu falo?")
	fmt.Scanf("%s", &nome)
	version := 1.6
	fmt.Println("Olá sr(a).", nome)
	fmt.Println("Monitora sites --version", version)
}
func menu() int {
	fmt.Println("1 - Iniciar Monitoramento de sites, por parametro")
	fmt.Println("2 - Iniciar Monitoramento de sites, por arquivo")
	fmt.Println("3 - Exibir Logs")
	fmt.Println("0 - Sair")
	// Captura a opção do usuário
	var comando int
	fmt.Scanf("%d", &comando)
	return comando
}
func handler(comando int) {
	switch comando {
	case 1:
		Monitoramento(PerguntaSites())
	case 2:
		Monitoramento(LeAqruivos())
	case 3:
		ExibeLog()
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Instrução invalida. Saindo...")
		os.Exit(-1)
	}
}

// Logica que percorre slice de sites
func Monitoramento(sites []string) {

	fmt.Println("Sr(a).", nome, "Iniciando o monitoramento para os sites:", sites, "...")
	/* TODO: integrar logica para perguntar por quanto tempo em minutos o usr
	gostaria de ter o monitoramento. */
	for {
		for _, site := range sites {
			Monitora(site)
		}
		time.Sleep(time.Millisecond * 5000)
	}
}

// Pega slice de sites
func PerguntaSites() []string {
	var sites []string
	var resposta string = "s"
	for resposta != "n" {
		switch len(sites) {
		case 0:
			fmt.Print("Bem vindo, Gostaria de continuar? (s/n): ")
		default:
			fmt.Println("Os sites informados até agora são:", sites)
			fmt.Println("Deseja informar mais algum site? (s/n)")
		}
		fmt.Scan(&resposta)
		if resposta == "n" {
			break
		}
		var siteNovo string
		fmt.Println("Informe o site que deseja monitorar:")
		fmt.Scan(&siteNovo)
		sites = append(sites, siteNovo)
	}
	return sites
}

// Logica para monitorar sites
func Monitora(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Site invalido!(", site, "), Erro:", err)
		fmt.Println("Tente novamente")
		Monitoramento(PerguntaSites())
	}
	fmt.Println("O site:", site, "retorna:", resp.Status)
	Logs(site, *resp)

}

// Leitura de arquivo de texto contendo sites
func LeAqruivos() []string {

	//Abre arquivo informado pelo usuario

	fmt.Println("Informe o nome do arquivo que contém os sites que deseja monitorar:")
	var arquivo string
	fmt.Scan(&arquivo)
	conteudo, err := os.Open(arquivo)
	defer conteudo.Close()
	if err != nil {
		fmt.Println("Nome incorreto :: Tente novamente", err)
		Monitoramento(LeAqruivos())
	}
	// Le arquivo utilizando bufio e retorna slice de sites
	leitor := bufio.NewReader(conteudo)
	var listaSites []string
	for {
		line, err := leitor.ReadString('\n')
		line = strings.TrimSpace(line)
		listaSites = append(listaSites, line)
		if err == io.EOF {
			break
		}

	}
	return listaSites
}

// Função para exibir logs
func Logs(site string, resp http.Response) {
	arquivo, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer arquivo.Close()
	if err != nil {
		fmt.Println("Erro ao abrir arquivo log.txt, Erro :", err)
	}
	arquivo.WriteString(time.Now().Format("2006-01-02 15:04:05: ") + site + " -> " + resp.Status + "\n")

}

func ExibeLog() {
	arquivo, err := os.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Falha ao exibir logs, ERRO: ", err, "\n Retornando ao menu inicial...")
		introdução()
	}
	fmt.Println("Logs:\n\n", string(arquivo))
}
