package main

import (
	"Project-2/database"
	"Project-2/router"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	router.StartServer().Run(PORT)

}
