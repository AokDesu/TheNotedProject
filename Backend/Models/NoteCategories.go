package models

type NoteCategories struct {
	NoteId       int `json:"note_id,omitempty"`
	CategoriesId int `json:"categories_id,omitempty"`
}
