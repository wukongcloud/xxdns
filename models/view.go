package models

import "gorm.io/gorm"

type View struct {
	Model
	Name string `gorm:"unique;size:16" json:"name" binding:"required"`
	Comment    string `gorm:"unique;size:128" json:"email" binding:"required"`
	Disabled   bool   `gorm:"default:False" json:"-"`
}

//查询视图是否存在
func CheckView(name string) int {
	var views View
	db.Select("id").Where("name=?",name).First(&views)
	// 视图已存在
	if views.ID > 0 {
		return 409
	}
	return 200
}

//新增视图
func CreateView(data *View) (code int) {
	// 插入视图
	err := db.Create(&data).Error
	if err != nil {
		return 500
	}
	return 201
}

// 获取视图列表
// pageNum 当前页数
// pageSize 页的条数
func GetUsers(pageSize int, pageNum int )[]View{
	var views []View
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&views).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return views
}

//编辑视图
func EditUsers( id int, data *View) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["comment"] = data.Comment

	err = db.Where("id=?", id).Model(&View{}).Updates(maps).Error
	if err != nil {
		return 500
	}
	return 200
}

//删除视图
func DeleteView(id int) int {
	err = db.Where("id=?",id).Delete(&View{}).Error
	if err != nil {
		return 500
	}
	return 200
}