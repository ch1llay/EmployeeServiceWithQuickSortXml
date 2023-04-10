package listWorker

import (
	"EmployeeServiceWithQuickSortXml/Model"
	"fmt"
	"github.com/ch1llay/GoDoublyLinkedList"
)

func Working(employees []*Model.EmployeeFull) string {
	l := GoDoublyLinkedList.ConvertSliceToList(employees)
	s := fmt.Sprintf("%v", l)
	return s
}
