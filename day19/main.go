package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 对应数据库中的表table
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// mysql的登录名:密码@(ip:mysql映射的端口)/db名称
	dsn := "root:123456@(139.155.239.206:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}

	// 已经没有了close()方法
	// defer db.Close()

	// 根据Go Struct结构自动生成对应的表结构
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		return
	}

	// 创建数据行
	//u1 := UserInfo{1, "June", "male", "basketball"}
	//db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("%v\n", u)
	// 更新
	db.Model(&u).Update("hobby", "game")
	// 删除
	db.Delete(&u)

}
