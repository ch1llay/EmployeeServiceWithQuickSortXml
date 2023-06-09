package XMLHelper

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"encoding/xml"
	"time"
)

func GetXmlFile(employees []*Model.EmployeeFull) (*Model.File, error) {
	tmp := struct {
		Employees []*Model.EmployeeFull `xml:"employee"`
		XMLName   xml.Name              `xml:"Employees"`
	}{Employees: employees}

	data, err := xml.MarshalIndent(tmp, "", "   ")
	responseData := []byte(xml.Header + string(data))
	return &Model.File{FileName: "file.xml", Data: responseData, InsertDate: time.Now()}, err
}
