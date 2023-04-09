package service

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"EmployeeServiceWithQuickSortXml/internal/repository"
)

type FileServ interface {
	Insert(file *Model.File) (string, error)
	GetById(guid string) (*Model.File, error)
	DeleteById(guid string) error
}

type FileService struct {
	FileRepository repository.FileRep
}

func NewFileService(fileRep repository.FileRep) *FileService {
	return &FileService{FileRepository: fileRep}
}

func (f *FileService) Insert(file *Model.File) (string, error) {
	return f.FileRepository.Insert(file)
}
func (f *FileService) GetByGuid(guid string) (*Model.File, error) {
	return f.FileRepository.GetById(guid)
}
func (f *FileService) DeleteById(guid string) error {
	return f.FileRepository.DeleteById(guid)
}
