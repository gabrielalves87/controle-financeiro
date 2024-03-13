package database

import (
	"log"

	"github.com/gabrielalves87/controle-financeiro/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConexaoComBancoDeDados() {
	stringDeConexao := "host=192.168.8.157 user=postgres password=password dbname=gofinance port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err := gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Error ao Conectar com o Banco de Dados")
	}
	// Crie as tabelas
	err = DB.AutoMigrate(&model.User{}, &model.Categorie{}, &model.Account{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration successful")
}
