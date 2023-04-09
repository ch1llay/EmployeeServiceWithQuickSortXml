package Sorting

import "EmployeeServiceWithQuickSortXml/Model"

func sortPartByReportCount(arr []*Model.EmployeeFull, from, to int) int {
	left := from
	right := to

	index := from + (to-from)/2
	pivot := arr[index].Reports

	for left <= right {

		for len(arr[left].Reports) < len(pivot) {
			left++
		}

		for len(arr[right].Reports) > len(pivot) {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}

func quickSortByReportCount(arr []*Model.EmployeeFull, from, to int) {
	if from < to {
		pivotIndex := sortPartByReportCount(arr, from, to)

		quickSortByReportCount(arr, from, pivotIndex-1)
		quickSortByReportCount(arr, pivotIndex, to)
	}
}
func QuickSortByReportCount(arr []*Model.EmployeeFull) {
	quickSortByReportCount(arr, 0, len(arr))
}

func sortPartByBirthday(arr []*Model.EmployeeFull, from, to int) int {
	left := from
	right := to

	index := from + (to-from)/2
	pivot := arr[index].Birthday

	for left <= right {

		for arr[left].Birthday.Unix() < pivot.Unix() {
			left++
		}

		for arr[right].Birthday.Unix() > pivot.Unix() {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return left
}

func quickSortByBirthday(arr []*Model.EmployeeFull, from, to int) {
	if from < to {
		pivotIndex := sortPartByBirthday(arr, from, to)

		quickSortByBirthday(arr, from, pivotIndex-1)
		quickSortByBirthday(arr, pivotIndex, to)
	}
}

func QuickSortByBirthday(arr []*Model.EmployeeFull) {
	quickSortByBirthday(arr, 0, len(arr))
}
