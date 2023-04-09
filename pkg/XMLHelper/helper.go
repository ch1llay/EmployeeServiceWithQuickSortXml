package XMLHelper

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"encoding/xml"
	"fmt"
	"os"
)

func Write(employees *Model.Employee) (file *Model.File, err error) {
	writer, err := os.Open("/tmp/tmp.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	encoder := xml.NewEncoder(writer)
	encoder.Encode(employees)
}
