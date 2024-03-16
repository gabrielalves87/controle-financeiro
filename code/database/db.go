package database

import (
	"log"

	"github.com/gabrielalves87/controle-financeiro/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConexaoComBancoDeDados() {
	stringDeConexao := "host=192.168.8.157 user=postgres password=password dbname=gofinance port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{TranslateError: true})
	print(DB)
	if err != nil {
		log.Panic("Error ao Conectar com o Banco de Dados")
	}
	// Crie as tabelas
	err = DB.AutoMigrate(
		&models.User{},
		&models.Categorie{},
		&models.Account{},
		&models.Trasactions{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration successful")
}
func CriaUsuario(user *models.User) {
	stringDeConexao := "host=192.168.8.157 user=postgres password=password dbname=gofinance port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err := gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	result := DB.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return // retornar imediatamente se houver um erro
	}
}
