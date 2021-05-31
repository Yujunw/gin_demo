package main

import (
	"database/sql"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

type User struct {
	// 如果字段中有ID。则默认为主键
	ID   int64
	Name sql.NullString `gorm:"default:'jack'"` // 设置默认值jack
	Age  int64
}

func main() {

	// mysql的登录名:密码@(ip:mysql映射的端口)/db名称
	dsn := "root:123456@(139.155.239.206:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}

	// 将模型与数据库中的表对应起来
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}

	// 正常创建
	//u := User{Name:"June", Age:26} // 在代码层面创建一个User对象
	// 不传入Name字段，则使用默认值Jack
	//u := User{Age:27}

	// 如果直接传入一个空值，还是会使用默认值jack
	//u := User{Name: "", Age: 29}

	// 如果想要传入空值，则需要使用指针类型，结构体中需要修改为 *string
	//u := User{Name: new(string), Age:30}

	// 也可以使用Scanner/Valuer接口方式实现零值存入数据库
	u := User{Name: sql.NullString{Valid: true}, Age: 20}
	db.Debug().Create(&u)

}
