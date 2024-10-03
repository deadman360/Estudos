package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func leArquivo(arquivo string) int {
	conteudo, _ := os.Open(arquivo)
	defer conteudo.Close()
	var ultimaLinha string
	leitor := bufio.NewReader(conteudo)
	conteudo.Seek(0, io.SeekEnd)
	line, _ := leitor.ReadString()
	line = strings.TrimSpace(line)
	inicio, _ := strconv.Atoi(line)
	return inicio
}
