package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gabrielalves87/controle-financeiro/api"
	db "github.com/gabrielalves87/controle-financeiro/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	// "github.com/gabrielalves87/controle-financeiro/database"
	// "github.com/gabrielalves87/controle-financeiro/routes"
)


func main() {
	// database.ConexaoComBancoDeDados()
	// log.Println("Tudo bem!!!")
	// routes.Handlerequest()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	  }
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	serverAddress := os.Getenv("SERVER_ADDRESS")
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("%s", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api: ", err)
	}
}
