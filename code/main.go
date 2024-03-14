package main

import (
	"log"

	"github.com/gabrielalves87/controle-financeiro/database"
	"github.com/gabrielalves87/controle-financeiro/models"
)

func main() {
	database.ConexaoComBancoDeDados()
	log.Println("Tudo bem!!!")
	usuario := models.User{Username: "Gabriel", Password: "batatinha", Email: "gabriel.alves@testes.com.br"}
	err := database.DB.Create(usuario)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Usuario criado com sucesso")

}
