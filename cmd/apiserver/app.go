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
	server  *http.Server
	config  *config.Config
	context context.Context
}

func NewApp(cfg *config.Config, ctx context.Context) *App {
	return &App{
		server:  nil,
		config:  cfg,
		context: ctx,
	}
}

func (a *App) Init() {
	fileRepository := repository.NewFileMongoRepository(a.config.MongoConnection, a.config.MongoDbName, a.config.MongoCollectionName)
	employeeRepository := repository.NewEmployeeRepository(a.config.PostgresConnection)
	fileService := service.NewFileService(fileRepository)
	employeeService := service.NewEmployeeService(employeeRepository)
	handlers := handler.NewHandler(employeeService, fileService)
	srv := apiserver.New()
	srv.Configure(a.config, handlers)
}

func (a *App) Start() {
	go func() {
		log.Println(fmt.Sprintf("Server started on port %s enviroment is %s", a.config.Port, a.config.Environment))
		a.server.ListenAndServe()
	}()
}
