package main

import (
	"fmt"

	"github.com/gabrielalves87/controle-financeiro/database"
)

func main() {
	database.ConexaoComBancoDeDados()
	fmt.Println("Tudo bem!!!")
}
