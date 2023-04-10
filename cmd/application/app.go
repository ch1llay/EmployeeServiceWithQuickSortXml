package application

import (
	"EmployeeServiceWithQuickSortXml/config"
	"EmployeeServiceWithQuickSortXml/internal/apiserver"
	"EmployeeServiceWithQuickSortXml/internal/handler"
	"EmployeeServiceWithQuickSortXml/internal/repository"
	"EmployeeServiceWithQuickSortXml/internal/service"
	"fmt"
	"log"
)

type App struct {
	server *apiserver.Server
	config *config.Config
}

func NewApp(cfg *config.Config) *App {
	return &App{
		server: nil,
		config: cfg,
	}
}

func (a *App) Init() {
	fileRepository := repository.NewFilePGRepository(a.config.PostgresConnection)
	employeeRepository := repository.NewEmployeeRepository(a.config.PostgresConnection)
	reportRepository := repository.NewReportRepository(a.config.PostgresConnection)
	fileService := service.NewFileService(fileRepository)
	employeeService := service.NewEmployeeService(employeeRepository, reportRepository)
	reportService := service.NewReportService(reportRepository)
	handlers := handler.NewHandler(employeeService, reportService, fileService)
	//repository.InitRepository(a.config.PostgresConnection)
	a.server = apiserver.New(a.config, handlers)
}

func (a *App) Start() {
	log.Println(fmt.Sprintf("Server started on port %s enviroment is %s", a.config.Port, a.config.Environment))
	a.server.Run()
}
