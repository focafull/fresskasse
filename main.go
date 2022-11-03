package main

import (
	"akaflieg/fresskasse/v2/config"
	"akaflieg/fresskasse/v2/db"
	"akaflieg/fresskasse/v2/server"
)

func main() {
	config.Init("config.yaml")
	db.InitDB()
	db.SetupDB()
	server.Init()
}
