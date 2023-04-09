package repository

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FileRep interface {
	Insert(file *Model.File) (string, error)
	GetById(guid string) (*Model.File, error)
	DeleteById(guid string) error
}
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

func (f *FileRepository) mongoSession(action func(*mgo.Session) error) error {
	session, err := mgo.Dial(f.connectionString)
	if err != nil {
		return fmt.Errorf("MongoDB session creation error: %s", err.Error())
	}
	defer session.Close()

	if err := action(session); err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) Insert(file *Model.File) (guid string, err error) {
	err = f.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(f.databaseName).C(f.collectionName)
		if err := collection.Insert(file); err != nil {
			return err
		}
		insertedFile := new(Model.File)
		if err := collection.Find(bson.M{"filename": file.FileName}).One(&insertedFile); err != nil {
			return err
		}
		guid = insertedFile.Id
		return nil
	})

	return
}

func (f *FileRepository) GetById(guid string) (file *Model.File, err error) {
	err = f.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(f.databaseName).C(f.collectionName)

		insertedFile := new(Model.File)
		if err := collection.FindId(guid).One(&insertedFile); err != nil {
			return err
		}
		guid = insertedFile.Id
		return nil
	})

	return

}
func (f *FileRepository) DeleteById(guid string) (err error) {
	err = f.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(f.databaseName).C(f.collectionName)

		if err := collection.RemoveId(guid); err != nil {
			return err
		}
		return nil
	})

	return
}
