package models

import "gorm.io/gorm"

type View struct {
	Model
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"unique;size:16" json:"name" binding:"required"`
	Comment  string `gorm:"unique;size:128" json:"comment" binding:"required"`
	Disabled bool   `gorm:"default:False" json:"disabled"`
}

func (View) TableName() string {
	return "views"
}

// CheckViewExist 查询视图是否存在
// 返回值：view存在返回true，不存在返回false
func (v View) CheckViewExist(name string) bool {
	db.Select("id").Where("name=?", name).First(&v)
	if v.ID > 0 {
		return true
	} else {
		return false
	}
}

// CreateView 新增视图
func CreateView(data *View) (err error) {
	// 插入视图
	if err = db.Create(&data).Error; err != nil {
		return
	}
	return nil
}

// GetViews 获取视图列表
// pageNum 当前页数
// pageSize 页的条数
func GetViews(pageSize int, pageNum int) []View {
	var views []View
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&views).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return views
}

// EditView 编辑视图
func EditView(id int, data *View) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["comment"] = data.Comment

	err = db.Where("id=?", id).Model(&View{}).Updates(maps).Error
	if err != nil {
		return 500
	}
	return 200
}

// DeleteView 删除视图
func DeleteView(id int) int {
	err = db.Where("id=?", id).Delete(&View{}).Error
	if err != nil {
		return 500
	}
	return 200
}
