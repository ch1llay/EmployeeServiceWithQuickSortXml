package main

import (
	"EmployeeServiceWithQuickSortXml/config"
	apiserver "EmployeeServiceWithQuickSortXml/internal"
	"EmployeeServiceWithQuickSortXml/internal/handler"
	"EmployeeServiceWithQuickSortXml/internal/repository"
	"EmployeeServiceWithQuickSortXml/internal/service"
	"context"
	"fmt"
	"log"
	"net/http"
)

type App struct {
	server   *http.Server
	settings *config.Config
	context  context.Context
}

func NewApp(cfg *config.Config, ctx context.Context) *App {
	return &App{
		server:   nil,
		settings: cfg,
		context:  ctx,
	}
}

func (a *App) InitApp() {
	fileRepository := repository.NewFileRepository(a.settings.MongoConnection, a.settings.MongoDbName, a.settings.MongoCollectionName)
	employeeRepository := repository.NewEmployeeRepository(a.settings.PostgresConnection, "")
	fileService := service.NewFileService(fileRepository)
	employeeService := service.NewEmployeeService(employeeRepository)
	handlers := handler.NewHandler(employeeService, fileService)
	srv := apiserver.New(handlers)
}

func (a *App) Start() {
	go func() {
		log.Println(fmt.Sprintf("Server started on port %s enviroment is %s", a.settings.Port, a.settings.Env))
		a.server.ListenAndServe()
	}()
}
