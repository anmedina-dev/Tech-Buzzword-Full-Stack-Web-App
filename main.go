package main

import (
	"tech-buzzword-service/buzzword"
	"tech-buzzword-service/cron"
	"tech-buzzword-service/db"
	"tech-buzzword-service/server"
)

func main() {
	db.Init()
	buzzword.Init()
	cron.Init()
	server.Init()
}
