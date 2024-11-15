package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func RunCrons(duck *sql.DB, mysql *sql.DB) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	_, err = scheduler.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(23, 59, 59), // Start log export daily at 23:59:59
			),
		),
		gocron.NewTask(
			func() {
				query := fmt.Sprintf(`INSERT INTO mysql_db.logs
					SELECT client_ip, timestamp, method, path, protocol,
					status_code, latency, user_agent, created_at
					FROM gopi.logs WHERE timestamp < %s;
				`, strconv.FormatInt(time.Now().Unix(), 10))
				_, err = duck.Exec(query)

				if err != nil {
					fmt.Printf("Error creating logs table: %v\n", err)
					return
				}
			},
			"Export Access logs",
			1,
		),
	)

	if err != nil {
	}

	go scheduler.Start()
}
