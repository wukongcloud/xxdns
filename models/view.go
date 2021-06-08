package models

type View struct {
	Model
	Name string `gorm:"unique;size:16" json:"name" binding:"required"`
	Comment    string `gorm:"unique;size:128" json:"email" binding:"required"`
	Disabled   bool   `gorm:"default:False" json:"-"`
}
