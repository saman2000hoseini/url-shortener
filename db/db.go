package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
	"urlShortener/model"
)

var db *gorm.DB

func New() *gorm.DB {
	var err error
	db, err = gorm.Open("sqlite3", "./links.db")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func Migrate() {
	db.AutoMigrate(&model.Link{})
}

func Cleanup() {
	db.Unscoped().Delete(&model.Link{}, "CreatedAt < ?", time.Now().Add(-1*time.Hour))
}

func MyDB() *gorm.DB {
	return db
}
