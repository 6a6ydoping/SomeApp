package main

import (
	"someApp/internal/app"
	"someApp/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	err = app.Run(*cfg)
	if err != nil {
		panic(err)
	}
}
