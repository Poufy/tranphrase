package main

import (
	"tranphrase/config"
	"tranphrase/db"
	"tranphrase/routes"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	router := routes.SetupRouter()

	// Serve static files
	router.Static("/", "./static")

	router.Run(config.Cfg.SERVER.Port)
}
