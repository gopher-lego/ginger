package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
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
	db, err := gorm.Open("mysql", viper.Get("database.mysql.dsn"))
	if err != nil {
		panic("MySql Connect error")
	}

	return db
}