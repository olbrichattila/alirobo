package main

import (
	"aliserver/server"
	"aliserver/storage"
)

func main() {
	store := storage.New()
	server := server.New(store)
	server.Serve()
}
