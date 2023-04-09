package application

import (
	"EmployeeServiceWithQuickSortXml/config"
	"EmployeeServiceWithQuickSortXml/internal/apiserver"
	"EmployeeServiceWithQuickSortXml/internal/handler"
	"EmployeeServiceWithQuickSortXml/internal/repository"
	"EmployeeServiceWithQuickSortXml/internal/service"
	"context"
	"fmt"
	"log"
)

type App struct {
	server  *apiserver.Server
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
	fileRepository := repository.NewFilePGRepository(a.config.PostgresConnection)
	employeeRepository := repository.NewEmployeeRepository(a.config.PostgresConnection)
	fileService := service.NewFileService(fileRepository)
	employeeService := service.NewEmployeeService(employeeRepository)
	handlers := handler.NewHandler(employeeService, fileService)
	a.server = apiserver.New(a.config, handlers)
}

func (a *App) Start() {
	log.Println(fmt.Sprintf("Server started on port %s enviroment is %s", a.config.Port, a.config.Environment))
	a.server.Run()
}
