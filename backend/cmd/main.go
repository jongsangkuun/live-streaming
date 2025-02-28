package main

import (
	"live-streaming-backend/pkg/config"
	"live-streaming-backend/pkg/utils"
	"log"
)

func main() {
	Config, err := config.EnvLoad()
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintStruct(Config)
}
