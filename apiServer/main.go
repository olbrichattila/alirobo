package main

import (
	"aliserver/server"
	"aliserver/storage"
	"fmt"
)

func main() {
	store := storage.New()
	err := store.Init()
	if err != nil {
		fmt.Println(err.Error())
	}
	server := server.New(store)
	server.Serve()
}
