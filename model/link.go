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
