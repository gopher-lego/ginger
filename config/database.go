package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var MyDB *gorm.DB

func init() {
}

func MySqlInit() {
	mySqlConnect()

	// Dump sql to screen on DebugMode, or use db.Debug().Find(&page)
}

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func mySqlConnect() {
	var err error
	dsn := viper.GetString("database.mysql.dsn")
	MyDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("MySql Connect error")
	}

	// https://gorm.io/docs/generic_interface.html
	// Connection Pool
	sqlDB, _ := MyDB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
