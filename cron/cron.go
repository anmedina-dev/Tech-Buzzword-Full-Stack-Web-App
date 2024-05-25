package cron

import (
	"fmt"
	"os"
	"tech-buzzword-service/models"

	"github.com/robfig/cron/v3"
)

type fn func()

func Init() {
	fmt.Println("Initializing Cron")
	StartCron(models.Init)
}

func StartCron(cronFunction fn) {
	c := cron.New()
	c.AddFunc(os.Getenv("CRON"), func() {
		cronFunction()
	})
	c.Start()
}
