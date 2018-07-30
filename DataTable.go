package main

import (
	"fmt"
)

//DataTableItem basic table element
type DataTableItem struct {
	itemValue int
	itemName  string
}

//DataTable basic data type for random choice
type DataTable struct {
	Items []DataTableItem
}

// AddValue to DataTable
func (dt *DataTable) AddValue(value string, wage int) []DataTableItem {
	dti := DataTableItem{wage, value}
	dt.Items = append(dt.Items, dti)
	return dt.Items
}

// GetSize get virtual (!) size of all elements
func (dt DataTable) GetSize() int {
	result := 0
	for k := range dt.Items {
		result += dt.Items[k].itemValue
	}
	return result
}

//GetElementForValue returns element from table at requested ID
func (dt DataTable) GetElementForValue(value int) string {
	localIndex := 0
	for k := range dt.Items {
		if (value >= localIndex) && (value <= localIndex+dt.Items[k].itemValue) {
			return dt.Items[k].itemName
		}
		localIndex += dt.Items[k].itemValue
	}
	return "None"
}

// DumpData from DataTable in human readable format
func (dt DataTable) DumpData() {
	index := 1
	for k := range dt.Items {
		fmt.Printf("%s : ", dt.Items[k].itemName)
		if dt.Items[k].itemValue > 1 {
			fmt.Printf("%d-%d\n", index, index+dt.Items[k].itemValue-1)
			index += dt.Items[k].itemValue
		} else {
			fmt.Println(index)
			index++
		}
	}
}
