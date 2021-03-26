package model

import (
	"goblog/pkg/logger"
	"gorm.io/driver/mysql"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm"
)


// DB gorm.DB 对象
var DB *gorm.DB

func ConnectDB() *gorm.DB  {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:joydata@tcp(10.211.55.6:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB,err = gorm.Open(config,&gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Warn),
	})

	logger.LogError(err)

	return DB
}