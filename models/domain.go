package models

import "gorm.io/gorm"

type Domain struct {
	Model
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"unique;size:255" json:"name" binding:"required"`
	Type     string `gorm:"unique;size:64" json:"type" binding:"required"` //master,slave,hit,forward
	TTL      uint   `gorm:"column:ttl" json:"ttl" binding:"required"`
	Provider string `gorm:"column:provider" provider:"provider" binding:"required"`
	Contact  string `gorm:"column:contract" json:"contract" binding:"required"`
	Serial   uint   `gorm:"column:serial" json:"serial" binding:"required"`
	Refresh  uint   `gorm:"column:refresh" json:"refresh" binding:"required"`
	Retry    uint   `gorm:"column:retry" json:"retry" binding:"required"`
	Expire   uint   `gorm:"column:expire" json:"expire" binding:"required"`
	Mininum  uint   `gorm:"column:minium" json:"mininum" binding:"required"`
	Comment  string `gorm:"column:comment" json:"comment" binding:"required"`
	Disabled bool   `gorm:"column:name;default:False" json:"disabled"`
}

func (Domain) TableName() string {
	return "domains"
}

// CheckDomainExist 查询域名是否存在
// 返回值：存在返回true，不存在返回false
func (d Domain) CheckDomainExist(name string) bool {
	db.Select("id").Where("name=?", name).First(&d)
	if d.ID > 0 {
		return true
	} else {
		return false
	}
}

// CreateDomain 新增域名
func CreateDomain(data *Domain) (err error) {
	// 插入域名
	if err = db.Create(&data).Error; err != nil {
		return
	}
	return nil
}

// GetDomains 获取域名列表
// pageNum 当前页数
// pageSize 页的条数
func GetDomains(pageSize int, pageNum int) []Domain {
	var domains []Domain
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&domains).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return domains
}

// GetDomainById 获取单个域名
func GetDomainById(id int) (domain Domain, err error) {
	if err = db.Where("id=?", id).First(&domain).Error; err != nil {
		return domain, err
	}
	return domain, err
}

// EditDomain 编辑域名
func EditDomain(id int, data *Domain) (err error) {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["comment"] = data.Comment

	if err = db.Where("id=?", id).Model(&Domain{}).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDomain 删除域名
func DeleteDomain(id int) (err error) {
	err = db.Where("id=?", id).Delete(&Domain{}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateDomainById(id int, v interface{}) (err error) {
	if err = db.Model(&Domain{}).Where("id = ?", id).Updates(v).Error; err != nil {
		return err
	}
	return nil
}
