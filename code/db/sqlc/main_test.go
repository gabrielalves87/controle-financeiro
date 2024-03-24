package db

import (
	"database/sql"
	"log"
	"testing"
	_ "github.com/lib/pq"
	"os"
	
)


const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/go_finance?sslmode=disable"
)

var testQueries *Queries


func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatalf("%s", err)
	}
	testQueries = New(conn) 
	os.Exit(m.Run())
}
