package model

import uuid "github.com/satori/go.uuid"

type Book struct {
	ID   uuid.UUID `gorm:"column:id;primary_key" json:"id"`
	Name string    `gorm:"column:name" json:"name"`
}
