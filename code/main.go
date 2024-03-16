package main

import (
	"log"

	"github.com/gabrielalves87/controle-financeiro/database"
	"github.com/gabrielalves87/controle-financeiro/routes"

)



func main() {
	database.ConexaoComBancoDeDados()
	log.Println("Tudo bem!!!")
	routes.Handlerequest()
}
