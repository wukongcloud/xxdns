package models

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type View struct {
	Model
	ID       int   `gorm:"column:id;primarykey;index;" json:"id"`
	Name     string `gorm:"column:name;unique;not null;size:32" json:"name" validate:"required"`
	Comment  string `gorm:"column:comment;size:512" json:"comment"`
	Disabled bool   `gorm:"column:disabled;default:False" json:"disabled"`
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

// GetViewById 获取单个视图
func GetViewById(id int) (view View, err error) {
	if err = db.Where("id=?", id).First(&view).Error; err != nil {
		return view, err
	}
	return view, err
}

// EditView 编辑视图
func EditView(id int, data *View) (err error) {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["comment"] = data.Comment

	if err = db.Where("id=?", id).Model(&View{}).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

// DeleteView 删除视图
func DeleteView(id int) (err error) {
	err = db.Where("id=?", id).Delete(&View{}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateViewById(id int, v interface{}) (err error) {
	if err = db.Model(&View{}).Where("id = ?", id).Updates(v).Error;err !=nil{
		return err
	}
	return nil
}