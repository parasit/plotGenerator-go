package main

import (
	"fmt"
)

func main() {
	var dt DataTable
	// dt.dataValues = make([0]DataItem)
	fmt.Println("Plot Generator")
	dt.ReadFromFile("Skarby/bronie.txt")
	fmt.Println(dt.GetSize())
	dt.DumpData()
	fmt.Println(dt.GetElementAtValue(1))

}
