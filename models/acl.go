package models

import (
	"gorm.io/gorm"
)

type Acl struct {
	Model
	ID       int    `gorm:"primarykey" json:"id"`
	Name     string `gorm:"unique;size:16" json:"name" binding:"required"`
	Filter   string
	Comment  string `gorm:"size:512" json:"comment,omitempty"`
	Disabled bool   `gorm:"default:False" json:"disabled"`
}

func (Acl) TableName() string {
	return "acls"
}

// CheckAclExist 查询访问控制列表是否存在
// 返回值：存在返回true，不存在返回false
func (a Acl) CheckAclExist(name string) bool {
	db.Select("id").Where("name=?", name).First(&a)
	if a.ID > 0 {
		return true
	} else {
		return false
	}
}

// CreateAcl 新增访问控制列表
func CreateAcl(data *Acl) (err error) {
	// 插入访问控制列表
	if err = db.Create(&data).Error; err != nil {
		return
	}
	return nil
}

// GetAcls 获取访问控制列表列表
// pageNum 当前页数
// pageSize 页的条数
func GetAcls(pageSize int, pageNum int) (err error, acls []Acl) {
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&acls).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return err, acls
	}
	return nil, acls
}

// GetAclById 获取单个访问控制列表
func GetAclById(id int) (acl Acl, err error) {
	if err = db.Where("id=?", id).First(&acl).Error; err != nil {
		return acl, err
	}
	return acl, err
}

// EditAcl 编辑访问控制列表
func EditAcl(id int, data *Acl) (err error) {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["comment"] = data.Comment

	if err = db.Where("id=?", id).Model(&Acl{}).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAcl 删除访问控制列表
func DeleteAcl(id int) (err error) {
	err = db.Where("id=?", id).Delete(&Acl{}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateAclById(id int, v interface{}) (err error) {
	if err = db.Model(&Acl{}).Where("id = ?", id).Updates(v).Error; err != nil {
		return err
	}
	return nil
}
