package entity

import "time"

type User struct {
	Id          string
	UserID      string
	FullName    string
	PhoneNumber string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
