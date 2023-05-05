package model

import "time"

type User struct {
	Username        string
	Password        string
	PasswordConfirm string
	Email           string
	RoleId          int32
	UpdatedAt       time.Time
	CreatedAt       time.Time
}
