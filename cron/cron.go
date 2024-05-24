package cron

import (
	"fmt"
	"os"
	"tech-buzzword-service/buzzword"

	"github.com/robfig/cron/v3"
)

type fn func()

func Init() {
	fmt.Println("Initializing Cron")
	StartCron(buzzword.Init)
}

func StartCron(cronFunction fn) {
	c := cron.New()
	c.AddFunc(os.Getenv("CRON"), func() {
		cronFunction()
	})
	c.Start()
}
