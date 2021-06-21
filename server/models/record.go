package models

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Record struct {
	Model
	ID       int    `gorm:"column:id;primarykey;index;" json:"id"`
	Domain   string `gorm:"column:domain" json:"domain" validate:"required"`
	View     string `gorm:"column:view" json:"view" validate:"required"`
	Name     string `gorm:"column:name" json:"name" validate:"required"`
	Type     string `gorm:"column:type" json:"type" validate:"required"`
	Content  string `gorm:"column:content" json:"content" validate:"required"`
	TTL      uint   `gorm:"column:ttl" json:"ttl" validate:"required"`
	Priority uint   `gorm:"column:priority" json:"priority"`
	Disabled bool   `gorm:"column:disabled;default:False" json:"disabled"`
	Comment  string `gorm:"column:comment" json:"comment"`
	//Acls     []Acl  `gorm:"-"`
}

func (Record) TableName() string {
	return "records"
}

// CheckRecordExist 查询记录是否存在
// 返回值：存在返回true，不存在返回false
func (r Record) CheckRecordExist(name string) bool {
	db.Select("id").Where("name=?", name).First(&r)
	if r.ID > 0 {
		return true
	} else {
		return false
	}
}

// CreateRecord 新增记录
func CreateRecord(data *Record) (err error) {
	// 插入记录
	if err = db.Create(&data).Error; err != nil {
		return
	}
	return nil
}

// GetRecords 获取记录列表
// pageNum 当前页数
// pageSize 页的条数
func GetRecords(pageSize int, pageNum int) []Record {
	var records []Record
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&records).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return records
}

// GetRecordById 获取单个记录
func GetRecordById(id int) (record Record, err error) {
	if err = db.Where("id=?", id).First(&record).Error; err != nil {
		return record, err
	}
	return record, err
}

// EditRecord 编辑记录
func EditRecord(id int, data *Record) (err error) {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	maps["comment"] = data.Comment

	if err = db.Where("id=?", id).Model(&Record{}).Updates(maps).Error; err != nil {
		return err
	}
	return nil
}

// DeleteRecord 删除记录
func DeleteRecord(id int) (err error) {
	err = db.Where("id=?", id).Delete(&Record{}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecordById(id int, r interface{}) (err error) {
	if err = db.Model(&Record{}).Where("id = ?", id).Updates(r).Error; err != nil {
		return err
	}
	return nil
}
