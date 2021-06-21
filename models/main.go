package models

import (
	"fmt"
	"github.com/wukongcloud/xxdns/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

//
var db *gorm.DB
var err error

func InitDb() {
	config,err:=utils.ReadConfig()
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	DSN := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",config.DBUser,config.DBPwd,config.DBHost,config.DBName)

	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	_ = db.AutoMigrate(&View{},&Acl{},&IPDB{},&Record{})

	sqlDB, _ := db.DB()

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
