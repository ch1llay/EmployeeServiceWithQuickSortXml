package main

import (
	"EmployeeServiceWithQuickSortXml/cmd/application"
	"EmployeeServiceWithQuickSortXml/config"
)

func main() {
	cfg := config.ReadCfg()
	//cxt, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//fmt.Printf("%v %v", cfg, cxt)

	app := application.NewApp(cfg, nil)
	app.Init()
	app.Start()

	//doneSignal := make(chan os.Signal, 1)
	//signal.Notify(doneSignal, syscall.SIGTERM, os.Interrupt)
	//<-doneSignal
}
