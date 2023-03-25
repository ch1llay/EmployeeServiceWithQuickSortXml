package service

import (
	"EmployeeServiceWithQuickSortXml/Models"
	"EmployeeServiceWithQuickSortXml/internal/repository"
)

type FService interface {
	Write(file *Models.File) string
	GetByGuid(guid string) *Models.File
}

type FileService struct {
	FileRepository *repository.FileRepository
}

func NewFileService(fileRep *repository.FileRepository) *FileService {
	return &FileService{FileRepository: fileRep}
}

func (e *FileService) GetByGuid(guid string) *Models.File {
	return &Models.File{}
}

func (e *FileService) Delete(guid string) error {
	return e.FileRepository.Delete(guid)
}
