package main

import (
	_ "gorm.io/gorm"
	"time"
)

type Event struct {
	ID         int    `gorm:"primary_key"`
	Title      string `gorm:"not null"`
	StartDate  time.Time
	EndDate    time.Time
	LastNotify time.Time
	NextNotify time.Time
	// daily, weekly, monthly, yearly, custom
	Repeat string
}
