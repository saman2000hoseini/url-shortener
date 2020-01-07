package handler

import (
	"github.com/jinzhu/gorm"
	"strings"
	"urlShortener/model"
)

func ClientHandler(c *model.Client, db *gorm.DB) {
	for {
		request, err := c.Reader.ReadString('\n')
		if err != nil {
			c.Connection.Close()
			return
		}
		req := strings.Split(request, ":")
		switch strings.ToLower(req[0]) {
		case "url":
			url := strings.ToLower(req[1])
			url = strings.Replace(url, "www.", "", 1)
			if !strings.Contains(url, "http://") {
				url = "http://" + url
			}
			link := model.NewLink(url)
			link.AddToDB(db)
			c.Writer.WriteString(link.Shortened)
		case "shortened":
			link
		}
	}
}
