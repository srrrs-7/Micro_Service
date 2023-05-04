package model

import "time"

type accessClient struct {
	Id        int
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
