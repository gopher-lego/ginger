package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
}

func MySqlInit() *gorm.DB {
	db := mySqlConnect()

	// Dump sql to screen on DebugMode, or use db.Debug().Find(&page)
	// db.LogMode(true)

	return db

	// @todo Why invalid ？
	// db.DB().SetMaxIdleConns(5)
	// db.DB().SetMaxOpenConns(100)
	// db.DB().SetConnMaxLifetime(time.Minute)
}

func mySqlConnect() *gorm.DB {
	dsn := viper.GetString("database.mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("MySql Connect error")
	}

	return db
}