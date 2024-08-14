package mysql

import (
	"fmt"
	"github.com/goodrain/rainbond-task-plug/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(dbConfig *config.DBConfig) error {
	logrus.Infof("init db")
	var err error
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
	db, err = gorm.Open(mysql.Open(dbSource), &gorm.Config{})
	logrus.Infof("init db success")
	return err
}

func GetDB() *gorm.DB {
	return db
}
