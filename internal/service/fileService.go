package service

import (
	"EmployeeServiceWithQuickSortXml/Models"
	"EmployeeServiceWithQuickSortXml/internal/repository"
	"errors"
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

func (e *FileService) GetByGuid(guid string) (*Models.File, error) {
	return &Models.File{}, errors.New("")
}

func (e *FileService) Delete(guid string) error {
	return e.FileRepository.Delete(guid)
}

func (e *FileService) GetAll() []*Models.File {
	return make([]*Models.File, 10)
}
