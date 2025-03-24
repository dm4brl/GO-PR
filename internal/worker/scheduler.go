package worker

import (
    "github.com/robfig/cron/v3"
)

func StartScheduler() {
    c := cron.New()

    c.AddFunc("@daily", func() {
        // проверяем, шаббат/йом тов, формируем задания
    })

    c.Start()
}
