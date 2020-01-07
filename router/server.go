package router

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"urlShortener/db"
	"urlShortener/handler"
	"urlShortener/utils"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Router() *echo.Echo {
	e := echo.New()
	db.New()
	db.Migrate()
	go utils.CleanDB()
	t := &Template{template.Must(template.ParseGlob("html/*.html"))}
	e.Renderer = t
	e.GET("/", handler.HomePage)
	e.POST("/generate", handler.AddUrl)
	e.GET("/:url", handler.RedirectUrl)
	return e
}
