package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"urlShortener/model"
)

func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./links.db")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Link{})
}
