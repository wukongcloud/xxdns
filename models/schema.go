package models

import (
	"database/sql"
	"time"
)

type Model struct {
	CreatedAt *time.Time   `json:"-"`
	UpdatedAt *time.Time   `json:"-"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}