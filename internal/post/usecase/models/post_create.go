package models

import (
	"github.com/google/uuid"
)

type PostCreate struct {
	TagId    *uuid.UUID `json:"tag_id"`
	Title    string     `json:"title"`
	AuthorId string     `json:"author_id"`
	Content  string     `json:"content"`
}
