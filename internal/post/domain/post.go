package domain

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	PostId   uuid.UUID  `json:"post_id,omitempty" gorm:"primary_key"`
	TagId    *uuid.UUID `json:"tag_id"`
	Title    string     `json:"title"`
	AuthorId string     `json:"author_id"`
	Content  string     `json:"content"`
	CreateAt time.Time  `json:"create_at,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
}
