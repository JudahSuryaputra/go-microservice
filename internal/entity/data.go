package entity

import "time"

type Data struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
