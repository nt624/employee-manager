package main

import (
	"example/employee-manager/db"
	"example/employee-manager/router"
)

func main() {
	dbConn := db.Init()
	router.Router(dbConn)
}
