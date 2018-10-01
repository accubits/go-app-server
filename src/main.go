package main

import (
	cfg "config"
	"db"
	"logger"
	"server"
)

var log = logger.NewLogger("APP")

type TempObj struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

func main() {
	cfg.Init()
	db.Init()
	defer db.CloseDB()
	server.Init()
}
