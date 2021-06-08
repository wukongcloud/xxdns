package models

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint         `gorm:"primarykey" json:"-"`
	CreatedAt *time.Time   `json:"-"`
	UpdatedAt *time.Time   `json:"-"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}