package models
type Domain struct {
	Model
	Name string `gorm:"unique;size:255" json:"name" binding:"required"`
	Type    string `gorm:"unique;size:64" json:"type" binding:"required"`
	TTL string `gorm:"unique;size:16" json:"ttl" binding:"required"`
	Provider string `gorm:"unique;size:16" provider:"name" binding:"required"`
	Contact string `gorm:"unique;size:16" json:"contract" binding:"required"`
	Serial string `gorm:"unique;size:16" json:"serial" binding:"required"`
	Refresh string `gorm:"unique;size:16" json:"refresh" binding:"required"`
	Retry string `gorm:"unique;size:16" json:"retry" binding:"required"`
	Expire string `gorm:"unique;size:16" json:"expire" binding:"required"`
	Mininum string `gorm:"unique;size:16" json:"mininum" binding:"required"`
	Comment string `gorm:"unique;size:128" json:"comment" binding:"required"`
	Disabled   bool   `gorm:"default:False" json:"-"`
}
