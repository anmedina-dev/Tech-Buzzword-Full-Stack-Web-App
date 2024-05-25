package main

import (
	"tech-buzzword-service/cron"
	"tech-buzzword-service/db"
	"tech-buzzword-service/models"
	"tech-buzzword-service/server"
)

func main() {
	db.Init()
	models.Init()
	cron.Init()
	server.Init()
}
