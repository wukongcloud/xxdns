package models

import (
	_ "github.com/go-playground/validator/v10"
)

type Server struct {
	Model
	ID       int    `gorm:"column:id;primarykey;index;" json:"id"`
	IP       string `gorm:"column:ip;unique;not null;size:32" json:"name"`
	Comment  string `gorm:"column:comment;size:512" json:"comment,omitempty"`
	Disabled bool   `gorm:"column:disabled;default:False" json:"disabled"`
}

func (Server) TableName() string {
	return "views"
}

// CheckServerExist 查询视图是否存在
// 返回值：view存在返回true，不存在返回false
func (s Server) CheckServerExist(name string) bool {
	db.Select("id").Where("name=?", name).First(&s)
	if s.ID > 0 {
		return true
	} else {
		return false
	}
}

// CreateServer 新增服务器
func CreateServer(data *Server) (err error) {
	if err = db.Create(&data).Error; err != nil {
		return
	}
	return nil
}
