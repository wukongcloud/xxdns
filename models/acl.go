package models

type Acl struct {
	Model
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"unique;size:16" json:"name" binding:"required"`
	Path     string `gorm:"unique;size:128" json:"path" binding:"required"`
	Disabled bool   `gorm:"default:False" json:"disabled"`
}

func (Acl) TableName() string {
	return "acls"
}
