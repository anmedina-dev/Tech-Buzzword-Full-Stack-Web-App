package cron

import (
	"os"
	"tech-buzzword-service/controllers"

	"github.com/robfig/cron/v3"
)

type fn func()

func Init() {
	controllers.InitBuzzword()
	StartCron(controllers.InitBuzzword)
}

func StartCron(cronFunction fn) {
	c := cron.New()
	c.AddFunc(os.Getenv("CRON"), func() {
		cronFunction()
	})
	c.Start()
}
