package domain

import "github.com/google/uuid"

type Tag struct {
	TagId   uuid.UUID `json:"tag_id" gorm:"primary_key"`
	TagName string    `json:"tag_name"`
	Post    []Post    `json:"post" gorm:"foreignkey:TagId"`
}
