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
	Id        int       `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	Role      UserRole  `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
