package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(d *gorm.DB) {
	db = d
	// FIX:暂时关闭AutoMigrate
	// err := AutoMigrate(new(model.Port), new(model.PortTrafficMonthly))
	// if err != nil {
	// 	log.Fatalf("failed migrate database: %s", err.Error())
	// }
}

func AutoMigrate(dst ...interface{}) error {
	var err error
	err = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(dst...)
	return err
}

func GetDb() *gorm.DB {
	return db
}

func Close() {
	log.Info("closing db")
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("failed to get db: %s", err.Error())
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Errorf("failed to close db: %s", err.Error())
		return
	}
}
