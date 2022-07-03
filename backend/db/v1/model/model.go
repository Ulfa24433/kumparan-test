package model

import (
	"time"
)

type Article struct {
	ID      int       `gorm:"column:id"`
	Author  string    `gorm:"column:author"`
	Title   string    `gorm:"column:title"`
	Body    string    `gorm:"column:body"`
	Created time.Time `gorm:"column:created"`
}
