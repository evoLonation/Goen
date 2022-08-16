package main

import (
	"Cocome/api"
	"Cocome/configuration"
)

func main() {
	config, err := configuration.LoadConfig("./")
	if err != nil {
		print(err.Error())
	}
	api.Start(config.ServerAddress)

}
