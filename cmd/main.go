package main

import (
	"EmployeeServiceWithQuickSortXml/cmd/application"
	"EmployeeServiceWithQuickSortXml/config"
)

func main() {
	cfg := config.ReadCfg()

	app := application.NewApp(cfg, nil)
	app.Init()
	app.Start()

}
