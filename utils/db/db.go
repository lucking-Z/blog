package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type ConfigDB struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
}

var Instance *gorm.DB

//db初始化
func InitDB(confObj ConfigDB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&&interpolateParams=true&loc=Local",
		confObj.Username,
		confObj.Password,
		confObj.Host,
		confObj.Port,
		confObj.DbName)
	loggerIns := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
	mysqlConfig := mysql.Config{DSN: dsn}
	gormConfig := &gorm.Config{
		Logger: loggerIns,
	}
	var err error
	if Instance, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		panic(fmt.Sprintf("db init failed:%s", err.Error()))
	}



	sqlDB, err := Instance.DB()
	if err != nil {
		panic(fmt.Sprintf("db init failed:%s", err.Error()))
	}
	//最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	//最大打开数量
	sqlDB.SetMaxOpenConns(100)
	//最大可服务时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}
