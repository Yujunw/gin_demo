package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	dsn := "root:123456@(139.155.239.206:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}

	//u1 := User{Name: "Rando", Age: 9, Active: true}
	//db.Create(&u1)
	//u2 := User{Name: "Kawhi", Age: 2, Active: false}
	//db.Create(&u2)

	var user User
	db.First(&user)
	fmt.Printf("user: %v\n", user)

	user.Name = "George"
	user.Age = 7
	//UPDATE `users` SET `created_at`='2021-06-04 07:55:39.1',`updated_at`='2021-06-04 07:57:55.248',`deleted_at`=NULL,`name`='George',`age`=7,`active`=true WHERE `id` = 1
	db.Debug().Save(&user) // 默认会修改所有字段

	// UPDATE `users` SET `name`='qiaozhi',`updated_at`='2021-06-04 08:01:46.977' WHERE `id` = 1
	db.Debug().Model(&user).Update("name", "qiaozhi") // 只更新一个指定字段

	m1 := map[string]interface{}{
		"name":   "paojiao",
		"age":    14,
		"active": true,
	}
	// UPDATE `users` SET `active`=true,`age`=14,`name`='paojiao',`updated_at`='2021-06-04 08:05:08.016' WHERE `id` = 1
	db.Debug().Model(&user).Updates(m1)                // 更新m1列出来的所有字段
	db.Debug().Model(&user).Select("age").Updates(m1)  // 只更新age字段
	db.Debug().Model(&user).Omit("active").Updates(m1) // 更新active以外的其他字段

	// UPDATE `users` SET `age`=30 WHERE `id` = 1
	db.Debug().Model(&user).UpdateColumn("age", 30) // 只更新age，不更新钩子函数

	// 让Users表中所有用户的age都加2
	db.Debug().Model(&User{}).Update("age", gorm.Expr("age+?", 2))

}
