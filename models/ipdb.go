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

// CreateIPDB 新增IP地址段
func CreateIPDB(data *IPDB) (err error) {
	// 插入IP地址段
	if err = db.Create(&data).Error; err != nil {
		return
	}
	return nil
}

// GetIPDB 获取IP地址段列表
// pageNum 当前页数
// pageSize 页的条数
func GetIPDB(pageSize int, pageNum int, country string, province string, isp string) []IPDB {
	var ipdb []IPDB
	dbIP := db.Model(IPDB{})
	if country != ""{
		dbIP.Where("country=?",country)
	}
	if province != ""{
		dbIP.Where("province=?",province)
	}
	if isp !=""{
		dbIP.Where("isp=?",isp)
	}
	err = dbIP.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&ipdb).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return ipdb
}
