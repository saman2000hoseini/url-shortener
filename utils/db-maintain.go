package utils

import (
	"time"
	"urlShortener/db"
)

func CleanDB() {
	for {
		db.Cleanup()
		time.Sleep(time.Hour)
	}
}
