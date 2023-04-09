package Model

import "time"

//type FileMongo struct {
//	Id       string `json:"_id,omitempty"`
//	FileName string `json:"file_name"`
//	Data     []byte `json:"data"`
//}

type File struct {
	Id         string    `json:"id"`
	FileName   string    `json:"file_name"`
	InsertDate time.Time `json:"insert_date"`
	Data       []byte    `json:"data"`
}

type FileResponse struct {
	Guid        string `json:"guid"`
	TypeSorting string `json:"type_sorting"`
}
