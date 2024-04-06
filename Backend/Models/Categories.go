package models

import "time"

type Categories struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	UserId    int       `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
