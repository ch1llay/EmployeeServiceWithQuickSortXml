package main

import (
	"EmployeeServiceWithQuickSortXml/cmd/application"
	"EmployeeServiceWithQuickSortXml/config"
)

func main() {
	cfg := config.ReadCfg()

	app := application.NewApp(cfg)
	app.Init()
	app.Start()

}
