package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type Todo struct {
	gorm.Model
	ID uuid.UUID `gorm:"column:uuid;"`

	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Users struct
type Todos struct {
	Todos []*Todo `json:"todos"`
}

func (todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	todo.ID = uuid.New()
	return
}
