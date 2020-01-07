package utils

import (
	"github.com/jinzhu/gorm"
	"time"
	db2 "urlShortener/db"
)

func CleanDB(db *gorm.DB) {
	for {
		db2.Cleanup(db)
		time.Sleep(time.Hour)
	}
}
