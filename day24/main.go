package main

import "gorm.io/gorm"
import "gorm.io/driver/mysql"

// 1. 定义模型
type User struct {
	gorm.Model
	Name   string
	No     int64
	Active bool
}

func main() {
	// 2. 连接数据库
	dsn := "root:123456@(139.155.239.206:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	// 3. 将模型与数据库中的表对应起来
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}

	// 4. 创建
	//u1 := User{Name: "Rando", No: 9, Active: true}
	//db.Create(&u1)
	//u2 := User{Name: "Kawhi", No: 2, Active: true}
	//db.Create(&u2)
	//u3 := User{Name: "Kobe", No: 24, Active: false}
	//db.Create(&u3)

	// 5. 删除
	//如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！
	//当调用 Delete 方法时， 记录不会真正的从数据库中被删除， 只会将DeletedAt 字段的值会被设置为当前时间
	//var u = User{}
	//u.ID = 1
	//db.Debug().Delete(&u)

	//删除一条记录时，删除对象需要指定主键，否则会触发 批量 Delete
	// 此时无法删除Name = "kobe"的记录，需要使用Where语句
	//var u = User{}
	//u.Name = "Kobe"
	//db.Debug().Delete(&u)
	// UPDATE `users` SET `deleted_at`='2021-06-05 09:46:57.47' WHERE name = 'Kobe' AND `users`.`deleted_at` IS NULL
	//db.Debug().Where("name = ?", "Kobe").Delete(&u)

	// 物理删除，使用Unscoped
	// DELETE FROM `users` WHERE name = 'Kobe'
	db.Debug().Unscoped().Where("name = ?", "Kobe").Delete(User{})

}
