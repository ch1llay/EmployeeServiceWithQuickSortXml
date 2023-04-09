package Model

type File struct {
	Id       string `json:"_id,omitempty"`
	FileName string `json:"file_name"`
	Data     []byte `json:"data"`
}
