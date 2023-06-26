package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/NopparootSuree/Learning-GO/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadCongfig("../..")
	if err != nil {
		log.Fatal("Can't load config:", err)
	}

	testDB, err = sql.Open(config.DBDRIVER, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
