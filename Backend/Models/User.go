package models

import (
	"time"
)

type UserRole string

const (
	Admin  UserRole = "admin"
	Member UserRole = "member"
)

type User struct {
	Id        int
	Username  string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}
