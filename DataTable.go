package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

//DataItem basic item for random choice
type DataItem struct {
	itemName  string
	itemValue int
}

//DataTable basic data type for random choice
type DataTable struct {
	dataValues []DataItem
}

// AddValue to DataTable
func (dt *DataTable) AddValue(name string, value int) {
	item := DataItem{name, value}
	dt.dataValues = append(dt.dataValues, item)
}

// AddItem adds item (struct) to DataTable
func (dt *DataTable) AddItem(item DataItem) {
	dt.dataValues = append(dt.dataValues, item)
}

// GetSize get virtual (!) size of all elements
func (dt DataTable) GetSize() int {
	result := 0
	for k := range dt.dataValues {
		result += dt.dataValues[k].itemValue
	}
	return result
}

//GetElementAtValue returns element from table at requested ID
func (dt DataTable) GetElementAtValue(value int) string {
	localIndex := 0
	for k := range dt.dataValues {
		if (value >= localIndex) && (value <= localIndex+dt.dataValues[k].itemValue) {
			return dt.dataValues[k].itemName
		}
		localIndex += dt.dataValues[k].itemValue
	}
	return "None"
}

// DumpData from DataTable in human readable format
func (dt DataTable) DumpData() {
	index := 1
	for k := range dt.dataValues {
		fmt.Printf("%s : ", dt.dataValues[k].itemName)
		if dt.dataValues[k].itemValue > 1 {
			fmt.Printf("%d-%d\n", index, index+dt.dataValues[k].itemValue-1)
			index += dt.dataValues[k].itemValue
		} else {
			fmt.Println(index)
			index++
		}
	}
}

// ReadFromFile loads DataItems from various type of files
func (dt *DataTable) ReadFromFile(filename string) (items int, err error) {
	rItems := 0
	var rErr error
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		item, err := ParseInputLine(line)
		if err == nil {
			dt.AddItem(item)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rItems, rErr
}

// ParseInputLine check and parse input data
func ParseInputLine(dataIn string) (DataItem, error) {

	// simple line
	// number:value
	re := regexp.MustCompile(`^(?P<value>\d{1,}):(?P<name>\S{1,})$`)
	match := re.FindStringSubmatch(dataIn)
	if len(match) > 0 {
		value, _ := strconv.Atoi(match[1])
		return DataItem{match[2], value}, nil
	}
	return DataItem{}, errors.New("Cannot parse line")
}
