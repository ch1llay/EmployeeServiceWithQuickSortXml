package Searching

import "EmployeeServiceWithQuickSortXml/Model"

func BinarySearch(arr []*Model.EmployeeFull, id int) (int, *Model.EmployeeFull) {
	startIndex := 0
	endIndex := len(arr) - 1
	middleIndex := 0
	for startIndex <= endIndex {
		middleIndex = startIndex + (endIndex-startIndex)/2
		if arr[middleIndex].Id == id {
			return middleIndex, arr[middleIndex]
		}

		if arr[middleIndex].Id > id {
			endIndex = middleIndex - 1
		} else {
			startIndex = middleIndex + 1
		}
	}
	return -1, nil
}
