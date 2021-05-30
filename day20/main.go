package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	// GORM内置了一个gorm.Model结构体。
	// gorm.Model包含了ID, CreatedAt, UpdatedAt, DeletedAt 4个字段
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

// 表名默认就是结构体名称的复数
// 将 User 的表名设置为 `profiles`
func (User) TableName() string {
	return "profiles"
}

func main() {
	// mysql的登录名:密码@(ip:mysql映射的端口)/db名称
	dsn := "root:123456@(139.155.239.206:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}

	err = db.AutoMigrate(&User{})
	err = db.AutoMigrate(&Animal{})
	if err != nil {
		return
	}

	//db.Create()

}
