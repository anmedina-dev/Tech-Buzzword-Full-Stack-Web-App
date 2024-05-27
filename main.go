package main

import (
	"tech-buzzword-service/db"
	"tech-buzzword-service/models"
	"tech-buzzword-service/server"
)

func main() {
	db.Init()
	models.Init()
	server.Init()
}
