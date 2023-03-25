package main

import (
	"EmployeeServiceWithQuickSortXml/internal"
	"EmployeeServiceWithQuickSortXml/internal/handler"
	"EmployeeServiceWithQuickSortXml/internal/repository"
	"EmployeeServiceWithQuickSortXml/internal/service"
	"log"
)

func main() {
	config := apiserver.NewConfig()
	fileRepository := repository.NewFileRepository(config.Mongo.ConnectionString, config.Mongo.DatabaseName, config.Mongo.CollectionName)
	employeeRepository := repository.NewEmployeeRepository(config.PG)
	fileService := service.NewFileService(fileRepository)
	employeeService := service.NewEmployeeService(employeeRepository)
	handlers := handler.NewHandler(employeeService, fileService)
	srv := apiserver.New()

	if err := srv.Run(config, handlers); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
