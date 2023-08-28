package main

import (
	"GOLANG/.vscode/api"
	db "GOLANG/db/sqlc"
	"database/sql"
	_ "database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "localhost:9000"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot start the server", err.Error())
	}
}
