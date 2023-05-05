package model

type UserUpdate struct {
	Username    string
	NewUsername *string
	Password    *string
	Email       *string
	RoleId      *int32
}
