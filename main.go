package main

import (
	"GOLANG/.vscode/api"
	db "GOLANG/db/sqlc"
	"GOLANG/db/util"
	"database/sql"
	_ "database/sql"
	"log"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "localhost:9000"
// )

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Unable to load config file", err)
	}
	// conn, err := sql.Open(dbDriver, dbSource)
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start the server", err.Error())
	}
}
