package main

import (
	"net"
	db2 "urlShortener/db"
	"urlShortener/handler"
	"urlShortener/model"
	"urlShortener/utils"
)

func main() {
	server, err := net.Listen("tcp", ":65431")
	db := db2.New()
	go utils.CleanDB(db)
	if err != nil {
		panic(err)
	}
	for {
		connection, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go handler.ClientHandler(model.NewClient(connection), db)
	}
}
