package main

import "urlShortener/router"

func main() {
	r := router.Router()
	r.Logger.Fatal(r.Start("localhost:65431"))
}
