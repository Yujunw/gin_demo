package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化一个全局的DB
var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	dsn := "root:123456@(139.155.239.206:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
