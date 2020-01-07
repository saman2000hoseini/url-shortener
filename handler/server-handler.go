package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"urlShortener/db"
	"urlShortener/model"
)

func AddUrl(c echo.Context) error {
	var err error
	url := c.FormValue("url")
	fmt.Println(url)
	if url == "" {
		err = errors.New("Invalid input!")
		return err
	}
	url = strings.Replace(url, "www.", "", 1)
	if !strings.Contains(url, "http://") {
		url = "http://" + url
	}
	link := model.NewLink(url)
	link.AddToDB(db.MyDB())
	//c.Render(http.StatusOK,"index.html",link)
	err = c.String(http.StatusOK, "Your shortened URL: "+c.Request().Host+"/"+link.Shortened)
	return err
}

func HomePage(c echo.Context) error {
	var err error
	c.Render(http.StatusOK, "index.html", nil)
	return err
}

func RedirectUrl(c echo.Context) error {
	var err error
	link := new(model.Link)
	link.Shortened = c.Param("url")
	link.GetUrl(db.MyDB())
	if link.Url != "" {
		c.Redirect(http.StatusTemporaryRedirect, link.Url)
		return nil
	} else {
		err = c.String(http.StatusBadRequest, "Invalid Request")
	}
	return err
}
