package main

import (
	"log"
	"database/sql"
	db "github.com/gabrielalves87/controle-financeiro/db/sqlc"
	_ "github.com/lib/pq"
	"github.com/gabrielalves87/controle-financeiro/api"
	// "github.com/gabrielalves87/controle-financeiro/database"
	// "github.com/gabrielalves87/controle-financeiro/routes"
)


const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	serverAddress string = "0.0.0.0:8000"
)

func main() {
	// database.ConexaoComBancoDeDados()
	// log.Println("Tudo bem!!!")
	// routes.Handlerequest()
	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatalf("%s", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil{
		log.Fatal("cannot start api: ", err)
	}
}
