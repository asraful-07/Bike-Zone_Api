package main

import (
	"bike_zone_api/internal/config"
	"bike_zone_api/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db := config.ConnectDatabase(cfg)

	server.StartServer( db, cfg)
}
