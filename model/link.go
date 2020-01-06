package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model
	Url       string `gorm:"primary_key"`
	Shortened string `gorm:"unique_index;not null"`
}

func NewLink(u string) *Link {
	uEnc := base64.RawURLEncoding.EncodeToString([]byte(u))
	return &Link{Url: u, Shortened: uEnc}
}

func (l *Link) AddToDB(db *gorm.DB) {
	if db.NewRecord(l) {
		db.Create(l)
	} else {
		db.Where("Url = ?", l.Url).First(l)
	}
}

func (l *Link) GetUrl(db *gorm.DB) {
	db.Where("Shortened = ?", l.Shortened).First(l)
}
