package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
)

func main() {
	// 创建数据库 bubble

	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}

	// 绑定模型
	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		fmt.Println("DB.AutoMigrate() failed, ", err)
		return
	}

	// 注册路由
	r := routers.SetupRouter()
	r.Run(":9000")

}
