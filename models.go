package main

import (
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Event struct {
	id          int
	title       string
	start_date  time.Time
	end_date    time.Time
	last_notify time.Time
	// daily, weekly, monthly, yearly, custom
	repeat string
}
