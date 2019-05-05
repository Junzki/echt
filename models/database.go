package models

import (
	"echt/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"runtime"
)


var database *gorm.DB


func InitDbConn(dbCfg *config.DatabaseArgs) error {
	dbType, dbArgs, err := dbCfg.Args()
	if nil != err {
		return err
	}

	database, err = gorm.Open(dbType, dbArgs)
	if nil != err {
		return err
	}

	database.SingularTable(true)
	database.DB().SetMaxOpenConns(655360)
	database.DB().SetMaxIdleConns(20 * runtime.NumCPU())

	database.AutoMigrate(&Content{})

	return nil
}


func GetDbConn() *gorm.DB {
	return database
}
