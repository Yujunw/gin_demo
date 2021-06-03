package main

import (
	"fmt"
	"gorm.io/driver/mysql"
)
import "gorm.io/gorm"

// 定义模型
type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	// 连接数据库，mysql的登录名:密码@(ip:mysql映射的端口)/db名称
	dsn := "root:123456@(139.155.239.206:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	// 将模型与数据库中的表对应起来
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}

	// 创建
	//u1 := User{Name: "Rando", Age: 9}
	//db.Create(&u1)
	//u2 := User{Name: "Kawhi", Age: 2}
	//db.Create(&u2)

	// 查询
	//var user User
	//db.First(&user)
	//fmt.Printf("user: %#v\n", user)
	//
	//var users []User
	//db.Find(&users)
	//fmt.Printf("users: %#v\n", users)

	var user User
	db.FirstOrInit(&user, User{Name: "George"})
	fmt.Printf("user: %#v\n", user)
}
