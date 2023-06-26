package main

import (
	"database/sql"
	"log"

	"github.com/NopparootSuree/Learning-GO/api"
	db "github.com/NopparootSuree/Learning-GO/db/sqlc"
	"github.com/NopparootSuree/Learning-GO/utils"
)

func main() {
	config, err := utils.LoadCongfig(".")
	if err != nil {
		log.Fatal("Can't load config:", err)
	}
	conn, err := sql.Open(config.DBDRIVER, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't start server:", err)
	}
}
