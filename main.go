package main

import (
	"iot/config"
	"iot/internal"
	"iot/migration"
)

func main() {
	config.New()
	Config := config.GetConfig()
	migration.InsertData(Config)
	internal.InitFiber(Config)

}
