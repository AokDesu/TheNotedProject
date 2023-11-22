package models

import "time"

type Note struct {
	Id        int       `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Detail    string    `json:"detail,omitempty"`
	UserId    int       `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type NoteImage struct {
	Id     int    `json:"id,omitempty"`
	Images []byte `json:"images,omitempty"`
	NoteId int    `json:"note_id,omitempty"`
}
