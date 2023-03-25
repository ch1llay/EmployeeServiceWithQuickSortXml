package repository

type FileRepository struct {
	connectionString, databaseName, collectionName string
}

func NewFileRepository(connectionString, databaseName, collectionName string) *FileRepository {
	return &FileRepository{
		connectionString: connectionString,
		databaseName:     databaseName,
		collectionName:   collectionName,
	}
}

func (f *FileRepository) Write() {

}
func (f *FileRepository) Delete(guid string) error {
	return *new(error)
}
