package models

import "gorm.io/gorm"

type IPDB struct {
	Model
	ID       int    `gorm:"primarykey" json:"id"`
	CIDR     string `gorm:"column:cidr" json:"cidr"`
	Country  string `gorm:"column:country" json:"country"`
	Province string `gorm:"column:province" json:"province"`
	ISP      string `gorm:"column:isp" json:"isp"`
	Comment  string `gorm:"size:512" json:"comment,omitempty"`
}

func (IPDB) TableName() string {
	return "ipdb"
}


// CreateIPDB 新增视图
func CreateIPDB(data *IPDB) (err error) {
	// 插入视图
	if err = db.Create(&data).Error; err != nil {
		return
	}
	return nil
}

// GetIPDB 获取视图列表
// pageNum 当前页数
// pageSize 页的条数
func GetIPDB(pageSize int, pageNum int) []IPDB {
	var ipdb []IPDB
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&ipdb).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return ipdb
}
