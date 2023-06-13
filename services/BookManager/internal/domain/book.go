package domain

import (
	"time"
)

type Book struct {
	Id       int64
	Title    string
	Author   string
	Year     int32
	Genres   []string
	CreateAt time.Time
}
