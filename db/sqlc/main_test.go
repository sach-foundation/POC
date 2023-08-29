package db

import (
	"GOLANG/db/util"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
// )

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	// var err error
	// testDb, err = sql.Open(dbDriver, dbSource)
	testDb, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
