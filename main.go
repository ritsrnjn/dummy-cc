package main

import (
	"net/http"

	"ritsrnjn/dummy-cc/config"
	"ritsrnjn/dummy-cc/routes"
	"ritsrnjn/dummy-cc/sqldb"
)

func main() {
	// read env configs
	config.ReadConfigs()

	// make a connection with the database
	err := sqldb.ConnectWithDb()
	if err != nil {
		panic(err)
	}

	router := routes.GetRouter()
	http.ListenAndServe(":3000", router)
}
