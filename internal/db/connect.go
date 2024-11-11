package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	// 构建连接字符串：用户名:密码@协议(地址)/数据库名
	dsn := "root:root@tcp(127.0.0.1:3306)/lcu_helper"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
	})

	fmt.Println("Connected to the database successfully!")
}
